package jobs

import (
	"fmt"
	"log"
	"time"
)

type MessageId int64

var currentId MessageId = 1

type JobType int

const (
	ServiceJob JobType = 1
	StatusJob  JobType = 2
)

type Job struct {
	id        MessageId
	timestamp time.Time
	jobType   JobType
}

type Jobs struct {
	table map[MessageId]Job
}

func NewJobs() *Jobs {
	var jobs Jobs
	jobs.table = make(map[MessageId]Job)
	return &jobs
}

func (jobs *Jobs) Add(job *Job) {
	if _, ok := jobs.table[job.id]; ok {
		panic("yow")
	}
	jobs.table[job.id] = *job
}

func (jobs *Jobs) Dump() {
	for k, v := range jobs.table {
		log.Printf("key=%v  value=%v\n", k, v)
	}
}

func NewJob(jobType JobType) *Job {
	var job = Job{id: currentId}
	currentId++

	job.timestamp = time.Now()
	job.jobType = jobType

	return &job
}

func (job Job) String() string {
	return fmt.Sprintf("{id:%v jobType:%v}", job.id, job.jobType)
}
