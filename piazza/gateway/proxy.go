package dispatcher

import (
	"github.com/mpgerlek/piazza-simulator/piazza"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"log"
	"github.com/mpgerlek/piazza-simulator/piazza/dispatcher"
)

//---------------------------------------------------------------------

type GatewayProxy struct {
	Host string
	DispatcherHost string
}

//---------------------------------------------------------------------

// used by services to connect to the registry
func NewGatewayProxy(serviceHost string, dispatcherProxy *dispatcher.DispatcherProxy) (*GatewayProxy, error) {
	host := fmt.Sprintf("http://%s", serviceHost)

	proxy := &GatewayProxy{Host: host, DispatcherHost: dispatcherProxy.Host}

	return proxy, nil
}


// send a message to the dispatcher
func (proxy *GatewayProxy) Send(m *piazza.Message) (piazza.PiazzaId, error) {
	log.Print("(GatewayProxy.Send)")

	url := fmt.Sprintf("%s/message", proxy.DispatcherHost)

	buf, err := m.ToBytes()
	if err != nil {
		return "", err
	}
	buff := bytes.NewBuffer(buf)

	resp, err := http.Post(url, "application/json", buff)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	buf, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}


	var obj piazza.Message
	err = json.Unmarshal(buf, &obj)
	if err != nil {
		return "", err
	}

	status := obj.CreateJobResponsePayload.JobStatus
	if status != piazza.JobStatusSubmitted {
		return "", errors.New(fmt.Sprintf("bad job status: %v", status))
	}

	pid := obj.CreateJobResponsePayload.JobId
	return pid, nil
}
