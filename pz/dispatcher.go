package main

import (
	"github.com/mpgerlek/piazza-simulator/piazza"
	"log"
)


func handleDispatcherRequest(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Print(err, string(body))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var request = new(userRequest)
	err = json.Unmarshal(body, &request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Print(request)

	w.Write([]byte(`{"status":"ok"}`))
}

func Dispatcher(serviceHost string, registryHost string) error {

	log.Printf("dispatcher started at registry host %v\n", registryHost)

	id, err := piazza.RegisterService(serviceHost, "dispatcher", "my fun dispatcher")
	if err != nil {
		return err
	}

	log.Printf("dispatcher id is %d", id)

	r := mux.NewRouter()
	r.HandleFunc("/service", handleGatewayService).
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
