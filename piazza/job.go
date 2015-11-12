package piazza

import (
	"fmt"
	"log"
	//"time"
)

type JobId int64

var currentJobId JobId = 1

type JobStatus int

type JobType int

const (
	StatusInvalid JobStatus = iota
	StatusSubmitted
	StatusDispatched
	StatusRunning
	StatusCompleted
	StatusFailed
)

type Job struct {
	id        JobId
	status    JobStatus
	messageId MessageId
	jtype     JobType
}

type JobTable struct {
	table map[JobId]Job
}

func NewJobTable() *JobTable {
	var t JobTable
	t.table = make(map[JobId]Job)
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
	var job = Job{id: currentJobId}
	currentJobId++

	job.jtype = jobType

	return &job
}

func (job Job) String() string {
	return fmt.Sprintf("{id:%v jobType:%v}", job.id, job.jtype)
}
