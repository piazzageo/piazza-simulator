package main

import (
	"bytes"
	"github.com/gorilla/mux"
	"github.com/mpgerlek/piazza-simulator/piazza"
	"io/ioutil"
	"log"
	"net/http"
)

func readUserRequest(r *http.Request) (*piazza.Message, error) {
	log.Printf("reading user request")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Print(err, string(body))
		return nil, err
	}

	m, err := piazza.NewMessageFromBytes(body)
	if err != nil {
		return nil, err
	}

	log.Printf("returning request response")
	return m, nil
}

func forwardUserRequest(dispatcherHost string, m *piazza.Message) (*piazza.Message, error) {
	log.Printf("forwarding user request")

	buf, err := m.ToBytes()
	var buff = bytes.NewBuffer(buf)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(dispatcherHost, "application/json", buff)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// resp has a message of type CreateJobResponsePayload, with the id filled in
	buff = new(bytes.Buffer)
	buff.ReadFrom(resp.Body)

	m2, err := piazza.NewMessageFromBytes(buff.Bytes())
	if err != nil {
		return nil, err
	}
	log.Printf("returning request response")

	return m2, nil
}

func Gateway(serviceHost string, registryHost string) error {

	log.Printf("gateway started at registry host %v\n", registryHost)

	id, err := piazza.RegisterService(registryHost, piazza.GatewayService, "my fun gateway")
	if err != nil {
		return err
	}

	log.Printf("gateway id is %d", id)

	dispatcherHost, err := RegistryLookup(registryHost, piazza.DispatcherService)
	if err != nil {
		return err
	}

	var handleRequest = func(w http.ResponseWriter, r *http.Request) {
		m, err := readUserRequest(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		m2, err := forwardUserRequest(dispatcherHost, m)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		// we now have the response to return to the outside caller
		buf, err := m2.ToBytes()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		w.Write(buf)
	}

	r := mux.NewRouter()
	r.HandleFunc("/service", handleRequest).Methods("POST")

	server := &http.Server{Addr: serviceHost, Handler: r}
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
		return err
	}

	// not reached
	return nil
}
