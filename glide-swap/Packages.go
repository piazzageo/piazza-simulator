package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Package stores the description of a single package,
// as read in from glide.lock.
type Package struct {
	Name        string
	Path        string
	Sha         string `yaml:"version"`
	Subpackages []string
}

type Packages map[string]Package

// glideLock is the struct representing the whole glide.lock file.
type glideLock struct {
	Hash         string
	Updated      string
	Packages     []Package `yaml:"imports"`
	TestPackages []Package `yaml:"testImports"`
}

func (a *Package) equal(b *Package) bool {
	if a.Name != b.Name ||
		a.Path != b.Path ||
		a.Sha != b.Sha {
		return false
	}

	if len(a.Subpackages) != len(b.Subpackages) {
		return false
	}
	for i := 0; i < len(a.Subpackages); i++ {
		// TODO: need to compare in sorted order
		if a.Subpackages[i] != b.Subpackages[i] {
			return false
		}
	}

	return true
}

func readPackagesFromRepos(repos []string) (*Packages, error) {
	var err error

	packages := Packages{}

	for _, repo := range repos {
		pkgs, err := readPackages(repo)
		if err != nil {
			return nil, err
		}

		mergeP(&packages, pkgs)
	}

	return &packages, err
}

func readPackages(dir string) (*Packages, error) {
	buf, err := ioutil.ReadFile(dir + "/glide.lock")
	if err != nil {
		return nil, err
	}

	data := &glideLock{}

	err = yaml.Unmarshal(buf, data)
	if err != nil {
		return nil, err
	}

	array := append(data.Packages, data.TestPackages...)

	pkgs := Packages{}
	for _, elem := range array {
		pkgs[elem.Name] = elem
	}

	return &pkgs, nil
}
