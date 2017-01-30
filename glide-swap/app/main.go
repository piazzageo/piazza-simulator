package main

import (
	"flag"
	"fmt"
	"log"

	glideswap "github.com/venicegeo/mpg-sandbox/glide-swap/glideswap"
)

var swapFile string

func init() {
	flag.StringVar(&swapFile, "swap", "./swap.txt", "location of swap file")
}

var packages *glideswap.Packages
var swaps *glideswap.Swaps
var dependences *glideswap.Dependences

func main() {
	var err error

	flag.Parse()
	repoDirs := flag.Args()

	packages, swaps, dependences, err = readAll(swapFile, repoDirs)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	ss, err := dependences.CheckDependences(packages, swaps)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	for _, s := range ss {
		fmt.Printf("*** %s\n", s)
	}

	fmt.Printf("Done.\n")
}

func readAll(swapFile string, repoDirs []string) (*glideswap.Packages, *glideswap.Swaps, *glideswap.Dependences, error) {
	packages, err := glideswap.ReadPackages(repoDirs)
	if err != nil {
		return nil, nil, nil, err
	}

	swaps, err := glideswap.ReadSwaps(swapFile)
	if err != nil {
		return nil, nil, nil, err
	}

	dependences, err := glideswap.ReadDependences(repoDirs)
	if err != nil {
		return nil, nil, nil, err
	}

	return packages, swaps, dependences, nil
}
