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
	"sync"
)

//---------------------------------------------------------------------

// IssueList is the list of all the issues from all the projects read in
type IssueList struct {
	data  map[int]*Issue
	maxID int
	mutex *sync.Mutex
}

// NewIssueList makes a new IssueList
func NewIssueList() *IssueList {
	list := &IssueList{}
	list.data = make(map[int]*Issue)
	list.mutex = &sync.Mutex{}
	return list
}

func (list *IssueList) getMap() map[int]*Issue {
	return list.data
}

func (list *IssueList) issue(id int) (*Issue, bool) {
	issue, ok := list.data[id]
	return issue, ok
}

func (list *IssueList) getMaxID() int {
	return list.maxID
}

// AddList IS threadsafe!
func (list *IssueList) AddList(issues []*Issue) {
	list.mutex.Lock()
	defer list.mutex.Unlock()
	for _, issue := range issues {
		list.add(issue)
	}
}

// add is NOT threadsafe!
func (list *IssueList) add(issue *Issue) {

	id := issue.ID

	list.data[id] = issue

	if id > list.maxID {
		list.maxID = id
	}

	issue.errors = nil

	issue.Issues = list
}

func (list *IssueList) preread(apiKey string, projects map[int]*Project) (map[int]int, int, error) {

	const offset = 0
	const limit = 1

	max := map[int]int{}
	tot := 0

	for index, project := range projects {
		resp, err := makeRequest(apiKey, project.ID, offset, limit)
		if err != nil {
			return nil, 0, err
		}
		max[index] += resp.TotalCount
		tot += resp.TotalCount
	}

	return max, tot, nil
}

func (list *IssueList) readChunk(apiKey string, project *Project, offset int, limit int, issueCount int) error {

	resp, err := makeRequest(apiKey, project.ID, offset, limit)
	if err != nil {
		return err
	}

	list.AddList(resp.Issues)

	return nil
}

// Read gets all the issues in
func (list *IssueList) Read(apiKey string, wg *sync.WaitGroup, projects map[int]*Project) error {

	issueCounts, issueTotal, err := list.preread(apiKey, projects)
	if err != nil {
		panic(err)
	}

	const limit = 100

	for index, project := range projects {

		for offset := 0; offset < issueCounts[index]; offset += limit {

			wg.Add(1)

			go func(project *Project, offset int, limit int, issueTotal int) {
				defer wg.Done()
				err := list.readChunk(apiKey, project, offset, limit, issueTotal)
				if err != nil {
					panic(err)
				}
				frac := 100.0 * float64(len(list.data)) / float64(issueTotal)
				fmt.Printf("\r%d%%", int(frac))
			}(project, offset, limit, issueTotal)

		}
	}

	return nil
}
