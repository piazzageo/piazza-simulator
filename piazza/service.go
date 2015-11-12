package piazza

import ()

type ServiceType int

const (
	RegistryService = iota
	GatewayService
	DispatcherService
	EchoService
)

type ServiceId int64

var nextServiceId ServiceId = 0

type Service struct {
	Id          string      `json:"id"`
	Name        string      `json:"name"`
	Type        ServiceType `json:type`
	Description string      `json:"description"`
}
