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

func TestGateway(t *testing.T) {

	var registryHost = startRegistry()

	Gateway(getRandomLocalhost())

	stopRegistry(registryHost)
}
