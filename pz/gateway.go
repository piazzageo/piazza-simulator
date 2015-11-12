package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/mpgerlek/piazza-simulator/piazza"
	"io/ioutil"
	"log"
	"net/http"

)


func readUserRequest(r *http.Request) (*userRequest, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Print(err, string(body))
		return nil, err
	}

	var request = new(userRequest)
	err = json.Unmarshal(body, &request)
	if err != nil {
		return nil, err
	}

	return request, nil
}

func forwardUserRequest(dispatcherHost string, request *userRequest) {

	var buf []byte
	json.Unmarshal(buf, request)

	resp, err := http.Post(dispatcherHost, "application/json", buf)

	w.Write([]byte(`{"status":"ok"}`))
}

func Gateway(serviceHost string, registryHost string) error {

	log.Printf("gateway started at registry host %v\n", registryHost)

	id, err := piazza.RegisterService(registryHost, "gateway", "my fun gateway")
	if err != nil {
		return err
	}

	log.Printf("gateway id is %d", id)

	dispatcherUrl, err = piazza.RegistryLookupURL(registryHost, "dispatcher")
	if err != nil {
		return err
	}

	var f = func(w http.ResponseWriter, r *http.Request) {

	}

	r := mux.NewRouter()
	r.HandleFunc("/service", f).Methods("POST")

	server := &http.Server{Addr: serviceHost, Handler: r}
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
		return err
	}

	// not reached
	return nil
}
