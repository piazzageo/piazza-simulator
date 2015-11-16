package piazza

import (
	//"fmt"
	//"log"
	"encoding/json"
	"time"
)

type Any interface{}

//---------------------------------------------------------

type MessageType int

const (
	// zero value
	InvalidMessage MessageType = iota

	// run a task on a service
	CreateJobMessage
	CreateJobResponseMessage

	// get status, get original params, get results
	ReadJobMessage
	ReadJobResponseMessage

	//UpdateJob

	// kill a job
	DeleteJobMessage
	DeleteJobResponseMessage

	// start up a service
	CreateServiceMessage
	CreateServiceResponseMessage

	// get info about a running service
	ReadServiceMessage
	ReadServiceResponseMessage

	//UpdateService

	// stop a service
	DeleteServiceMessage
	DeleteServiceResponseMessage
)

type StatusCode int

const (
	InvalidStatus StatusCode = iota
	SuccessStatus
	ErrorStatus
	FailureStatus
)

type Message struct {
	// onl set by Gateway (or by Dispatcher if the Message bypasses the Gateway)
	Id PiazzaId `json:"id"`

	Type      MessageType `json:"type"`
	Timestamp time.Time   `json:"timestamp"`
	User      *User       `json:"user,omitempty"`

	// only used for Response messages
	Status StatusCode `json:status`
	Error  *error     `json:error,omitEmpty`

	// exactly one of the following must be present
	CreateJobPayload             *CreateJobPayload             `json:"create_job_payload,omitempty"`
	CreateJobResponsePayload     *CreateJobResponsePayload     `json:"create_job_response_payload,omitempty"`
	ReadJobPayload               *ReadJobPayload               `json:"read_job_payload,omitempty"`
	ReadJobResponsePayload       *ReadJobResponsePayload       `json:"read_job_response_payload,omitempty"`
	DeleteJobPayload             *DeleteJobPayload             `json:"delete_job_payload,omitempty"`
	DeleteJobResponsePayload     *DeleteJobResponsePayload     `json:"delete_job_response_payload,omitempty"`
	CreateServicePayload         *CreateServicePayload         `json:"create_service_payload,omitempty"`
	CreateServiceResponsePayload *CreateServiceResponsePayload `json:"create_service_response_payload,omitempty"`
	ReadServicePayload           *ReadServicePayload           `json:"read_service_payload,omitempty"`
	ReadServiceResponsePayload   *ReadServiceResponsePayload   `json:"read_service_response_payload,omitempty"`
	DeleteServicePayload         *DeleteServicePayload         `json:"delete_service_payload,omitempty"`
	DeleteServiceResponsePayload *DeleteServiceResponsePayload `json:"delete_service_response_payload,omitempty"`
}

func NewMessage() *Message {
	m := new(Message)
	return m
}

func NewMessageFromBytes(buf []byte) (*Message, error) {
	m := new(Message)
	err := json.Unmarshal(buf, m)
	return m, err
}

func (m *Message) ToBytes() ([]byte, error) {
	buf, err := json.Marshal(m)
	return buf, err
}

func (m *Message) ToJson() (string, error) {
	buf, err := m.ToBytes()
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

//---------------------------------------------------------

////////////////////////////

type CreateJobPayload struct {
	ServiceType ServiceType    `json:"stype"`
	Parameters  map[string]Any `json:"parameters"`

	// set by user for his own reasons (optional)
	Comment string `json:"comments"`
}

type CreateJobResponsePayload struct {
	JobId     PiazzaId  `json:"id"`
	JobStatus JobStatus `json:"status"`
}

type ReadJobPayload struct {
	JobId PiazzaId `json:"job_id"`
}

type ReadJobResponsePayload struct {
	Status          StatusCode     `json:"status"`
	PercentComplete float32        `json:"percent_complete"`
	TimeRemaining   time.Duration  `json:"time_remaining"`
	ServiceName     string         `json:"service_name"`
	Parameters      map[string]Any `json:"parameters"`
	Comments        string         `json:"comments"`
}

type DeleteJobPayload struct {
	JobId PiazzaId `json:"job_id"`
}

type DeleteJobResponsePayload struct {
}

//---------------------------------------------------------

type CreateServicePayload struct {
	Type ServiceType `json:"service_type"`
}

type CreateServiceResponsePayload struct {
	Service Service `json:"service"`
}

type ReadServicePayload struct {
	Id PiazzaId `json:"service_id"`
}

type ReadServiceResponsePayload struct {
	running bool `json:"is_running"`
}

type DeleteServicePayload struct {
	Id PiazzaId `json:"service_id"`
}

type DeleteServiceResponsePayload struct {
}

//---------------------------------------------------------
