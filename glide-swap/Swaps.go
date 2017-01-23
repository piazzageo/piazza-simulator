package main

import (
	"encoding/csv"
	"os"
)

type Swap struct {
	Name    string
	Sha     string
	Version string
}

type Swaps map[string]Swap

func readSwaps(swapFile string) (*Swaps, error) {
	file, err := os.Open(swapFile)
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(file)
	reader.Comma = ','
	reader.FieldsPerRecord = 3
	reader.Comment = '#'
	reader.TrimLeadingSpace = true

	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	swaps := Swaps{}
	for _, elem := range data {
		swap := Swap{
			Name:    elem[0],
			Sha:     elem[1],
			Version: elem[2],
		}
		swaps[elem[0]] = swap
	}

	//	log.Printf("%v", data)

	return &swaps, nil
}
