package main

import (
	"log"
	"net/http"
	"time"
)

func Sleeper(registryHost string, length time.Duration) {
	log.Printf("sleeping started at registry host %v, for %v\n", registryHost, length)

	log.Fatal(http.ListenAndServe(registryHost, nil))
}
