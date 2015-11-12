package main

import
(
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
	"net/http"
	"encoding/json"
)

func TestGateway(t *testing.T) {

	registryHost := startRegistry(t)
	serviceHost := startGateway(t, registryHost)

	var payload = `{"user": "mpg", "service": "sleeper", "params": { "a": 1, "b": 2 }}`

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

	var x map[string]interface{}
	err = json.Unmarshal(buf, &x)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(x)

	stopService(t, serviceHost)
	stopRegistry(t, registryHost)
}
