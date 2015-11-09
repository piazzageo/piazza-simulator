package piazza

import (
	"net/http"
)

// https://gist.github.com/tristanwietsma/8444cf3cb5a1ac496203

/*func HandlePost(w http.ResponseWriter, r *http.Request) {
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
		log.Printf("%v", m)
	}
	io.WriteString(w, "post\n")
}*/

/*
type Result struct {
	FirstName string `json:"first"`
	LastName  string `json:"last"`
}

func HandleJSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result, _ := json.Marshal(Result{"tee", "dub"})
	io.WriteString(w, string(result))
}
*/

func GetOnly(h http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			h(w, r)
			return
		}
		http.Error(w, "get only", http.StatusMethodNotAllowed)
	}
}

func PostOnly(h http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			h(w, r)
			return
		}
		http.Error(w, "post only", http.StatusMethodNotAllowed)
	}
}

func PostOrGetOnly(h http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" || r.Method == "GET" {
			h(w, r)
			return
		}
		http.Error(w, "only POST or GET allowed", http.StatusMethodNotAllowed)
	}
}
