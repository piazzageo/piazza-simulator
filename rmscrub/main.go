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
	"sync"
)

// DEBUG just for testing
var DEBUG = false

func getAPIKey() (string, error) {
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

func buildOwnerToIssuesMap(list *IssueList) map[string][]*Issue {
	m := map[string][]*Issue{}

	for i := 0; i <= list.getMaxID(); i++ {
		issue, ok := list.issue(i)
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
	for owner := range m {
		keys = append(keys, owner)
	}

	sort.Strings(keys)

	return keys
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
	}

	name = "@" + strings.Replace(name, " ", ".", 1)
	name = strings.ToLower(name)
	return name
}

func showErrors(list *IssueList) {

	totalErrors := 0

	ownerToIssuesMap := buildOwnerToIssuesMap(list)

	names := getSortedKeys(ownerToIssuesMap)

	for _, owner := range names {
		issues := ownerToIssuesMap[owner]

		errs := 0
		for _, issue := range issues {
			errs += len(issue.errors)
		}

		if errs > 0 {
			for _, issue := range issues {
				for _, s := range issue.errors {
					Printf("%s  ==>  %d: %s", slackStyle(owner), issue.ID, s)
				}
			}

			totalErrors += errs
		}
	}

	//Printf("Found %d errors ", totalErrors)
}

func main() {
	targetProjectNames := []string(nil)
	if len(os.Args) == 2 {
		targetProjectNames = []string{os.Args[1]}
	} else {
		targetProjectNames = []string{"Piazza", "Beachfront"}
	}

	apiKey, err := getAPIKey()
	if err != nil {
		Errorf("Failed to get API key: %s", err)
	}

	availableProjects, err := NewProjectList(apiKey)
	if err != nil {
		Errorf("Failed to get projects: %s", err)
	}

	availableProjects.filter(targetProjectNames)
	if len(availableProjects.getMap()) == 0 {
		Errorf("Unknown project name(s): %s", strings.Join(targetProjectNames, " ,"))
	}

	wg := &sync.WaitGroup{}
	allIssues := NewIssueList()
	allIssues.Read(apiKey, wg, availableProjects.getMap())
	wg.Wait()

	fmt.Fprintf(os.Stderr, "\n")

	Logf("%d issues in projects", len(allIssues.getMap()))

	rules := Rules{}
	rules.run(allIssues)

	showErrors(allIssues)

	fmt.Printf("\n")

	tagChecker := NewTagCheker()
	tagChecker.run(allIssues)
	tagChecker.report()
}
