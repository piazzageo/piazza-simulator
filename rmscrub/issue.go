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
)

//---------------------------------------------------------------------

type RmObject struct {
	Id   int
	Name string
}

type Issue struct {
	Id             int      `json:"id"`
	Project        RmObject `json:"project"`
	Tracker        RmObject `json:"tracker"`
	Status         RmObject `json:"status"`
	Priority       RmObject `json:"priority"`
	Author         RmObject `json:"author"`
	AssignedTo     RmObject `json:"assigned_to"`
	Category       RmObject `json:"category"`
	FixedVersion   RmObject `json:"fixed_version"`
	Subject        string   `json:"subject"`
	Description    string   `json:"description"`
	DoneRatio      int      `json:"done_ratio"`
	CreatedOn      string   `json:"created_on"`
	UpdatedOn      string   `json:"updated_on"`
	StartDate      string   `json:"start_date"`
	DueDate        string   `json:"due_date"`
	EstimatedHours float64  `json:"estimated_hours"`
	Parent         RmObject `json:"parent"`
	StoryPoints    int      `json:"story_points"`

	Issues *Issues // back link

	warnings []string
	errors   []string
}

//---------------------------------------------------------------------

func (issue *Issue) Warnf(mssg string, args ...interface{}) {
	s := fmt.Sprintf(mssg, args...)
	issue.warnings = append(issue.warnings, s)
}

func (issue *Issue) Errorf(mssg string, args ...interface{}) {
	s := fmt.Sprintf(mssg, args...)
	issue.errors = append(issue.errors, s)
}

func (issue *Issue) category() string {
	return issue.Category.Name
}

func (issue *Issue) tracker() string {
	return issue.Tracker.Name
}

func (issue *Issue) hasValidParent() bool {
	pid := issue.Parent.Id
	if pid == 0 {
		return true
	}

	_, ok := issue.Issues.Map[pid]
	return ok
}

func (issue *Issue) parent() *Issue {
	if !issue.hasValidParent() {
		return nil
	}
	pid := issue.Parent.Id
	p, _ := issue.Issues.Map[pid]
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
	return issue.targetVersion() == CURRENT_SPRINT
}

func (issue *Issue) isReadySprint() bool {
	return issue.targetVersion() == READY_SPRINT
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
	return issue.targetVersion() == BACKLOG_SPRINT
}

func (issue *Issue) isEpicSprint() bool {
	return issue.targetVersion() == EPIC_SPRINT
}

func (issue *Issue) isNoSprint() bool {
	return issue.targetVersion() == ""
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
		Errorf("Regep failed for issue %d", issue.Id)
	}
	return matched
}
