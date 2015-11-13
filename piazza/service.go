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
	RegistryService: "RegistryService",
	GatewayService: "GatewayService",
	DispatcherService: "DispatcherService",
	EchoService: "EchoService",
}

type ServiceId int64

var nextServiceId ServiceId = 1

type Service struct {
	Id      ServiceId   `json:"id"`
	Type    ServiceType `json:type`
	Host    string      `json:"host"`
	Comment string      `json:"comment"`
}


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
func ServiceTypeToString(stype ServiceType) string {
	s, ok := serviceTypeNames[stype]
	if !ok {
		return "InvalidService"
	}
	return s
}

func (e *Service) String() string {
	return fmt.Sprintf("[id:%d, type:%v, host:%s, comment:%s]", e.Id, e.Type, e.Host, e.Comment)
}

func NewService(stype ServiceType, host string, comment string) (*Service) {
	id := nextServiceId
	nextServiceId++

	e := &Service{Id: id, Type: stype, Host: host, Comment: comment}

	return e
}

func NewServiceFromBytes(buf []byte) (*Service, error) {
	var e Service
	err := json.Unmarshal(buf, &e)
	if err != nil {
		return nil, err
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
