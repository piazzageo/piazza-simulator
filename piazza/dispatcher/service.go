package dispatcher

import (
	"github.com/mpgerlek/piazza-simulator/piazza"
	"github.com/mpgerlek/piazza-simulator/piazza/registry"
	//"io/ioutil"
	"log"
	//"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)

type Dispatcher struct {
	id piazza.PiazzaId
}

func (dispatcher *Dispatcher) postService(w http.ResponseWriter, r *http.Request) {

	/*body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mIn, err := piazza.NewMessageFromBytes(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Print(mIn)*/

	mOut := piazza.NewMessage()
	{
		mOut.Id = piazza.BadId // left empty
		mOut.Type = piazza.CreateJobResponseMessage
		mOut.Status = piazza.SuccessStatus

		mOut.CreateJobResponsePayload = &piazza.CreateJobResponsePayload{piazza.NewId(), piazza.JobStatusSubmitted}
	}

	buf, err := json.Marshal(mOut)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write(buf)
}

func NewDispatcherService(serviceHost string, registryProxy *registry.RegistryProxy) error {

	log.Printf("dispatcher started: %s (proxy: %s)", serviceHost, registryProxy.Host)

	id, err := registryProxy.Register(serviceHost, piazza.DispatcherService, "my fun dispatcher")

	dispatcher := new(Dispatcher)
	dispatcher.id = id

	log.Printf("dispatcher id: %s", id)

	r := mux.NewRouter()
	r.HandleFunc("/message", dispatcher.postService).Methods("POST")

	server := &http.Server{Addr: serviceHost, Handler: r}
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
		return err
	}

	// not reached
	return nil
}
