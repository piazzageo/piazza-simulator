package glideswap

import (
	"encoding/csv"
	"log"
	"os"
)

type Swap struct {
	Name string
	Ver  map[string]string // sha string to version number
}

type Swaps map[string]Swap

func ReadSwaps(swapFile string) (*Swaps, error) {
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
		nam := elem[0]
		_, ok := swaps[nam]
		if ok {
			// already have it
			swaps[nam].Ver[elem[1]] = elem[2]
		} else {
			swaps[nam] = Swap{
				Name: elem[0],
				Ver:  map[string]string{elem[1]: elem[2]},
			}
		}
	}

	//	log.Printf("%v", data)

	return &swaps, nil
}

func (swap *Swap) ContainsSha(sha string) bool {
	_, ok := swap.Ver[sha]
	return ok
}

func (swap *Swap) addVer(newsha string, newver string) {
	curver, ok := swap.Ver[newsha]
	if ok {
		if curver != newver {
			log.Fatalf("internal error: version mismatch for %s, sha %s", swap.Name, newsha)
		}
		return
	}

	swap.Ver[newsha] = newver
}
