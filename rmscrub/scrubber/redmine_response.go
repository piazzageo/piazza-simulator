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

package scrubber

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// IssueRequestResponse is what is returned from a query for issues
type IssueRequestResponse struct {
	Issues     []*Issue `json:"issues"`
	TotalCount int      `json:"total_count"`
	Offset     int      `json:"offset"`
	Limit      int      `json:"limit"`
}

func makeRequest(apiKey string, projectID int, offset int, limit int) (*IssueRequestResponse, error) {

	url := fmt.Sprintf("%s/%s?project_id=%d&offset=%d&limit=%d&status_id=*",
		"https://redmine.devops.geointservices.io",
		"/issues.json",
		projectID, offset, limit)

	client := &http.Client{}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	request.SetBasicAuth(apiKey, "random")

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("GET failed with status code %s", response.Status)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	resp := IssueRequestResponse{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
