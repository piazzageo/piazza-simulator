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

func buildOwnerToIssuesMap(issues *Issues) map[string][]*Issue {
	m := map[string][]*Issue{}

	for i := 0; i <= issues.MaxId; i++ {
		issue, ok := issues.Issue(i)
		if !ok {
			continue
		}

		owner := "(no owner)"
		if issue.isAssigned() {
			owner = issue.assignee()
		}

		_, ok2 := m[owner]
		if !ok2 {
			m[owner] = make([]*Issue, 0)
		}
		m[owner] = append(m[owner], issue)
	}

	return m
}

func getSortedKeys(m map[string][]*Issue) []string {
	keys := []string{}
	for owner, _ := range m {
		keys = append(keys, owner)
	}

	sort.Strings(keys)

	return keys
}

func showErrors(issues *Issues) {

	totalErrors := 0

	ownerToIssuesMap := buildOwnerToIssuesMap(issues)

	names := getSortedKeys(ownerToIssuesMap)

	for _, owner := range names {
		issues := ownerToIssuesMap[owner]

		errs := 0
		for _, issue := range issues {
			errs += len(issue.errors)
		}

		if errs > 0 {
			Printf("%s (%d)", owner, errs)
			for _, issue := range issues {
				for _, s := range issue.errors {
					Printf("\t%d: %s", issue.Id, s)
				}
			}

			totalErrors += errs
		}
	}

	Printf("Found %d errors ", totalErrors)
}

func main() {
	targetProjectNames := []string(nil)
	if len(os.Args) == 2 {
		targetProjectNames = []string{os.Args[1]}
	} else {
		targetProjectNames = []string{"Piazza", "Beachfront"}
	}

	apiKey, err := getApiKey()
	if err != nil {
		Errorf("Failed to get API key: %s", err)
	}

	availableProjects, err := getProjects(apiKey)
	if err != nil {
		Errorf("Failed to get projects: %s", err)
	}

	targetProjects := availableProjects.Filter(targetProjectNames)
	if len(targetProjects.GetMap()) == 0 {
		Errorf("Unknown project name(s): %s", strings.Join(targetProjectNames, " ,"))
	}

	allIssues := NewIssues()
	for _, project := range targetProjects.GetMap() {
		issues, err := getIssues(project)
		if err != nil {
			Errorf(err.Error())
		}
		allIssues.Merge(issues)
	}

	Logf("%d issues in project", len(allIssues.GetMap()))

	for _, issue := range allIssues.GetMap() {

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

	showErrors(allIssues)
}
