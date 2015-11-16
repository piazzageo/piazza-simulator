package dispatcher

import (
	//"bytes"
	//"encoding/json"
	//"errors"
	//"fmt"
	"time"
	"log"
	//"io/ioutil"
	//"net/http"
	//"os"
	"testing"
	"github.com/mpgerlek/piazza-simulator/piazza"
	"github.com/mpgerlek/piazza-simulator/piazza/registry"
	"github.com/mpgerlek/piazza-simulator/piazza/dispatcher"
)

func setup(t *testing.T) *GatewayProxy {
	var err error

	var registryHost = piazza.GetRandomHost()
	var registryProxy *registry.RegistryProxy

	var dispatcherHost = piazza.GetRandomHost()
	var dispatcherProxy *dispatcher.DispatcherProxy

	var gatewayHost = piazza.GetRandomHost()
	var gatewayProxy *GatewayProxy

	{
		go registry.NewRegistryService(registryHost, "my registry")
		time.Sleep(250 * time.Millisecond)

		registryProxy, err = registry.NewRegistryProxy(registryHost)
		if err != nil {
			t.Error(err)
		}
	}
	{
		go dispatcher.NewDispatcherService(dispatcherHost, registryProxy)
		time.Sleep(250 * time.Millisecond)

		dispatcherProxy, err = dispatcher.NewDispatcherProxy(dispatcherHost)
		if err != nil {
			t.Error(err)
		}
	}
	{
		go NewGatewayService(gatewayHost, registryProxy)
		time.Sleep(250 * time.Millisecond)

		gatewayProxy, err = NewGatewayProxy(gatewayHost, dispatcherProxy)
		if err != nil {
			t.Error(err)
		}
	}

	return gatewayProxy
}

func TestGateway(t *testing.T) {

	gatewayProxy := setup(t)

	m := piazza.NewMessage()
	{
		m.Id = piazza.BadId
		m.Type = piazza.CreateJobMessage
		m.User = nil
		m.Status = 0 // left empty

		params := map[string]piazza.Any{"A": "alpha", "B": 222}
		m.CreateJobPayload = &piazza.CreateJobPayload{ServiceType: piazza.EchoService, Parameters: params, Comment: "a run of the echo service"}
	}

	pid, err := gatewayProxy.Send(m)
	if err != nil {
		t.Error(err)
	}
	log.Print(pid)

/*
	// send job to gateway
	var myHost = fmt.Sprintf("http://%s/service", gatewayHost)
	resp, err := http.Post(myHost, "application/json", bytes.NewBufferString(payload))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatal("bonk")
	}

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if buf == nil {
		t.Fatal(err)
	}

	m2, err := piazza.NewMessageFromBytes(buf)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(m2)
	*/
}
