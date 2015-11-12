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
	//"time"
	"time"
)

func TestSleeper(t *testing.T) {
	registryHost := startRegistry(t)
	serviceHost := startSleeper(t, registryHost, 1*time.Second)

	stopService(t, serviceHost)
	stopRegistry(t, registryHost)
}
