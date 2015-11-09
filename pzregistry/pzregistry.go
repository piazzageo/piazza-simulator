package main

import (
	"net/http"
	"log"
	"io"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
)

type entryId int64

type serviceEntry struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Id          entryId  `json:"id"`
}

func (e serviceEntry) String() string {
	return fmt.Sprintf("[name:%s, description:%s, id:%d]", e.Name, e.Description, e.Id)
}

type serviceTable struct {
	table     map[entryId]*serviceEntry
	currentId entryId
}

var table serviceTable


func (t *serviceTable) add(entry *serviceEntry) {
	if t.table == nil {
		t.table = make(map[entryId]*serviceEntry)
	}

	entry.Id = t.currentId
	t.currentId++

	t.table[entry.Id] = entry
}

func (t *serviceTable) get(id entryId) (entry *serviceEntry, ok bool) {
	entry, ok = t.table[id]
	return entry, ok
}


func HandleServicePost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if len(r.PostForm) != 1 {
		http.Error(w, "post only", http.StatusMethodNotAllowed)
		return
	}

	for k, _ := range r.PostForm {

		b := []byte(k)

		var m serviceEntry
		err := json.Unmarshal(b, &m)
		if err != nil {
			http.Error(w, "unmarshall failed", http.StatusMethodNotAllowed)
			return
		}

		table.add(&m)
		log.Printf("added entry [%d, %s]", m.Id, m.Name)

		io.WriteString(w, fmt.Sprintf("added id:%d\n", m.Id))
	}
}


func HandleService(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		HandleServicePost(w, r)
	} else if r.Method == "GET" {
		HandleServiceGet(w, r)
	} else {
		http.Error(w, "only GET and POST allowed", http.StatusMethodNotAllowed)
	}
}


func HandleServiceGet(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, fmt.Sprintf("item item item"))
	log.Println(r.URL)
}

func HandleServiceItem(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		HandleServiceItemGet(w, r)
	} else {
		http.Error(w, "only GET and POST allowed", http.StatusMethodNotAllowed)
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

	var v int64
	if v, err = strconv.ParseInt(ms[1], 10, 64); err != nil {
		http.Error(w, "internal atoi failure", http.StatusMethodNotAllowed)
		log.Printf("internal atoi failure, %q", ms[1])
	}
	var id entryId = entryId(v)
	var entry *serviceEntry

	if entry, ok = table.get(id); !ok {
		http.Error(w, "internal match failure", http.StatusMethodNotAllowed)
		log.Printf("internal item match failure, path:%q", ms)
	}

	w.Header().Set("Content-Type", "application/json")
	//var b []byte
	b, err := json.Marshal(entry)
	io.WriteString(w, string(b)+"\n")
}


func main() {

	// TODO: /service/00
	// TODO: /service/-1

	http.HandleFunc("/service", PostOrGetOnly(HandleService))
	http.HandleFunc("/service/", GetOnly(HandleServiceItem))

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
