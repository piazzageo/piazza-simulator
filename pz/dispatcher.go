package main

import (
	"github.com/mpgerlek/piazza-simulator/piazza"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func handleDispatcherRequest(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Print(err, string(body))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	m, err := piazza.NewMessageFromBytes(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Print(m)

	w.Write([]byte(`{"status":"ok"}`))
}

func Dispatcher(serviceHost string, registryHost string) error {

	log.Printf("dispatcher started at registry host %v\n", registryHost)

	id, err := piazza.RegisterService(serviceHost, piazza.DispatcherService, "my fun dispatcher")
	if err != nil {
		return err
	}

	log.Printf("dispatcher id is %d", id)

	r := mux.NewRouter()
	r.HandleFunc("/service", handleDispatcherRequest).
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
