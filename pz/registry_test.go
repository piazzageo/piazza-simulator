package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/mpgerlek/piazza-simulator/piazza"
	"io/ioutil"
	"net/http"
	"testing"
)

func fetchTable(t *testing.T, registryHost string) (*piazza.ServiceTable, error) {
	var myHost = fmt.Sprintf("http://%s/service", registryHost)

	resp, err := http.Get(myHost)
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

	table, err := piazza.NewServiceTableFromBytes(buf)
	if err != nil {
		return nil, err
	}

	return table, nil
}

func fetchEntry(t *testing.T, registryHost string, id int) (*piazza.ServiceEntry, error) {
	var myHost = fmt.Sprintf("http://%s/service/%d", registryHost, id)

	resp, err := http.Get(myHost)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	entry, err := piazza.NewServiceEntryFromBytes(buf)
	if err != nil {
		return nil, err
	}

	return entry, nil
}

func TestRegistry(t *testing.T) {
	var resp *http.Response
	var err error
	var buf []byte

	registryHost := startRegistry(t)

	var testEntryJson1 = `{"name": "svc1", "description": "service one"}`
	var testEntryJson2 = `{"name": "svc2", "description": "service two"}`

	testEntry1, err := piazza.NewServiceEntryFromBytes([]byte(testEntryJson1))
	if err != nil {
		t.Fatal(err)
	}
	testEntry2, err := piazza.NewServiceEntryFromBytes([]byte(testEntryJson2))
	if err != nil {
		t.Fatal(err)
	}

	// table starts out empty
	{
		var table *piazza.ServiceTable

		table, err = fetchTable(t, registryHost)
		if err != nil {
			t.Error(err)
		}

		if table == nil || table.Count() != 0 {
			t.Fatal("initial table not empty")
		}

		var testTable0 = piazza.NewServiceTable()
		if !table.Compare(testTable0) {
			t.Fatal("fetched table incorrect")
		}
	}

	// add a new entry, check response
	{
		var entry *piazza.ServiceEntry

		var myHost = fmt.Sprintf("http://%s/service", registryHost)
		resp, err = http.Post(myHost, "application/json", bytes.NewBufferString(testEntryJson1))
		if err != nil {
			t.Error(err)
		}
		defer resp.Body.Close()

		buf, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Error(err)
		}
		if buf == nil {
			t.Fatal("returned buf empty")
		}

		entry, err = piazza.NewServiceEntryFromBytes(buf)
		if err != nil {
			t.Error(err)
		}

		if !entry.Compare(testEntry1, false) {
			t.Fatal("return entry incorrect")
		}
	}

	// fetch all from DB
	{
		var table *piazza.ServiceTable

		table, err = fetchTable(t, registryHost)
		if err != nil {
			t.Error(err)
		}

		var testTable1 = piazza.NewServiceTable()
		testTable1.Add(testEntry1)
		if !table.Compare(testTable1) {
			t.Fatal("fetched table incorrect")
		}
	}

	// add a 2nd entry
	{
		var myHost = fmt.Sprintf("http://%s/service", registryHost)
		resp, err = http.Post(myHost, "application/json", bytes.NewBufferString(testEntryJson2))
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()

	}

	// fetch all from DB
	{
		var table *piazza.ServiceTable

		table, err = fetchTable(t, registryHost)
		if err != nil {
			t.Error(err)
		}

		var testTable2 = piazza.NewServiceTable()
		testTable2.Add(testEntry1)
		testTable2.Add(testEntry2)
		if !table.Compare(testTable2) {
			t.Fatal("fetched table incorrect")
		}
	}

	stopRegistry(t, registryHost)
}

func TestRegistration(t *testing.T) {

	var err error

	registryHost := startRegistry(t)

	// table starts out empty
	{
		var table *piazza.ServiceTable
		table, err = fetchTable(t, registryHost)
		if err != nil {
			t.Error(err)
		}

		if table == nil || table.Count() != 0 {
			t.Fatal("initial table not empty")
		}
	}

	id, err := piazza.RegisterService(registryHost, "myservice", "my fun service")
	if err != nil {
		t.Error(err)
	}
	if id >= 10 { // reasonable upper bound on entries in this test file
		t.Error("got %v, expected less than 10", id)
	}

	// now retrieve that entry
	entry, err := fetchEntry(t, registryHost, id)
	if err != nil {
		t.Fatal(err)
	}
	if entry.Name != "myservice" {
		t.Fatal("got name \"%s\", expected \"myservice\"")
	}

	stopRegistry(t, registryHost)
}
