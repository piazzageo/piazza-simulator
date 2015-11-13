package piazza

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

//---------------------------------------------------------------------

type ServiceTable struct {
	table map[ServiceId]*Service
}

//---------------------------------------------------------------------

func NewServiceTable() *ServiceTable {
	var t ServiceTable
	t.table = make(map[ServiceId]*Service)
	return &t
}

func (t *ServiceTable) Add(entry *Service) error {
	t.table[entry.Id] = entry
	return nil
}

func (t *ServiceTable) LookupByTypeString(key string) (e *Service, ok bool) {

	stype := ServiceTypeFromString(key)
	if stype == InvalidService {
		return nil, false
	}
	return t.LookupByType(stype)
}

func (t *ServiceTable) LookupByType(stype ServiceType) (e *Service, ok bool) {

	for _, e = range t.table {
		if e.Type == stype {
			return e, true
		}
	}
	return e, false
}

func (t *ServiceTable) LookupById(id ServiceId) (e *Service, ok bool) {
	e, ok = t.table[id]
	return e, ok
}

func (t *ServiceTable) String() string {
	return fmt.Sprintf("[tablelen:%d]", len(t.table))
}

func (t *ServiceTable) Count() int {
	i := len(t.table)
	return i
}

func (a *ServiceTable) Compare(b *ServiceTable) bool {
	if a.Count() != b.Count() {
		return false
	}
	for k, _ := range a.table {
		if a.table[k].Type != b.table[k].Type {
			return false
		}
	}
	return true
}

func (t *ServiceTable) ToBytes() ([]byte, error) {
	var buf []byte
	buf, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

//---------------------------------------------------------------------

// used by services to communicate with registry
func RegisterService(registryHost string, serviceType ServiceType, comment string) (int, error) {
	entry := Service{Type: serviceType, Comment: comment}

	buf, err := entry.ToBytes()
	if err != nil {
		return 0, err
	}
	var s = string(buf)
	var myhost = fmt.Sprintf("http://%s/service", registryHost)
	log.Print(myhost)
	resp, err := http.Post(myhost, "application/json", bytes.NewBufferString(s))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	buf, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var x map[string]interface{}
	err = json.Unmarshal(buf, &x)
	if err != nil {
		return 0, err
	}

	v, ok := x["id"]
	if !ok {
		return 0, errors.New("invalid response from registry server")
	}

	vs := v.(string)
	id, err := strconv.Atoi(vs)
	if err != nil {
		return 0, err
	}
	return id, nil
}
