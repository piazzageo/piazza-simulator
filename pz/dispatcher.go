package main

import (
	"github.com/mpgerlek/piazza-simulator/piazza"
	"log"
)

func Dispatcher(registryHost string) error {

	log.Printf("dispatcher started at registry host %v\n", registryHost)

	id, err := piazza.RegisterService(registryHost, "dispatcher", "my fun dispatcher")
	if err != nil {
		return err
	}

	log.Printf("dispatcher id is %d", id)

	return nil
}
