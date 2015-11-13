package main

import
(
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
	"net/http"
	"encoding/json"
	"github.com/mpgerlek/piazza-simulator/piazza"
)

func TestGateway(t *testing.T) {

	registryHost := startRegistry(t)
	serviceHost := startGateway(t, registryHost)

	m := NewMessage()
	m.Id = 0 // left empty
	m.Type = CreateJobMessage
	m.User = User{Id: 123, UserType: NormalUser, Name: "Bob"}
	m.Status = 0 // left empty

	params := map[string]any{"count": 5, "compress": true}
	m.CreateJobPayload = CreateJobPayload{ServiceName: "echo", Parameters: params, Comments: "This is the Echo Service. This is the Echo Service."}

	// send job to gateway
	var myHost = fmt.Sprintf("http://%s/service", serviceHost)
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

	stopService(t, serviceHost)
	stopRegistry(t, registryHost)
}
