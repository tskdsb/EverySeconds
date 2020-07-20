package main

import "time"

type JobStatus string

const (
	ToDo     JobStatus = "ToDo"
	Doing    JobStatus = "Doing"
	Pending  JobStatus = "Pending"
	Finished JobStatus = "Finished"

	ActionCreateJob = "CreateJob"
)

type (
	Event struct {
		Time   time.Time
		Who    string
		Action string
		What   interface{}
	}

	Job struct {
		Name   string
		Status JobStatus
		Events []*Event
	}
)
