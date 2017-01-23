package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// stores the description of a single package,
// as read in from glide.lock

type GlidePackage struct {
	Hash    string
	Updated string
	Imports map[string]interface{}
}

func readGlideLock() ([]GlidePackage, error) {
	buf, err := ioutil.ReadFile("./glide.lock")
	if err != nil {
		return nil, err
	}

	data := &GlidePackage{}

	err = yaml.Unmarshal(buf, data)
	if err != nil {
		return nil, err
	}

	log.Printf("%#v", data)
	return nil, nil
}
