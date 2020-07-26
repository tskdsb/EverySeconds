package job

import (
	"github.com/google/uuid"
	"time"
)

type Status string

const (
	ToDo  Status = "ToDo"
	Doing Status = "Doing"
	Did   Status = "Did"
	Done  Status = "Done"

	ActionJobCreate = "Create"
	ActionJobStart  = "Start"
	ActionJobStop   = "Stop"
	ActionJobFinish = "Finish"
)

type (
	Event struct {
		Time time.Time
		//Who    string
		Action string
		//JobRef string
	}

	Job struct {
		Name   string
		ID     string
		Status Status
		Cost   time.Duration
		Events []*Event
	}
)

func New(jobName string) *Job {
	job := &Job{
		Name:   jobName,
		ID:     uuid.New().String(),
		Status: ToDo,
	}
	job.Events = []*Event{
		{
			Time: time.Now(),
			//Who:    user,
			Action: ActionJobCreate,
			//JobRef: job.ID,
		},
	}

	return job
}

func (j *Job) Start() {
	if j.Status == ToDo || j.Status == Did {
		j.Status = Doing
		j.Events = append(j.Events, &Event{
			Time:   time.Now(),
			Action: ActionJobStart,
		})
	}
}

func (j *Job) Stop() {
	if j.Status == Doing {
		j.Status = Did
		j.Events = append(j.Events, &Event{
			Time:   time.Now(),
			Action: ActionJobStop,
		})
		j.Cost += j.Events[len(j.Events)-1].Time.Sub(j.Events[len(j.Events)-2].Time)
	}
}

func (j *Job) Finish() {
	if j.Status != Done {
		j.Status = Done
		j.Events = append(j.Events, &Event{
			Time:   time.Now(),
			Action: ActionJobFinish,
		})
		last2Event := j.Events[len(j.Events)-2]
		if last2Event.Action == ActionJobStart {
			j.Cost += j.Events[len(j.Events)-1].Time.Sub(last2Event.Time)
		}
	}
}
