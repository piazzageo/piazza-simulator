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

// This app reads in a CSV export of a Redmine project and "validates" all
// the issues based on various rules such as:
//    a Story must have a parent, and that parent must be an Epic
//    a Task in the current milestone must have an estimate, and that
//      estimate must be <= 16 hours
// This app is still underdevelopment -- will be adding rules as we need them.
//
// To use:
//   - in Redmine, click on "View all issues" (right-hand column, top)
//   - then click on "Also available in... CSV" (main panel, bottom-right)
//   - run the app:  $ go run redmine-scrub.go downloadedfile.csv
//

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

var DEBUG = false

func getApiKey() (string, error) {
	home := os.Getenv("HOME")
	if home == "" {
		return "", fmt.Errorf("$HOME not found")
	}

	key, err := ioutil.ReadFile(home + "/.rm-scrub")
	if err != nil {
		return "", err
	}

	s := strings.TrimSpace(string(key))
	Logf("API Key: %s", s)

	return s, nil
}

func showErrors(issues *Issues) {

	errors := 0
	cleanIssues := 0
	dirtyIssues := 0

	byOwner := map[string][]*Issue{}

	// sort the issues by owner
	for i := 0; i <= issues.MaxId; i++ {
		issue, ok := issues.Map[i]
		if !ok {
			continue
		}

		owner := "(no owner)"
		if issue.isAssigned() {
			owner = issue.assignee()
		}

		_, ok2 := byOwner[owner]
		if !ok2 {
			byOwner[owner] = make([]*Issue, 0)
		}
		byOwner[owner] = append(byOwner[owner], issue)

		if len(issue.errors) > 0 {
			dirtyIssues++
		} else {
			cleanIssues++
		}
	}

	names := []string{}
	for owner, _ := range byOwner {
		names = append(names, owner)
	}

	sort.Strings(names)

	for _, owner := range names {
		issues := byOwner[owner]

		titled := false

		for _, issue := range issues {
			for _, s := range issue.errors {
				if !titled {
					Printf("%s", owner)
					titled = true
				}
				Printf("\t%d: %s", issue.Id, s)
				errors++
			}
		}
	}

	Printf("Found %d errors (across %d of %d total issues)",
		errors, dirtyIssues, dirtyIssues+cleanIssues)
}

func main() {
	if len(os.Args) != 2 {
		Errorf("usage:  $ rm-scrub project")
	}

	apiKey, err := getApiKey()
	if err != nil {
		Errorf("Failed to get API key: %s", err)
	}

	projectsResponse, err := getProjects(apiKey, 0, 100)
	if err != nil {
		Errorf("Failed to get projects: %s", err)
	}

	project := os.Args[1]
	projectId := -1
	for _, p := range (*projectsResponse).Projects {
		if p.Identifier == project || p.Name == project {
			projectId = p.Id
			break
		}
	}
	if projectId == -1 {
		Errorf("Unknown project name: %s", project)
	}

	issues, err := getIssues(projectId)
	if err != nil {
		Errorf(err.Error())
	}

	Logf("%d issues in project", len(issues.Map))

	for _, issue := range issues.Map {

		/*
			if !issue.isPastSprint() {
				fmt.Printf("%d - %s - %t\n", issue.Id, issue.targetVersion(), issue.isPastSprint())
			}
		*/

		if issue.isCurrentSprint() {
			runCurrentSprintRules(issue)
		} else if issue.isReadySprint() {
			runReadySprintRules(issue)
		} else if issue.isPastSprint() {
			runPastSprintRules(issue)
		} else if issue.isFutureSprint() {
			runFutureSprintRules(issue)
		} else if issue.isBacklogSprint() {
			runBacklogSprintRules(issue)
		} else if issue.isEpicSprint() {
			runEpicSprintRules(issue)
		} else {
			runInvalidSprintRules(issue)
		}

	}

	showErrors(issues)
}
