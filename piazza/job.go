package piazza

import (
	"fmt"
	"log"
	//"time"
)

type JobStatus int

type JobType int

const (
	JobStatusInvalid JobStatus = iota
	JobStatusSubmitted
	JobStatusDispatched
	JobStatusRunning
	JobStatusCompleted
	JObStatusFailed
)

type Job struct {
	id     PiazzaId
	status    JobStatus
	messageId PiazzaId
	jtype     JobType
}

type JobTable struct {
	table map[PiazzaId]Job
}

func NewJobTable() *JobTable {
	var t JobTable
	t.table = make(map[PiazzaId]Job)
	return &t
}

func (t *JobTable) Add(job *Job) {
	if _, ok := t.table[job.id]; ok {
		panic("yow")
	}
	t.table[job.id] = *job
}

func (t *JobTable) Dump() {
	for k, v := range t.table {
		log.Printf("key=%v  value=%v\n", k, v)
	}
}

func NewJob(jobType JobType) *Job {
	var job = Job{id: NewId()}

	job.jtype = jobType

	return &job
}

func (job Job) String() string {
	return fmt.Sprintf("{id:%v jobType:%v}", job.id, job.jtype)
}
