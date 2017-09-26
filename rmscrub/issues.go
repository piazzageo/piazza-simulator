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

//---------------------------------------------------------------------

type Issues struct {
	data  map[int]*Issue
	MaxId int
}

func NewIssues() *Issues {
	issues := &Issues{}
	issues.data = make(map[int]*Issue)
	return issues
}

func (issues *Issues) GetMap() map[int]*Issue {
	return issues.data
}

func (issues *Issues) Issue(id int) (*Issue, bool) {
	issue, ok := issues.data[id]
	return issue, ok
}

func (issues *Issues) Add(issue *Issue) {
	issues.data[issue.Id] = issue

	if issue.Id > issues.MaxId {
		issues.MaxId = issue.Id
	}

	issue.errors = make([]string, 0)

	issue.Issues = issues
}

// returns issues table and highest id value
func getIssues(project *Project) (*Issues, error) {

	apiKey, err := getApiKey()
	if err != nil {
		return nil, err
	}

	issues := NewIssues()

	offset := 0
	const limit = 100

	for {
		resp, err := makeRequest(apiKey, project.Id, offset, limit)
		if err != nil {
			return nil, err
		}

		for _, issue := range resp.Issues {
			issues.Add(issue)
		}

		offset += limit
		if offset > resp.TotalCount {
			break
		}
	}

	return issues, nil
}

// returns issues table and highest id value
func (issues *Issues) Merge(newIssue *Issues) error {
	for _, v := range newIssue.GetMap() {
		issues.Add(v)
	}
	return nil
}
