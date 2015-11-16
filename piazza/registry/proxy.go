package registry

import (
	"github.com/mpgerlek/piazza-simulator/piazza"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"log"
)

//---------------------------------------------------------------------

type RegistryProxy struct {
	Host string
}

//---------------------------------------------------------------------

// used by services to connect to the registry
func NewRegistryProxy(host string) (*RegistryProxy, error) {

	proxy := &RegistryProxy{Host: host}

	return proxy, nil
}


// used by services to add a service to the registry, returns the new service's id
func (proxy *RegistryProxy) Register(serviceHost string, serviceType piazza.ServiceType, comment string) (piazza.PiazzaId, error) {
	log.Printf("(Registry.Register) proxyhost=%s host:%s comment=\"%s\"", proxy.Host, serviceHost, comment)

	url := fmt.Sprintf("http://%s/service", proxy.Host)

	service := piazza.Service{Host: serviceHost, Type: serviceType, Comment: comment}

	var badId = piazza.BadId

	buf, err := service.ToBytes()
	if err != nil {
		return badId, err
	}

	var s = string(buf)
	resp, err := http.Post(url, "application/json", bytes.NewBufferString(s))
	if err != nil {
		return badId, err
	}

	defer resp.Body.Close()

	buf, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return badId, err
	}

	var obj map[string]interface{}
	err = json.Unmarshal(buf, &obj)
	if err != nil {
		return badId, err
	}

	value, ok := obj["id"]
	if !ok {
		return badId, errors.New("invalid response from registry server")
	}
	pid, err := piazza.NewIdFromString(value.(string))
	if err != nil {
		return badId, err
	}
	return pid, nil
}

// used by services to add a service to the registry, returns the new service's id
func (proxy *RegistryProxy) Remove(pid piazza.PiazzaId) error {
	return nil // TODO

	/**
	url := fmt.Sprintf("%s/service/%s", proxy.host, pid.String())

	var s = string(buf)
	resp, err := http.Delete(url, "application/json", bytes.NewBufferString(s))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	buf, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return badId, err
	}

	var obj map[string]interface{}
	err = json.Unmarshal(buf, &obj)
	if err != nil {
		return badId, err
	}

	value, ok := obj["id"]
	if !ok {
		return badId, errors.New("invalid response from registry server")
	}
	**/
}

// returns the host name of the service
func (proxy *RegistryProxy) Lookup(stype piazza.ServiceType) (*piazza.ServiceList, error) {
	url := fmt.Sprintf("%s/service/%v", proxy.Host, stype)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, errors.New("buf != nil, expected ==")
	}

	list := new(piazza.ServiceList)
	err = json.Unmarshal(buf, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
