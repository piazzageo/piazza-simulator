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
	Map   map[int]*Issue
	MaxId int
}

// returns issues table and highest id value
func getIssues(projectId int) (*Issues, error) {

	apiKey, err := getApiKey()
	if err != nil {
		return nil, err
	}

	issues := &Issues{}
	issues.Map = make(map[int]*Issue)

	offset := 0
	const limit = 100

	for {
		resp, err := makeRequest(apiKey, projectId, offset, limit)
		if err != nil {
			return nil, err
		}

		for _, issue := range resp.Issues {
			issues.Map[issue.Id] = issue
			if issue.Id > issues.MaxId {
				issues.MaxId = issue.Id
			}

			issue.errors = make([]string, 0)

			issue.Issues = issues
		}

		offset += limit
		if offset > resp.TotalCount {
			break
		}
	}

	return issues, nil
}
