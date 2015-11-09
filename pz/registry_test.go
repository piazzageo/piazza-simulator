package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
	"time"
)

func unpackTable(buf []byte) (*serviceTable, error) {
	t := NewServiceTable()

	err := json.Unmarshal(buf, &t.Table)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func unpackEntry(buf []byte) (*serviceEntry, error) {
	var e serviceEntry

	err := json.Unmarshal(buf, &e)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func compareEntries(a *serviceEntry, b *serviceEntry) bool {
	return a.Name == b.Name &&
		a.Description == b.Description &&
		a.Ids == b.Ids
}

func compareTables(a *serviceTable, b *serviceTable) bool {
	if len(a.Table) != len(b.Table) {
		return false
	}
	for k, _ := range a.Table {
		if compareEntries(a.Table[k], b.Table[k]) == false {
			return false
		}
	}
	return true
}

func fetchTable() (*serviceTable, error) {
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

	table, err := unpackTable(buf)
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
	var testEntry1 = &serviceEntry{Name: "no-op", Description: "some text"}
	var testEntryJson2 = `{"name": "foo", "description": "bar"}`
	var testEntry2 = &serviceEntry{Name: "foo", Description: "bar"}

	var testTable *serviceTable = NewServiceTable()

	// table starts out empty
	{
		var table *serviceTable

		table, err = fetchTable()
		if err != nil {
			t.Error(err)
		}

		if table == nil || len(table.Table) != 0 {
			t.Fatal("initial table not empty")
		}
	}

	err = testTable.add(testEntry1)
	if err != nil {
		t.Error(err)
	}


	// add a new entry, check response
	{
		var entry *serviceEntry

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

		entry, err = unpackEntry(buf)
		if err != nil {
			t.Error(err)
		}

		if !compareEntries(entry, testEntry1) {
			t.Fatal("return entry incorrect")
		}
	}

	// fetch from DB
	{
		resp, err = http.Get("http://localhost:8080/service/")
		if err != nil {
			t.Error(err)
		}
		defer resp.Body.Close()

		buf, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Error(err)
		}

		table, err := unpackTable(buf)
		if err != nil {
			t.Error(err)
		}

		if !compareTables(testTable, table) {
			t.Error("fetched table incorrect")
		}
	}

	err = testTable.add(testEntry2)
	if err != nil {
		t.Error(err)
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
		resp, err = http.Get("http://localhost:8080/service/")
		if err != nil {
			t.Error(err)
		}
		defer resp.Body.Close()

		buf, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Error(err)
		}

		table, err := unpackTable(buf)
		if err != nil {
			t.Error(err)
		}

		if !compareTables(testTable, table) {
			t.Error("fetched table incorrect")
		}
	}}

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}
