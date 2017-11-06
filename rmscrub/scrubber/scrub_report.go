/*
Copyright 2016, RadiantBlue Technologies, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package scrubber

import (
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// ScrubReport is what embodies the ScrubReport
type ScrubReport struct {
	list *IssueList
}

// ScrubIssueResult is the issue object to be stored
type ScrubIssueResult struct {
	Errors []string `json:"errors"`
	Owner  string   `json:"owner"`
}

// ScrubResults is the report map object to be stored
type ScrubResults struct {
	Date time.Time                `json:"date"`
	Data map[int]ScrubIssueResult `json:"data"`
}

// NewScrubReport returns a new ScrubReport
func NewScrubReport(list *IssueList) *ScrubReport {
	r := &ScrubReport{
		list: list,
	}
	return r
}

// Run runs
func (r *ScrubReport) Run() error {
	for _, issue := range r.list.GetMap() {

		/*
			if !issue.isPastSprint() {
				fmt.Printf("%d - %s - %t\n", issue.Id, issue.targetVersion(), issue.isPastSprint())
			}
		*/

		if issue.isCurrentSprint() {
			r.runCommonRules(issue)
			r.runCurrentSprintRules(issue)
		} else if issue.isReadySprint() {
			r.runCommonRules(issue)
			r.runReadySprintRules(issue)
		} else if issue.isPastSprint() {
			r.runPastSprintRules(issue)
		} else if issue.isFutureSprint() {
			r.runCommonRules(issue)
			r.runFutureSprintRules(issue)
		} else if issue.isBacklogSprint() {
			r.runCommonRules(issue)
			r.runBacklogSprintRules(issue)
		} else if issue.isEpicSprint() {
			r.runEpicSprintRules(issue)
		} else {
			r.runInvalidSprintRules(issue)
		}

	}

	return nil
}

//
// just some common rules used by the functions below
//
func (r *ScrubReport) runCommonRules(issue *Issue) {

	parent := issue.parent()

	if !issue.isValidStatus() {
		issue.Error("status is invalid")
	}

	if issue.hasStartDate() {
		issue.Error("start date is not empty")
	}

	if issue.hasDueDate() {
		issue.Error("due date is not empty")
	}

	if issue.hasEstimatedTime() {
		issue.Error("estimated time is not empty")
	}

	if parent != nil && parent.isClosed() && (!issue.isClosed() && !issue.isRejected()) {
		issue.Error("issue's parent is closed but issue is not")
	}

	if parent != nil && parent.isResolved() && (!issue.isResolved() && !issue.isClosed()) {
		issue.Error("issue's parent is resolved but issue is not resolved or closed")
	}

	if parent != nil && parent.isRejected() && !issue.isRejected() {
		issue.Error("issue's parent is rejected but issue is not")
	}

	if !issue.hasValidParent() {
		issue.Error(fmt.Sprintf("issue's parent %d is invalid", issue.Parent.ID))
	}

	if issue.hasBadTitleTag() {
		issue.Error(fmt.Sprintf("issue has invalid title tag: %s", issue.Subject))
	}
}

//
// rules for issues in the Current sprint
//
func (r *ScrubReport) runCurrentSprintRules(issue *Issue) {

	parent := issue.parent()

	if issue.isEpic() {
		if parent != nil {
			issue.Error("epic's parent is not set")
		}
	} else if issue.isStory() {
		if parent == nil {
			issue.Error("story's parent is not set")
		} else if parent.tracker() != "Epic" {
			issue.Error("story's parent is not an Epic")
		}
	} else if issue.isTask() {
		if parent == nil {
			issue.Error("task's parent is not set")
		} else if parent.tracker() != "Story" && parent.tracker() != "Bug" {
			issue.Error("task's parent is not an Story or Bug")
		}
	} else if issue.isBug() {
		if parent == nil {
			issue.Error("bug's parent is not set")
		} else if parent.tracker() != "Epic" && parent.tracker() != "Story" {
			issue.Error("bug's parent is not a Story or Bug")
		}
	} else {
		issue.Error(fmt.Sprintf("issue's type (\"tracker\") is invalid: \"%s\"", issue.tracker()))
	}

	if !issue.isAssigned() && (!issue.isNew() && !issue.isRejected()) {
		issue.Error("issue in current sprint is not assigned to anyone")
	}

	if issue.category() == "" {
		issue.Error("'category' not set")
	}

	if issue.isClosed() && issue.percentDone() != 100 {
		issue.Error("'percent done' on closed issue is not 100%")
	}

	if issue.isResolved() && issue.percentDone() != 100 {
		issue.Error("'percent done' on resolved issue is not 100%")
	}

	if issue.isNew() && issue.percentDone() != 0 {
		issue.Error("'percent done' on new issue is not 0%")
	}

	if issue.isInProgress() && issue.percentDone() == 0 {
		issue.Error("'percent done' on in-progress issue is 0%")
	}

	if issue.isInProgress() && issue.percentDone() == 100 {
		issue.Error("'percent done' on in-progress issue is 100%")
	}

	if issue.isStory() && !issue.isRejected() && !issue.hasAcceptanceCriteria() {
		issue.Error("story does not have any acceptance criteria")
	}

	if issue.isBug() && !issue.isRejected() && !issue.hasAcceptanceCriteria() {
		issue.Error("bug does not have any acceptance criteria")
	}
}

