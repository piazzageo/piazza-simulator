package glideswap

import (
	"io/ioutil"
	"log"

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

func ReadDependences(repos []string) (*Dependences, error) {
	var err error

	dependences := Dependences{}

	for _, repo := range repos {
		deps, err := readDependencesFromRepo(repo)
		if err != nil {
			return nil, err
		}

		dependences.addDependences(deps)
	}

	return &dependences, err
}

func readDependencesFromRepo(repo string) (*Dependences, error) {
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

func (dependences *Dependences) CheckDependences(packages *Packages, swaps *Swaps) ([]string, error) {
	ss := []string{}

	for _, dep := range *dependences {
		s, err := packages.CheckDependence(swaps, dep)
		if err != nil {
			return nil, err
		}
		if s != "" {
			ss = append(ss, s)
		}
	}

	return ss, nil
}

func (a *Dependences) addDependences(b *Dependences) {
	for k, vb := range *b {
		va, ok := (*a)[k]
		if !ok {
			(*a)[k] = vb
		} else {
			if va != vb {
				log.Fatalf("merged failed: %s\n%#v\n%#v", k, va, vb)
			}
		}
	}
}
