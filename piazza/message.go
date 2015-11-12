package piazza

import (
	"fmt"
	//"log"
	"time"
)

//---------------------------------------------------------

// id of 0 is always invalid
type MessageId int64

var nextMessageId MessageId = 1

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

type Status int

const (
	InvalidStatus Status = iota
	SuccessStatus
	ErrorStatus
	FailureStatus
)

type Message struct {
	// onl set by Gateway (or by Dispatcher if the Message bypasses the Gateway)
	Id MessageId `json:message_id`

	Type      MessageType `json:type`
	Timestamp time.Time   `json:timestamp`
	User      User        `json:user`

	// only used for Response messages
	Status Status `json:status`

	// only used for Response messages (optional)
	Error error `json:error`

	// exactly one of the following must be present
	CreateJobPayload             `json:create_job_payload`
	CreateJobResponsePayload     `json:create_job_response_payload`
	ReadJobPayload               `json:read_job_payload`
	ReadJobResponsePayload       `json:read_job_response_payload`
	DeleteJobPayload             `json:delete_job_payload`
	DeleteJobResponsePayload     `json:delete_job_response_payload`
	CreateServicePayload         `json:create_service_payload`
	CreateServiceResponsePayload `json:create_service_response_payload`
	ReadServicePayload           `json:read_service_payload`
	ReadServiceResponsePayload   `json:read_service_response_payload`
	DeleteServicePayload         `json:delete_service_payload`
	DeleteServiceResponsePayload `json:delete_service_response_payload`
}

//---------------------------------------------------------

////////////////////////////

type CreateJobPayload struct {
	ServiceName string            `json:service_name`
	Parameters  map[string]string `json:parameters`

	// set by user for his own reasons (optional)
	Comments string `json:comments`
}

type CreateJobResponsePayload struct {
	JobId JobId `json:job_id`
}

type ReadJobPayload struct {
	JobId JobId `json:job_id`
}

type ReadJobResponsePayload struct {
	Status          StatusCode        `json:status`
	PercentComplete float             `json:percent_complete`
	TimeRemaining   time.Duration     `json:time_remaining`
	ServiceName     string            `json:service_name`
	Parameters      map[string]string `json:parameters`
	Comments        string            `json:comments`
}

type DeleteJobPayload struct {
	JobId JobId `json:job_id`
}

type DeleteJobResponsePayload struct {
}

//---------------------------------------------------------

type CreateServicePayload struct {
	Type ServiceType `json:service_type`
}

type CreateServiceResponsePayload struct {
	Service Service `json:service`
}

type ReadServicePayload struct {
	Id ServiceId `json:service_id`
}

type ReadServiceResponsePayload struct {
	running bool `json:is_running`
}

type DeleteServicePayload struct {
	Id ServiceId `json:service_id`
}

type DeleteServiceResponsePayload struct {
}

//---------------------------------------------------------
