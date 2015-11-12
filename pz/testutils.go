package main

import (
	"fmt"
	"testing"
	"time"
)

var currPort = 12300

func getRandomHost() string {
	host := fmt.Sprintf("localhost:%d", currPort)
	currPort++
	return host
}

func startRegistry(t *testing.T) (string) {
	var host = getRandomHost()
	go Registry(host)
	time.Sleep(1 * time.Second)
	return host
}

func startGateway(t *testing.T, registryHost string) (string) {
	var host = getRandomHost()
	go Gateway(host, registryHost)
	time.Sleep(1 * time.Second)
	return host
}

func startDispatcher(t *testing.T, registryHost string) (string) {
	var host = getRandomHost()
	go Dispatcher(host, registryHost)
	time.Sleep(1 * time.Second)
	return host
}

func startSleeper(t *testing.T, registryHost string, duration time.Duration) (string) {
	var host = getRandomHost()
	go Sleeper(host, registryHost, duration)
	time.Sleep(1 * time.Second)
	return host
}

func stopRegistry(t *testing.T, host string) {
	// TODO
}

func stopService(t *testing.T, host string) {
	// TODO
}
