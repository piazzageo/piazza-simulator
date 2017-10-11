package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ProjectList is just a list of Projects
type ProjectList struct {
	data map[int]*Project
}

func (list *ProjectList) add(p *Project) {
	list.data[p.ID] = p
}

func (list *ProjectList) getMap() map[int]*Project {
	return list.data
}

func (list *ProjectList) filter(names []string) {

	newList := &ProjectList{
		data: map[int]*Project{},
	}

	for _, project := range list.data {

		for _, name := range names {
			if project.Identifier == name || project.Name == name {
				newList.add(project)
				break
			}
		}

	}

	list.data = newList.data
}

// NewProjectList makes a new ProjectList
func NewProjectList(apiKey string) (*ProjectList, error) {

	list := &ProjectList{
		data: map[int]*Project{},
	}

	type response struct {
		Projects   []*Project `json:"projects"`
		TotalCount int        `json:"total_count"`
		Offset     int        `json:"offset"`
		Limit      int        `json:"limit"`
	}

	const offset = 0
	const limit = 100

	url := fmt.Sprintf("%s/%s?offset=%d&limit=%d&status_id=*",
		"https://redmine.devops.geointservices.io",
		"/projects.json",
		offset, limit)

	client := &http.Client{}

	Logf("url: %s", url)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	request.SetBasicAuth(apiKey, "random")

	resp1, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	if resp1.StatusCode != 200 {
		return nil, fmt.Errorf("GET failed with status code %s", resp1.Status)
	}

	body, err := ioutil.ReadAll(resp1.Body)
	if err != nil {
		return nil, err
	}

	Logf("Raw result: %s...", string(body[0:60]))

	resp2 := response{}
	err = json.Unmarshal(body, &resp2)
	if err != nil {
		return nil, err
	}
	Logf("request: offset=%d, limit=%d, totalCount=%d, len=%d",
		resp2.Offset, resp2.Limit, resp2.TotalCount, len(resp2.Projects))

	for _, p := range resp2.Projects {
		list.add(p)
	}
	return list, nil
}
