package registry

import (
	"testing"
	"time"
	"github.com/mpgerlek/piazza-simulator/piazza"
	//"log"
)


func TestRegistry(t *testing.T) {
	var host string
	var proxy *RegistryProxy
	var err error

	// create it, start the proxy
	{
		host = piazza.GetRandomHost()
		go NewRegistryService(host, "my registry")
		time.Sleep(1 * time.Second)

		proxy, err = NewRegistryProxy(host)
		if err != nil {
			t.Error(err)
		}
	}

	// verify only entry is itself
	{
		list, err := proxy.Lookup(piazza.InvalidService)
		if err != nil {
			t.Error(err)
		}

		if list.Len() != 1 {
			t.Errorf("list length %d, expected 1", list.Len())
		}

		service, ok := list.GetOne()
		if !ok {
			t.Error("serivce list is not length 1")
		}
		if service.Type != piazza.RegistryService {
			t.Error("single service is not a RegistryService")
		}
		if service.Id == "" {
			t.Errorf("service id is invalid: %s", service.Id)
		}
		//t.Log(service)
	}

	// add a service
	{
		pid, err := proxy.Register(host, piazza.DispatcherService, "This is my service.")
		if err != nil {
			t.Error(err)
		}
		t.Logf("pid: %v", pid)
	}

	// check what we added
	var addedId piazza.PiazzaId
	{
		list, err := proxy.Lookup(piazza.DispatcherService)
		if err != nil {
			t.Error(err)
		}

		if list.Len() != 1 {
			t.Errorf("list length %d, expected 1", list.Len())
		}

		service, ok := list.GetOne()
		if !ok {
			t.Error("serivce list is not length 1")
		}
		if service.Type != piazza.DispatcherService {
			t.Error("service is not a DispatcherService")
		}
		if service.Comment != "This is my service." {
			t.Error("service comment is \"%s\", expected \"%s\"", service.Comment, "This is my service.")
		}
		addedId = service.Id
	}

	// remove the service
	{
		err := proxy.Remove(addedId)
		if err != nil {
			t.Error(err)
		}
	}

	// verify the table is empty
	{
		/** TODO
		list, err := proxy.Lookup(piazza.InvalidService)
		if err != nil {
			t.Error(err)
		}

		if list.Len() != 0 {
			t.Errorf("list length %d, expected 0", list.Len())
		}
		**/
	}
}
