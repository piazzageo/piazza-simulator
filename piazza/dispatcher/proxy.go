package dispatcher

import (
	//"github.com/mpgerlek/piazza-simulator/piazza"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	//"log"
)

//---------------------------------------------------------------------

type DispatcherProxy struct {
	Host string
}

//---------------------------------------------------------------------

// used by services to connect to the registry
func NewDispatcherProxy(serviceHost string) (*DispatcherProxy, error) {
	host := fmt.Sprintf("http://%s", serviceHost)

	proxy := &DispatcherProxy{Host: host}

	return proxy, nil
}


// send a message to the dispatcher
func (proxy *DispatcherProxy) Send() (string, error) {
	url := fmt.Sprintf("%s/message", proxy.Host)

	message := "Yow!"

	resp, err := http.Post(url, "application/json", bytes.NewBufferString(message))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var obj map[string]interface{}
	err = json.Unmarshal(buf, &obj)
	if err != nil {
		return "", err
	}

	value, ok := obj["status"]
	if !ok {
		return "", errors.New("invalid response from dispatcher server")
	}
	s := value.(string)
	return s, nil
}
