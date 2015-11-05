package main


import (
    //"fmt"
    //"log"

    "github.com/mpgerlek/piazza-simulator/jobs"
    //"github.com/mpgerlek/piazza-simulator/kafka"    
)

func main() {
    // set up the database
    jobTable := jobs.NewJobs()
    
    job := jobs.NewJob(jobs.ServiceJob)
    jobTable.Add(job)
    
    jobTable.Dump()
}
