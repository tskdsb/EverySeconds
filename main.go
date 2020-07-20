package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var (
	JobList []*Job
)

func init() {
	JobList = make([]*Job, 0, 4)
}

func main() {
	router := gin.Default()

	// Create job
	router.POST("/job/:name", func(c *gin.Context) {
		jobName := c.Param("name")
		user := c.GetHeader("user")
		JobList = append(JobList, &Job{
			Name:   jobName,
			Status: ToDo,
			Events: []*Event{
				{
					Time:   time.Now(),
					Who:    user,
					Action: ActionCreateJob,
					What:   nil,
				},
			},
		})
	})

	// List job
	router.GET("/job/*name", func(c *gin.Context) {
		jobName := c.Param("name")
		//user := c.GetHeader("user")

		// TODO: filter by user
		if jobName == "/" {
			c.IndentedJSON(http.StatusOK, JobList)
		} else {
			for _, job := range JobList {
				if job.Name == jobName {
					c.IndentedJSON(http.StatusOK, job)
				}
			}
		}
	})

	panic(router.Run(":26666"))
}
