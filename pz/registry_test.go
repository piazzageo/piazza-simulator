package main

import (
	"bytes"
	//"encoding/json"
	"errors"
	"flag"
	"github.com/mpgerlek/piazza-simulator/piazza"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
	"time"
)



func fetchTable() (*piazza.ServiceTable, error) {
	resp, err := http.Get("http://localhost:8080/service")
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


func TestRegistry(t *testing.T) {
	go Registry(8080)
	time.Sleep(1 * time.Second)

	var resp *http.Response
	var err error
	var buf []byte

	var testEntryJson1 = `{"name": "no-op", "description": "some text"}`
	var testEntry1 = &piazza.ServiceEntry{Name: "no-op", Description: "some text"}
	var testEntryJson2 = `{"name": "foo", "description": "bar"}`
	var testEntry2 = &piazza.ServiceEntry{Name: "foo", Description: "bar"}

	// table starts out empty
	{
		var table *piazza.ServiceTable

		table, err = fetchTable()
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

		resp, err = http.Post("http://localhost:8080/service", "application/json", bytes.NewBufferString(testEntryJson1))
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

		if !entry.Compare(testEntry1) {
			t.Fatal("return entry incorrect")
		}
	}

	// fetch all from DB
	{
		var table *piazza.ServiceTable

		table, err = fetchTable()
		if err != nil {
			t.Error(err)
		}

		var testTable1 = piazza.NewServiceTable()
		testTable1.Add(testEntry1)
		if !table.Compare(testTable1) {
			t.Log(table)
			t.Log(testTable1)

			t.Fatal("fetched table incorrect")
		}
	}

	// add a 2nd entry
	{
		resp, err = http.Post("http://localhost:8080/service", "application/json", bytes.NewBufferString(testEntryJson2))
		if err != nil {
			t.Error(err)
		}
		defer resp.Body.Close()

	}

	// fetch from DB
	{
		var table *piazza.ServiceTable

		table, err = fetchTable()
		if err != nil {
			t.Error(err)
		}

		var testTable2 = piazza.NewServiceTable()
		testTable2.Add(testEntry1)
		testTable2.Add(testEntry2)
		if !table.Compare(testTable2) {
			t.Log(table)
			t.Log(testTable2)
			t.Fatal("fetched table incorrect")
		}
	}
}

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}
