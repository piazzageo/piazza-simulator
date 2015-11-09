package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	//"bytes"
	//"errors"

	"github.com/mpgerlek/piazza-simulator/piazza"
)

var table *piazza.ServiceTable

func HandleServicePost(w http.ResponseWriter, r *http.Request) {
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

func HandleService(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		HandleServicePost(w, r)
	} else if r.Method == "GET" {
		HandleServiceGet(w, r)
	} else {
		http.Error(w, "1 only GET and POST allowed", http.StatusMethodNotAllowed)
	}
}

func HandleServiceGet(w http.ResponseWriter, r *http.Request) {
	var buf []byte
	buf, err := table.ToBytes()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Print("sending back")
	log.Print(buf)
	w.Write(buf)
}

func HandleServiceItem(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		HandleServiceItemGet(w, r)
	} else {
		http.Error(w, "2 only GET allowed", http.StatusMethodNotAllowed)
	}
}

func HandleServiceItemGet(w http.ResponseWriter, r *http.Request) {
	var err error
	var ok bool

	path := r.URL.Path
	if path == "/service/" {
		r.URL.Path = "/service"
		HandleServiceGet(w, r)
		return
	}

	re := regexp.MustCompile("^/service/([0-9]+)$")
	ms := re.FindStringSubmatch(path)
	log.Printf("url matches: %q\n", ms)
	if len(ms) != 2 {
		http.Error(w, "internal match failure", http.StatusMethodNotAllowed)
		log.Printf("internal item match failure, path:%q", ms)
	}

	if _, err = strconv.ParseInt(ms[1], 10, 64); err != nil {
		http.Error(w, "internal atoi failure", http.StatusMethodNotAllowed)
		log.Printf("internal atoi failure, %q", ms[1])
	}
	id := ms[1]
	var entry *piazza.ServiceEntry

	if entry, ok = table.Get(id); !ok {
		http.Error(w, "internal match failure", http.StatusMethodNotAllowed)
		log.Printf("internal item match failure, path:%q", ms)
	}

	w.Header().Set("Content-Type", "application/json")
	//var b []byte
	b, err := json.Marshal(entry)
	io.WriteString(w, string(b)+"\n")
}

func Registry(portNumber int) {

	table = piazza.NewServiceTable()

	log.Printf("registry started on port %d", portNumber)
	// TODO: /service/00
	// TODO: /service/-1
	// TODO: DELETE

	http.HandleFunc("/service", HandleService)
	http.HandleFunc("/service/", HandleServiceItem)

	log.Fatal(http.ListenAndServe("localhost:"+strconv.Itoa(portNumber), nil))
}
