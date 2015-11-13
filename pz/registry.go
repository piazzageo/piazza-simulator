package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"fmt"
	"errors"
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

	entry, err := piazza.NewServiceFromBytes(body)
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

	log.Printf("added entry [%s, %s], now cnt:%d", entry.Id, entry.Type, table.Count())
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

func handleRegistry_GetServiceItem(w http.ResponseWriter, r *http.Request) {
	var err error
	var ok bool

	key := mux.Vars(r)["key"]

	var entry *piazza.Service

	if id, err := strconv.ParseInt(key, 10, 64); err == nil {
		if entry, ok = table.LookupById(piazza.ServiceId(id)); !ok {
			http.Error(w, "id not found", http.StatusBadRequest)
			return
		}
	} else {
		stype := piazza.ServiceTypeFromString(key)
		if stype == piazza.InvalidService {
			http.Error(w, "service type not found", http.StatusBadRequest)
			return
		}
		if entry, ok = table.LookupByType(stype); !ok {
			http.Error(w, "service type not found", http.StatusBadRequest)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	buf, err := entry.ToBytes()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write(buf)
}

func Registry(registryHost string) error {

	table = piazza.NewServiceTable()

	log.Printf("registry started on %s", registryHost)
	// TODO: /service/00
	// TODO: /service/-1
	// TODO: DELETE

	r := mux.NewRouter()
	r.HandleFunc("/service", handleRegistryService)
	r.HandleFunc("/service/{key}", handleRegistry_GetServiceItem).Methods("GET")

	server := &http.Server{Addr: registryHost, Handler: r}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
		return err
	}

	// not reached
	return nil
}


func RegistryLookup(registryHost string, stype piazza.ServiceType) (value string, err error) {
	entry, ok := table.LookupByType(stype)
	if !ok {
		return "", errors.New(fmt.Sprintf("%v not found in registry", stype))
	}

	return entry.Host, nil
}