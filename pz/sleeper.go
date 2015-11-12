package main

import (
	"github.com/gorilla/mux"
	"github.com/mpgerlek/piazza-simulator/piazza"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func handleSleeperService(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	body = body
}

func Sleeper(serviceHost string, registryHost string, length time.Duration) error {
	log.Printf("sleeping started at registry host %v, for %v\n", registryHost, length)

	id, err := piazza.RegisterService(registryHost, "sleeper", "my fun sleeper")
	if err != nil {
		return err
	}

	log.Printf("sleeper id is %d", id)

	r := mux.NewRouter()
	r.HandleFunc("/service", handleSleeperService).
		Methods("POST")

	server := &http.Server{Addr: serviceHost, Handler: r}
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
		return err
	}

	// not reached
	return nil
}
