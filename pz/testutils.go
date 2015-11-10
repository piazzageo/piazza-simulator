package main

import (
	"fmt"
	"time"
)

var currPort = 12300

func getRandomLocalhost() string {
	host := fmt.Sprintf("localhost:%d", currPort)

	currPort++

	return host
}

func startRegistry() string {
	host := getRandomLocalhost()

	go Registry(host)
	time.Sleep(1 * time.Second)

	return host
}

func stopRegistry(host string) {
	// TODO
}
