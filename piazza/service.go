package piazza

import (
	"encoding/json"
	"fmt"
)

type ServiceType int64

const (
	InvalidService = iota
	RegistryService
	GatewayService
	DispatcherService
	EchoService
)

var serviceTypeNames = map[ServiceType]string{
	RegistryService:   "RegistryService",
	GatewayService:    "GatewayService",
	DispatcherService: "DispatcherService",
	EchoService:       "EchoService",
}

type Service struct {
	Id      PiazzaId    `json:"id"`
	Type    ServiceType `json:type`
	Host    string      `json:"host"`
	Comment string      `json:"comment,omitempty"`
}

//---------------------------------------------------------

type ServiceList struct {
	Data map[PiazzaId]Service `json:"data"`
}

func (list *ServiceList) init() {
	if list.Data == nil {
		list.Data = make(map[PiazzaId]Service)
	}
}

func (list *ServiceList) Add(service Service) {
	list.init()
	list.Data[service.Id] = service
}

func (list *ServiceList) Get(id PiazzaId) (Service, bool) {
	list.init()
	service, ok := list.Data[id]
	return service, ok
}

func (list *ServiceList) GetOne() (Service, bool) {
	list.init()
	for _, service := range list.Data {
		return service, true
	}
	return Service{}, false
}

func (list *ServiceList) Iterator() <-chan Service {
	list.init()
	ch := make(chan Service)
	go func() {
		for _, val := range list.Data {
			ch <- val
		}
		close(ch) // Remember to close or the loop never ends!
	}()
	return ch
}

func (list *ServiceList) Len() int {
	list.init()
	return len(list.Data)
}

//---------------------------------------------------------

// returns InvalidService if failure
func ServiceTypeFromString(s string) ServiceType {
	for k, v := range serviceTypeNames {
		if v == s {
			return k
		}
	}
	return InvalidService
}

// returns InvalidService if failure
func (stype ServiceType) String() string {
	s, ok := serviceTypeNames[stype]
	if !ok {
		return "InvalidService"
	}
	return s
}

//---------------------------------------------------------

func (e Service) String() string {
	s := fmt.Sprintf("[id:%v, type:%v, host:%s", e.Id, e.Type, e.Host)

	if e.Comment != "" {
		s += fmt.Sprintf(", comment:\"%s\"", e.Comment)
	}

	s += "]"

	return s
}

func NewService(stype ServiceType, host string, comment string) *Service {

	e := &Service{Id: NewId(), Type: stype, Host: host, Comment: comment}

	return e
}

func NewServiceFromBytes(buf []byte) (*Service, error) {
	var e Service
	err := json.Unmarshal(buf, &e)
	if err != nil {
		return nil, err
	}

	if e.Id == "" {
		e.Id = NewId()
	}

	return &e, nil
}

func (e1 *Service) Compare(e2 *Service) bool {
	return e1.Id == e2.Id
}

func (e *Service) ToBytes() ([]byte, error) {
	var buf []byte
	buf, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}

	return buf, nil
}
