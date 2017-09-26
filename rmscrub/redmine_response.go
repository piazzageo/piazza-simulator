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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type RedmineResponse struct {
	Issues     []*Issue `json:"issues"`
	TotalCount int      `json:"total_count"`
	Offset     int      `json:"offset"`
	Limit      int      `json:"limit"`
}



func makeRequest(apiKey string, projectId int, offset int, limit int) (*RedmineResponse, error) {

	url := fmt.Sprintf("%s/%s?project_id=%d&offset=%d&limit=%d&status_id=*",
		"https://redmine.devops.geointservices.io",
		"/issues.json",
		projectId, offset, limit)

	client := &http.Client{}

	Logf("\n((((((((((((((((((((((\nurl: %s", url)

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

	Logf("Raw result: %s...", string(body)[0:60])

	resp := RedmineResponse{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	Logf("request: offset=%d, limit=%d, totalCount=%d, len=%d\n))))))))))))))))))))))))\n",
		resp.Offset, resp.Limit, resp.TotalCount, len(resp.Issues))

	return &resp, nil
}
