package main

import (
	"github.com/gorilla/mux"
	"github.com/mpgerlek/piazza-simulator/piazza"
	"io/ioutil"
	"log"
	"net/http"
)

func handleEchoService(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	body = body
}

func Echo(serviceHost string, registryHost string) error {
	log.Printf("echo started at registry host %s\n", registryHost)

	id, err := piazza.RegisterService(registryHost, piazza.EchoService, "my fun echo")
	if err != nil {
		return err
	}

	log.Printf("echo id is %d", id)

	r := mux.NewRouter()
	r.HandleFunc("/service", handleEchoService).
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
