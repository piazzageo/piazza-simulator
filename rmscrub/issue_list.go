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
)

//---------------------------------------------------------------------

type IssueList struct {
	data  map[int]*Issue
	maxId int
}

func NewIssueList() *IssueList {
	list := &IssueList{}
	list.data = make(map[int]*Issue)
	return list
}

func (list *IssueList) GetMap() map[int]*Issue {
	return list.data
}

func (list *IssueList) Issue(id int) (*Issue, bool) {
	issue, ok := list.data[id]
	return issue, ok
}

func (list *IssueList) MaxId() int {
	return list.maxId
}

func (list *IssueList) Add(issue *Issue) {
	list.data[issue.Id] = issue

	if issue.Id > list.maxId {
		list.maxId = issue.Id
	}

	issue.errors = make([]string, 0)

	issue.Issues = list
}

// returns issues table and highest id value
func (list *IssueList) Read(project *Project) error {

	apiKey, err := getApiKey()
	if err != nil {
		return err
	}

	offset := 0
	const limit = 100

	for {
		resp, err := makeRequest(apiKey, project.Id, offset, limit)
		if err != nil {
			return err
		}

		for _, issue := range resp.Issues {
			list.Add(issue)
		}

		offset += limit
		if offset > resp.TotalCount {
			break
		}
		perc := (float64(offset) / float64(resp.TotalCount)) * 100.0
		fmt.Fprintf(os.Stderr, "\r%d%%", int(perc))
	}
	fmt.Fprintf(os.Stderr, "\r100%%\n")

	return nil
}

// returns issues table and highest id value
func (list *IssueList) Merge(newIssue *IssueList) error {
	for _, v := range newIssue.GetMap() {
		list.Add(v)
	}
	return nil
}
