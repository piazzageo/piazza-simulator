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

package main

import (
	"fmt"
	"regexp"
	"strings"
)

//---------------------------------------------------------------------

// RedmineObject is how Redmine represents a generic Thing in JSON
type RedmineObject struct {
	ID   int
	Name string
}

// Issue represents a Redmine Issue
type Issue struct {
	ID             int           `json:"id"`
	Project        RedmineObject `json:"project"`
	Tracker        RedmineObject `json:"tracker"`
	Status         RedmineObject `json:"status"`
	Priority       RedmineObject `json:"priority"`
	Author         RedmineObject `json:"author"`
	AssignedTo     RedmineObject `json:"assigned_to"`
	Category       RedmineObject `json:"category"`
	FixedVersion   RedmineObject `json:"fixed_version"`
	Subject        string        `json:"subject"`
	Description    string        `json:"description"`
	DoneRatio      int           `json:"done_ratio"`
	CreatedOn      string        `json:"created_on"`
	UpdatedOn      string        `json:"updated_on"`
	StartDate      string        `json:"start_date"`
	DueDate        string        `json:"due_date"`
	EstimatedHours float64       `json:"estimated_hours"`
	Parent         RedmineObject `json:"parent"`
	StoryPoints    int           `json:"story_points"`

	Issues *IssueList // back link

	errors []string
}

//---------------------------------------------------------------------

// Errorf records an error message
func (issue *Issue) Errorf(mssg string, args ...interface{}) {
	s := fmt.Sprintf(mssg, args...)
	if issue.errors == nil {
		issue.errors = make([]string, 0)
	}
	issue.errors = append(issue.errors, s)
}

func (issue *Issue) category() string {
	return issue.Category.Name
}

func (issue *Issue) tracker() string {
	return issue.Tracker.Name
}

func (issue *Issue) hasValidParent() bool {
	pid := issue.Parent.ID
	if pid == 0 {
		return true
	}

	_, ok := issue.Issues.issue(pid)
	return ok
}

func (issue *Issue) parent() *Issue {
	if !issue.hasValidParent() {
		return nil
	}
	pid := issue.Parent.ID
	p, _ := issue.Issues.issue(pid)
	return p
}

func (issue *Issue) targetVersion() string {
	return issue.FixedVersion.Name
}

func (issue *Issue) isEpic() bool {
	return issue.tracker() == "Epic"
}

func (issue *Issue) isStory() bool {
	return issue.tracker() == "Story"
}

func (issue *Issue) isTask() bool {
	return issue.tracker() == "Task"
}

func (issue *Issue) isBug() bool {
	return issue.tracker() == "Bug"
}

func (issue *Issue) isCurrentSprint() bool {
	return issue.targetVersion() == CurrentSprint
}

func (issue *Issue) isReadySprint() bool {
	return issue.targetVersion() == ReadySprint
}

func (issue *Issue) isPastSprint() bool {
	for _, s := range PastSprints {
		if issue.targetVersion() == s {
			return true
		}
	}
	return false
}

func (issue *Issue) isFutureSprint() bool {
	for _, s := range FutureSprints {
		if issue.targetVersion() == s {
			return true
		}
	}
	return false
}

func (issue *Issue) isBacklogSprint() bool {
	return issue.targetVersion() == BacklogSprint
}

func (issue *Issue) isEpicSprint() bool {
	return issue.targetVersion() == EpicSprint1 || issue.targetVersion() == EpicSprint2
}

func (issue *Issue) isNoSprint() bool {
	return issue.targetVersion() == ""
}

func (issue *Issue) isValidStatus() bool {
	return issue.isResolved() || issue.isRejected() ||
		issue.isClosed() || issue.isNew() || issue.isInProgress()
}

func (issue *Issue) isResolved() bool {
	return issue.Status.Name == "Resolved"
}

func (issue *Issue) isRejected() bool {
	return issue.Status.Name == "Rejected"
}

func (issue *Issue) isClosed() bool {
	return issue.Status.Name == "Closed"
}

func (issue *Issue) isNew() bool {
	return issue.Status.Name == "New"
}

func (issue *Issue) isInProgress() bool {
	return issue.Status.Name == "In Progress"
}

func (issue *Issue) assignee() string {
	return issue.AssignedTo.Name
}

func (issue *Issue) isAssigned() bool {
	return issue.AssignedTo.Name != ""
}

func (issue *Issue) hasStartDate() bool {
	return issue.StartDate != ""
}

func (issue *Issue) hasDueDate() bool {
	return issue.DueDate != ""
}

func (issue *Issue) hasEstimatedTime() bool {
	return issue.EstimatedHours != 0
}

func (issue *Issue) percentDone() int {
	return issue.DoneRatio
}

func (issue *Issue) hasAcceptanceCriteria() bool {
	d := issue.Description
	matched, err := regexp.MatchString("A/C", d)
	if err != nil {
		Errorf("Regep failed for issue %d", issue.ID)
	}
	return matched
}

func (issue *Issue) hasBadTitleTag() bool {
	subj := issue.Subject

	if !strings.ContainsAny(subj, "[]") {
		return false
	}

	// remove the legal tags
	for _, tag := range TitleTags {
		subj = strings.Replace(subj, tag, "", 1)
	}

	// if we still have markers, we've bad
	return strings.ContainsAny(subj, "[]")
}
