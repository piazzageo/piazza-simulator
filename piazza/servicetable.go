package piazza

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"errors"
	"log"
)

//---------------------------------------------------------------------

type ServiceTable struct {
	Table     map[string]*ServiceEntry `json:table`
	CurrentId int                      `json:current_id`
}

type ServiceEntry struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Id          string `json:"id"`
}

//---------------------------------------------------------------------

func NewServiceTable() *ServiceTable {
	var t ServiceTable
	t.Table = make(map[string]*ServiceEntry)
	return &t
}

func (t *ServiceTable) Add(entry *ServiceEntry) error {
	entry.Id = strconv.Itoa(t.CurrentId)
	t.CurrentId++

	t.Table[entry.Id] = entry
	return nil
}

func (t *ServiceTable) Get(id string) (entry *ServiceEntry, ok bool) {
	entry, ok = t.Table[id]
	return entry, ok
}

func (t *ServiceTable) String() string {
	return fmt.Sprintf("[...table (len:%d)...]", len(t.Table))
}

func (t *ServiceTable) ToBytes() ([]byte, error) {
	var buf []byte
	buf, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (t *ServiceTable) Count() int {
	i := len(t.Table)
	return i
}

func NewServiceTableFromBytes(buf []byte) (*ServiceTable, error) {
	var t = NewServiceTable()
	err := json.Unmarshal(buf, &t)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (a *ServiceTable) Compare(b *ServiceTable) bool {
	if a.Count() != b.Count() {
		return false
	}
	for k, _ := range a.Table {
		if a.Table[k].Compare(b.Table[k], true) == false {
			return false
		}
	}
	return true
}

//---------------------------------------------------------------------

func (e *ServiceEntry) String() string {
	return fmt.Sprintf("[name:%s, description:%s, id:%s]", e.Name, e.Description, e.Id)
}

func NewServiceEntryFromBytes(buf []byte) (*ServiceEntry, error) {
	var e ServiceEntry

	err := json.Unmarshal(buf, &e)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (a *ServiceEntry) Compare(b *ServiceEntry, checkIds bool) bool {
	if checkIds && a.Id != b.Id {
		return false
	}
	return a.Name == b.Name && a.Description == b.Description
}

func (e *ServiceEntry) ToBytes() ([]byte, error) {
	var buf []byte
	buf, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

//---------------------------------------------------------------------

// used by services to communicate with registry
func RegisterService(registryHost string, name string, description string) (int, error) {
	entry := ServiceEntry{Name: name, Description: description}

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
