package glideswap

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

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

func ReadPackages(repos []string) (*Packages, error) {
	var err error

	packages := Packages{}

	for _, repo := range repos {
		pkgs, err := readPackagesFromRepo(repo)
		if err != nil {
			return nil, err
		}

		packages.addPackages(pkgs)
	}

	return &packages, err
}

func readPackagesFromRepo(dir string) (*Packages, error) {
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

func (packages *Packages) CheckDependence(swaps *Swaps, dep Dependence) (string, error) {
	//fmt.Printf("Dep: %s\n", dep.Name)
	pkg, ok := (*packages)[dep.Name]
	if !ok {
		return "", fmt.Errorf("internal error: %s not found in package list", dep.Name)
	}

	if strings.Contains(dep.Name, "github.com/venicegeo/") {
		return "", nil
	}

	swap, ok := (*swaps)[dep.Name]
	if !ok {
		s := fmt.Sprintf("%s not found in swap list", dep.Name)
		return s, nil
	}

	if !swap.ContainsSha(pkg.Sha) {
		s := fmt.Sprintf("swap list does not contain needed vesion of %s", dep.Name)
		return s, nil
	}

	return "", nil
}

func (a *Packages) addPackages(b *Packages) {
	for k, vb := range *b {
		va, ok := (*a)[k]
		if !ok {
			(*a)[k] = vb
		} else {
			if !va.isEqual(&vb) {
				log.Fatalf("merged failed: %s\n%#v\n%#v", k, va, vb)
			}
		}
	}
}

func (a *Package) isEqual(b *Package) bool {
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
