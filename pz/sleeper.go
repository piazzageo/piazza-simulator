package main

import (
	"log"
	"net/http"
	//	"io"
	//	"encoding/json"
	//	"fmt"
	//	"regexp"
	"strconv"

	//	"github.com/mpgerlek/piazza-simulator/piazza"
	"time"
)

func Sleeper(registryPort int, length time.Duration) {
	log.Printf("sleeping started at registry port %d, for %v\n", registryPort, length)
	log.Fatal(http.ListenAndServe("localhost:"+strconv.Itoa(registryPort), nil))
}
