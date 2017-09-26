package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Project struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Identifier string `json:"identifier"`
}

type ProjectList struct {
	data map[int]*Project
}

type ProjectsResponse struct {
	Projects   []*Project `json:"projects"`
	TotalCount int        `json:"total_count"`
	Offset     int        `json:"offset"`
	Limit      int        `json:"limit"`
}

func NewProjectList() *ProjectList {
	p := &ProjectList{
		data: map[int]*Project{},
	}
	return p
}

func (list *ProjectList) Add(p *Project) {
	list.data[p.Id] = p
}

func (list *ProjectList) GetMap() map[int]*Project {
	return list.data
}

func (list *ProjectList) Filter(names []string) *ProjectList {

	newList := NewProjectList()

	for _, project := range list.data {

		for _, name := range names {
			if project.Identifier == name || project.Name == name {
				newList.Add(project)
				break
			}
		}

	}

	return newList
}

func getProjects(apiKey string) (*ProjectList, error) {

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

	Logf("Raw result: %s...", string(body[0:60]))

	resp := ProjectsResponse{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	Logf("request: offset=%d, limit=%d, totalCount=%d, len=%d",
		resp.Offset, resp.Limit, resp.TotalCount, len(resp.Projects))

	list := NewProjectList()
	for _, p := range resp.Projects {
		list.Add(p)
	}
	return list, nil
}
