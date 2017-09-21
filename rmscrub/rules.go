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

//
// just some common rules used by the functions below
//
func runCommonRules(issue *Issue) {

	parent := issue.parent()

	if issue.hasStartDate() {
		issue.Errorf("start date is not empty")
	}

	if issue.hasDueDate() {
		issue.Errorf("due date is not empty")
	}

	if issue.hasEstimatedTime() {
		issue.Errorf("estimated time is not empty")
	}

	if parent != nil && parent.isClosed() && (!issue.isClosed() && !issue.isRejected()) {
		issue.Errorf("issue's parent is closed but issue is not")
	}

	if parent != nil && parent.isResolved() && !issue.isResolved() {
		issue.Errorf("issue's parent is resolved but issue is not")
	}

	if parent != nil && parent.isRejected() && !issue.isRejected() {
		issue.Errorf("issue's parent is rejected but issue is not")
	}

	if !issue.hasValidParent() {
		issue.Errorf("issue's parent %d is invalid", issue.Parent.Id)
	}
}

//
// rules for issues in the Current sprint
//
func runCurrentSprintRules(issue *Issue) {
	runCommonRules(issue)

	parent := issue.parent()

	if issue.isEpic() {
		if parent != nil {
			issue.Errorf("epic's parent is not set")
		}
	} else if issue.isStory() {
		if parent == nil {
			issue.Errorf("story's parent is not set")
		} else if parent.tracker() != "Epic" {
			issue.Errorf("story's parent is not an Epic")
		}
	} else if issue.isTask() {
		if parent == nil {
			issue.Errorf("task's parent is not set")
		} else if parent.tracker() != "Story" && parent.tracker() != "Bug" {
			issue.Errorf("task's parent is not an Story or Bug")
		}
	} else if issue.isBug() {
		if parent == nil {
			issue.Errorf("bug's parent is not set")
		} else if parent.tracker() != "Epic" && parent.tracker() != "Story" {
			issue.Errorf("bug's parent is not a Story or Bug")
		}
	} else {
		issue.Errorf("issue's type (\"tracker\") is invalid: \"%s\"", issue.tracker())
	}

	if !issue.isAssigned() {
		issue.Errorf("issue is not assigned to anyone")
	}

	if issue.category() == "" {
		issue.Errorf("'category' not set")
	}

	if issue.isClosed() && issue.percentDone() != 100 {
		issue.Errorf("'percent done' on closed issue is not 100%%")
	}

	if issue.isResolved() && issue.percentDone() != 100 {
		issue.Errorf("'percent done' on resolved issue is not 100%%")
	}

	if issue.isNew() && issue.percentDone() != 0 {
		issue.Errorf("'percent done' on new issue is not 0%%")
	}

	if issue.isInProgress() && issue.percentDone() == 0 {
		issue.Errorf("'percent done' on in-progress issue is 0%%")
	}

	if issue.isInProgress() && issue.percentDone() == 100 {
		issue.Errorf("'percent done' on in-progress issue is 100%%")
	}

	if issue.isStory() && !issue.isRejected() && !issue.hasAcceptanceCriteria() {
		issue.Errorf("story does not have any acceptance criteria")
	}

	if issue.isBug() && !issue.isRejected() && !issue.hasAcceptanceCriteria() {
		issue.Errorf("bug does not have any acceptance criteria")
	}
}

//
// rules for issues in a Past sprint
//
func runPastSprintRules(issue *Issue) {
	if !issue.isClosed() && !issue.isRejected() {
		issue.Errorf("issue from past sprint is not closed or rejected")
	}
}

//
// rules for issues in a Future sprint
//
func runFutureSprintRules(issue *Issue) {

	runCommonRules(issue)

	if issue.isClosed() {
		issue.Errorf("issue from future sprint is already closed")
	}
	if issue.isRejected() {
		issue.Errorf("issue from future sprint is already rejected")
	}
	if issue.isResolved() {
		issue.Errorf("issue from future sprint is already resolved")
	}
}

//
// rules for issues in the Backlog
//
func runBacklogSprintRules(issue *Issue) {

	runCommonRules(issue)

	if !issue.isNew() && !issue.isRejected() {
		issue.Errorf("issue in backlog does not have status 'new' or 'rejected'")
	}
}

//
// rules for issues with target version set to "Pz Epic"
//
func runEpicSprintRules(issue *Issue) {
	if !issue.isEpic() {
		issue.Errorf("issue is not an Epic but has target version set to \"Pz Epic\"")
	}
}

//
// rules for issues with no target version set at all
//
func runInvalidSprintRules(issue *Issue) {
	if issue.targetVersion() == "" {
		issue.Errorf("no sprint set")
	} else {
		issue.Errorf("invalid sprint: %s", issue.targetVersion())
	}
}

//
// rules for issues with target version set to "Ready"
//
func runReadySprintRules(issue *Issue) {
	//	if !issue.isEpic() {
	//		issue.Errorf("Issue is not an Epic but has target version set to \"Pz Epic\"")
	//	}
}
