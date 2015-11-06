package main

import (
	"net/http"
	"log"
	"io"
	"encoding/json"
	"fmt"
	"regexp"
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
	if r.Method=="POST" {
		HandleServicePost(w,r)
	} else if r.Method == "GET" {
		HandleServiceGetList(w,r)
	}
}

func HandleServiceGetList(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, fmt.Sprintf("item item item"))
	log.Println(r.URL)
}

func HandleServiceItem(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if path == "/service/" {
		HandleService(w,r)
		return
	}

	re := regexp.MustCompile("^/service/([0-9]+)$")
	fmt.Printf("%q\n", re.FindStringSubmatch(path))
	m := re.FindStringSubmatch(path)
	io.WriteString(w, fmt.Sprintf("item %d\n", m[1]))

	log.Println(path)
}


func main() {

	// private views
	http.HandleFunc("/service", PostOrGetOnly(HandleService))
	http.HandleFunc("/service/", GetOnly(HandleServiceItem))

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
