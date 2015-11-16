package dispatcher

import (
	//"bytes"
	//"encoding/json"
	//"errors"
	//"flag"
	//"fmt"
	//"io/ioutil"
	//"net/http"
	//"os"
	"testing"
	"time"
	"log"
	"github.com/mpgerlek/piazza-simulator/piazza"
	"github.com/mpgerlek/piazza-simulator/piazza/registry"
)

func TestDispatcher(t *testing.T) {

	var registryProxy *registry.RegistryProxy
	var dispatcherHost string
	var dispatcherProxy *DispatcherProxy
	var err error

	// start up the registry
	{
		host := piazza.GetRandomHost()
		go registry.NewRegistryService(host, "my registry")
		time.Sleep(1 * time.Second)

		registryProxy, err = registry.NewRegistryProxy(host)
		if err != nil {
			t.Error(err)
		}
	}

	// start the dispatcher
	{
		dispatcherHost = piazza.GetRandomHost()
		go NewDispatcherService(dispatcherHost, registryProxy)
		time.Sleep(1 * time.Second)
	}

	// start the dispatcher proxy
	{
		dispatcherProxy, err = NewDispatcherProxy(dispatcherHost)
		if err != nil {
			t.Error(err)
		}
		log.Printf("%v", dispatcherProxy)
	}
}
