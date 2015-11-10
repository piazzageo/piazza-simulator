package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	//"bytes"
	//"errors"
	"github.com/gorilla/mux"
	"github.com/mpgerlek/piazza-simulator/piazza"
)

var table *piazza.ServiceTable

func handleRegistryServicePost(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	entry, err := piazza.NewServiceEntryFromBytes(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err2 := table.Add(entry)
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusBadRequest)
		return
	}

	var bytes []byte
	bytes, err = entry.ToBytes()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(bytes)

	log.Printf("added entry [%s, %s], now cnt:%d", entry.Id, entry.Name, table.Count())
}

func handleRegistryService(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		handleRegistryServicePost(w, r)
	} else if r.Method == "GET" {
		handleRegistryServiceGet(w, r)
	} else {
		http.Error(w, "1 only GET and POST allowed", http.StatusMethodNotAllowed)
	}
}

func handleRegistryServiceGet(w http.ResponseWriter, r *http.Request) {
	var buf []byte
	buf, err := table.ToBytes()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(buf)
}

func handleRegistryServiceItem(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		handleRegistryServiceItemGet(w, r)
	} else {
		http.Error(w, "2 only GET allowed", http.StatusMethodNotAllowed)
	}
}

func handleRegistryServiceItemGet(w http.ResponseWriter, r *http.Request) {
	var err error
	var ok bool

	path := r.URL.Path
	if path == "/service/" {
		r.URL.Path = "/service"
		handleRegistryServiceGet(w, r)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	if _, err = strconv.ParseInt(id, 10, 64); err != nil {
		http.Error(w, "internal atoi failure", http.StatusMethodNotAllowed)
	}

	var entry *piazza.ServiceEntry

	if entry, ok = table.Get(id); !ok {
		http.Error(w, "internal match failure", http.StatusMethodNotAllowed)
	}

	w.Header().Set("Content-Type", "application/json")
	//var b []byte
	b, err := json.Marshal(entry)
	io.WriteString(w, string(b)+"\n")
}

func Registry(registryHost string) error {

	table = piazza.NewServiceTable()

	log.Printf("registry started on %s", registryHost)
	// TODO: /service/00
	// TODO: /service/-1
	// TODO: DELETE

	r := mux.NewRouter()
	r.HandleFunc("/service", handleRegistryService)
	r.HandleFunc("/service/{id}", handleRegistryServiceItem)

	server := &http.Server{Addr: registryHost, Handler: r}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
		return err
	}

	// not reached
	return nil
}
