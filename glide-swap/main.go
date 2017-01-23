package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

var swapFile string

func init() {
	flag.StringVar(&swapFile, "swap", "./swap.txt", "location of swap file")
}

func main() {
	flag.Parse()
	repoDirs := flag.Args()

	packages, err := readPackagesFromRepos(repoDirs)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	//log.Printf("%#v", packages)

	swaps, err := readSwaps(swapFile)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	//log.Printf("%#v", swaps)

	dependences, err := readDependencesFromRepos(repoDirs)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	//log.Printf("%#v", dependences)

	for _, dep := range *dependences {
		//fmt.Printf("Dep: %s\n", dep.Name)
		dep, ok := (*packages)[dep.Name]
		if !ok {
			log.Fatalf("internal error: %s not found in package list", dep.Name)
		}

		if strings.Contains(dep.Name, "github.com/venicegeo/") {
			continue
		}

		swap, ok := (*swaps)[dep.Name]
		if !ok {
			fmt.Printf("TASK: %s not found in swap list\n", dep.Name)
			continue
		}

		if !swap.containsSha(dep.Sha) {
			fmt.Printf("TASK: swap list does not contain needed vesion of %s\n", dep.Name)
		}
	}
}

func mergeD(a *Dependences, b *Dependences) {
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

func mergeP(a *Packages, b *Packages) {
	for k, vb := range *b {
		va, ok := (*a)[k]
		if !ok {
			(*a)[k] = vb
		} else {
			if !va.equal(&vb) {
				log.Fatalf("merged failed: %s\n%#v\n%#v", k, va, vb)
			}
		}
	}
}

func contains(a []string, b string) bool {
	for _, s := range a {
		if b == s {
			return true
		}
	}
	return false
}
