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

func TestEcho(t *testing.T) {
	registryHost := startRegistry(t)
	serviceHost := startSleeper(t, registryHost)

	stopService(t, serviceHost)
	stopRegistry(t, registryHost)
}
