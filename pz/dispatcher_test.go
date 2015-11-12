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
)

func TestDispatcher(t *testing.T) {

	registryHost := startRegistry(t)
	serviceHost := startDispatcher(t, registryHost)

	stopService(t, serviceHost)
	stopRegistry(t, registryHost)
}
