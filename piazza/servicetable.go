package piazza

import (
	"fmt"
	"strconv"
	"encoding/json"
)


//---------------------------------------------------------------------


type ServiceTable struct {
	table     map[string]*ServiceEntry `json:table`
	currentId int	                   `json:current_id`
}

type ServiceEntry struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Id         string 	`json:"id"`
}


//---------------------------------------------------------------------


func NewServiceTable() (*ServiceTable) {
	var t ServiceTable
	t.table = make(map[string]*ServiceEntry)
	return &t
}


func (t *ServiceTable) Add(entry *ServiceEntry) error {
	entry.Id = strconv.Itoa(t.currentId)
	t.currentId++

	t.table[entry.Id] = entry
	return nil
}


func (t *ServiceTable) Get(id string) (entry *ServiceEntry, ok bool) {
	entry, ok = t.table[id]
	return entry, ok
}


func (t *ServiceTable) String() string {
	return fmt.Sprintf("[...table (len:%d)...]", len(t.table))
}


func (t *ServiceTable) ToBytes() ([]byte, error) {
	var buf []byte
	buf, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}

	return buf, nil
}


func (t *ServiceTable) ToJSON() (string, error) {
	buf, err := t.ToBytes()
	if err != nil {
		return "", err
	}
	return string(buf), nil
}


func (t *ServiceTable) Count() int {
	i := len(t.table)
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


func NewServiceTableFromJSON(s string) (*ServiceTable, error) {
	return NewServiceTableFromBytes([]byte(s))
}


func (a *ServiceTable) Compare(b *ServiceTable) bool {
	if a.Count() != b.Count() {
		return false
	}
	for k, _ := range a.table {
		if a.table[k].Compare(b.table[k]) == false {
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


func (a *ServiceEntry) Compare(b *ServiceEntry) bool {
	return a.Name == b.Name && a.Description == b.Description && a.Id == b.Id
}


func (e *ServiceEntry) ToBytes() ([]byte, error) {
	var buf []byte
	buf, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}

	return buf, nil
}