//
// rules for issues in a Past sprint
//
func (r *ScrubReport) runPastSprintRules(issue *Issue) {
	if !issue.isClosed() && !issue.isRejected() {
		issue.Error("issue from past sprint is not closed or rejected")
	}
}

//
// rules for issues in a Future sprint
//
func (r *ScrubReport) runFutureSprintRules(issue *Issue) {

	if issue.isClosed() {
		issue.Error("issue from future sprint is already closed")
	}
	if issue.isRejected() {
		issue.Error("issue from future sprint is already rejected")
	}
	if issue.isResolved() {
		issue.Error("issue from future sprint is already resolved")
	}
}

//
// rules for issues in the Backlog
//
func (r *ScrubReport) runBacklogSprintRules(issue *Issue) {

	if !issue.isNew() && !issue.isRejected() {
		issue.Error("issue in backlog does not have status 'new' or 'rejected'")
	}
}

//
// rules for issues with target version set to "Pz Epic"
//
func (r *ScrubReport) runEpicSprintRules(issue *Issue) {
	if !issue.isEpic() {
		issue.Error("issue is not an Epic but has target version set to \"Pz Epic\"")
	}
}

//
// rules for issues with no target version set at all
//
func (r *ScrubReport) runInvalidSprintRules(issue *Issue) {
	if issue.targetVersion() == "" {
		issue.Error("no sprint set")
	} else {
		issue.Error(fmt.Sprintf("invalid sprint: %s", issue.targetVersion()))
	}
}

//
// rules for issues with target version set to "Ready"
//
func (r *ScrubReport) runReadySprintRules(issue *Issue) {
	//	if !issue.isEpic() {
	//		issue.Errorf("Issue is not an Epic but has target version set to \"Pz Epic\"")
	//	}
}

// Report returns the text report
func (r *ScrubReport) Report() *ScrubResults {

	result := &ScrubResults{
		Date: time.Now(),
		Data: map[int]ScrubIssueResult{},
	}

	for id, issue := range r.list.GetMap() {

		if issue.errors == nil || len(issue.errors) == 0 {
			continue
		}

		owner := "(no owner)"
		if issue.isAssigned() {
			owner = issue.assignee()
		}

		data := ScrubIssueResult{
			Owner:  owner,
			Errors: []string{},
		}

		for _, s := range issue.errors {
			data.Errors = append(data.Errors, s)
		}

		result.Data[id] = data
	}

	return result
}

func (r *ScrubResults) String() string {
	s := ""
	for id, result := range r.Data {
		for _, mssg := range result.Errors {
			s += fmt.Sprintf("%d:  %s  %s\n", id, result.Owner, mssg)
		}
	}
	return s
}

// Store writes to the database
func (r *ScrubResults) Store() error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	av, err := dynamodbattribute.MarshalMap(r)

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("RedmineScrubberTable"),
	}

	_, err = svc.PutItem(input)

	if err != nil {
		return err
	}

	return nil
}

func slackStyle(name string) string {
	switch name {

	case "(no owner)":
		return "*NO OWNER*"

	case "Gregory Barrett":
		return "@gbarrett"

	case "Jeffrey Yutzler":
		name = "Jeff Yutzler"
	case "Ben Hosmer":
		name = "Benjamin Hosmer"
	case "Benjamin Peizer":
		name = "Ben Peizer"
	case "Marjorie Lynum":
		name = "Marge Lynum"
	}

	name = "@" + strings.Replace(name, " ", ".", 1)
	name = strings.ToLower(name)
	return name
}
