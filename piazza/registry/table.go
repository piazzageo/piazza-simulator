package registry

import (
	"encoding/json"
	"fmt"
	"github.com/mpgerlek/piazza-simulator/piazza"
)

//---------------------------------------------------------------------

type serviceTable struct {
	table *piazza.ServiceList
}

//---------------------------------------------------------------------

func NewServiceTable() *serviceTable {
	var t serviceTable
	t.table = new(piazza.ServiceList)
	return &t
}

func (t *serviceTable) Add(entry *piazza.Service) error {
	t.table.Add(*entry)
	return nil
}

func (t *serviceTable) LookupByTypeString(key string) (e *piazza.ServiceList, ok bool) {

	stype := piazza.ServiceTypeFromString(key)
	if stype == piazza.InvalidService {
		return nil, false
	}
	return t.LookupByType(stype)
}

func (t *serviceTable) LookupByType(stype piazza.ServiceType) (e *piazza.ServiceList, ok bool) {

	list := new(piazza.ServiceList)
	for service := range t.table.Iterator() {
		if service.Type == stype {
			list.Add(service)
		}
	}
	return list, true
}

func (t *serviceTable) LookupAll() (e *piazza.ServiceList) {

	list := new(piazza.ServiceList)
	for service := range t.table.Iterator() {
		list.Add(service)
	}
	return list
}

// returns a list of size 0 or 1
func (t *serviceTable) LookupById(id piazza.PiazzaId) (e *piazza.ServiceList, ok bool) {
	service, ok := t.table.Get(id)
	if !ok {
		return nil, false
	}
	list := new(piazza.ServiceList)
	list.Add(service)
	return list, true
}

func (t *serviceTable) String() string {
	return fmt.Sprintf("[tablelen:%d]", t.table.Len())
}

func (t *serviceTable) Len() int {
	return t.table.Len()
}

func (a *serviceTable) Compare(b *serviceTable) bool {
	if a.Len() != b.Len() {
		return false
	}
	for service := range a.table.Iterator() {
		if bService, ok := b.table.Get(service.Id); !ok || service.Id != bService.Id {
			return false
		}
	}
	return true
}

func (t *serviceTable) ToBytes() ([]byte, error) {
	var buf []byte
	buf, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return buf, nil
}
