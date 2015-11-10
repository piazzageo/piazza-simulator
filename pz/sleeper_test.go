package main

import (
	//"bytes"
	//"encoding/json"
	//"errors"
	//"flag"
	//"fmt"
	//"github.com/mpgerlek/piazza-simulator/piazza"
	//"io/ioutil"
	//"net/http"
	//"os"
	"testing"
	//"time"
	//"log"
	"time"
)

func TestSleeper(t *testing.T) {

	var registryHost = startRegistry()

	Sleeper(getRandomLocalhost(), time.Second * 1)

	stopRegistry(registryHost)
}
