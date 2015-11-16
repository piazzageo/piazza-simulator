package dispatcher

import (
	"github.com/gorilla/mux"
	"github.com/mpgerlek/piazza-simulator/piazza"
	"github.com/mpgerlek/piazza-simulator/piazza/registry"
	"io/ioutil"
	"log"
	"net/http"
	"bytes"
	"github.com/mpgerlek/piazza-simulator/piazza/dispatcher"
	"errors"
)

type Gateway struct {
	id piazza.PiazzaId
	registryProxy *registry.RegistryProxy
	dispatcherProxy *dispatcher.DispatcherProxy
}

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


func (gateway *Gateway) postMessage(w http.ResponseWriter, r *http.Request) {
	m, err := readUserRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	m2, err := forwardUserRequest(gateway.dispatcherProxy.Host, m)
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


func NewGatewayService(serviceHost string, registryProxy *registry.RegistryProxy) error {
	var err error
	gateway := new(Gateway)

	gateway.registryProxy = registryProxy
	if err != nil {
		return err
	}

	gateway.id, err = gateway.registryProxy.Register(serviceHost, piazza.GatewayService, "my fun gateway")
	if err != nil {
		log.Fatalf("(Gateway.NewGatewayService) %v", err)
		return err
	}

	log.Printf("gateway id is %d", gateway.id)

	dispatcherServiceList, err := gateway.registryProxy.Lookup(piazza.DispatcherService)
	if err != nil {
		return err
	}

	dispatcherService, ok := dispatcherServiceList.GetOne()
	if !ok {
		return errors.New("dispatcher service not registered")
	}

	gateway.dispatcherProxy, err = dispatcher.NewDispatcherProxy(dispatcherService.Host)
	if err != nil {
		return err
	}

	r := mux.NewRouter()
	r.HandleFunc("/service", gateway.postMessage).Methods("POST")

	server := &http.Server{Addr: serviceHost, Handler: r}
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
		return err
	}

	// not reached
	return nil
}
