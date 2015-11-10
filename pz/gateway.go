package main

import (
	"github.com/gorilla/mux"
	"github.com/mpgerlek/piazza-simulator/piazza"
	"io/ioutil"
	"log"
	"net/http"
)

func handleGatewayService(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	body = body
}

func Gateway(registryHost string) error {

	log.Printf("gateway started at registry host %v\n", registryHost)

	id, err := piazza.RegisterService(registryHost, "gateway", "my fun gateway")
	if err != nil {
		return err
	}

	log.Printf("gateway id is %d", id)

	r := mux.NewRouter()
	r.HandleFunc("/service", handleGatewayService).
		Methods("POST")

	server := &http.Server{Addr: registryHost, Handler: r}
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
		return err
	}

	// not reached
	return nil
}
