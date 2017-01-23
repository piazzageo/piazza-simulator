package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Package stores the description of a single package,
// as read in from glide.lock.
type Dependence struct {
	Name string `yaml:"package"`
}

type Dependences map[string]Dependence

// GlideLock is the struct representing the whole glide.lock file.
type glideYaml struct {
	BasePackage     string       `yaml:"package"`
	Dependences     []Dependence `yaml:"import"`
	TestDependences []Dependence `yaml:"testImport"`
}

func readDependencesFromRepos(repos []string) (*Dependences, error) {
	var err error

	dependences := Dependences{}

	for _, repo := range repos {
		deps, err := readDependences(repo)
		if err != nil {
			return nil, err
		}

		mergeD(&dependences, deps)
	}

	return &dependences, err
}

func readDependences(repo string) (*Dependences, error) {
	buf, err := ioutil.ReadFile(repo + "/glide.yaml")
	if err != nil {
		return nil, err
	}

	data := &glideYaml{}

	err = yaml.Unmarshal(buf, data)
	if err != nil {
		return nil, err
	}

	array := append(data.Dependences, data.TestDependences...)

	deps := Dependences{}
	for _, elem := range array {
		deps[elem.Name] = elem
	}

	return &deps, nil
}
