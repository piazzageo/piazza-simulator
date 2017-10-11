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
	"os"
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

// Read returns issues table and highest id value
func (list *IssueList) Read(wg *sync.WaitGroup, project *Project) error {

	apiKey, err := getAPIKey()
	if err != nil {
		return err
	}

	offset := 0
	const limit = 100

	resp, err := makeRequest(apiKey, project.ID, offset, limit)
	if err != nil {
		return err
	}
	max := resp.TotalCount

	readChunk := func(offset, limit int) error {
		resp, err := makeRequest(apiKey, project.ID, offset, limit)
		if err != nil {
			return err
		}

		list.AddList(resp.Issues)

		fmt.Fprintf(os.Stderr, ".")
		return nil
	}

	for offset := 0; offset < max; offset += limit {
		wg.Add(1)
		go func(offset, limit int) {
			defer wg.Done()
			err = readChunk(offset, limit)
			if err != nil {
				panic(err)
			}
		}(offset, limit)
	}

	fmt.Fprintf(os.Stderr, "*")

	return nil
}
