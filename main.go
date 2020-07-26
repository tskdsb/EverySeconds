package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tskdsb/EverySeconds/job"
	"net/http"
)

var (
	JobList map[string]*job.Job
)

func init() {
	JobList = make(map[string]*job.Job)
}

func main() {
	router := gin.Default()

	// Create job
	router.PUT("/job/:name", func(c *gin.Context) {
		jobName := c.Param("name")
		//user := c.GetHeader("user")
		newJob := job.New(jobName)
		JobList[newJob.ID] = newJob
	})

	// List job
	router.GET("/job/*name", func(c *gin.Context) {
		//jobName := c.Param("name")
		//user := c.GetHeader("user")
		// TODO: filter by user

		c.IndentedJSON(http.StatusOK, JobList)

	})

	// Start job
	router.POST("/job/:id", func(c *gin.Context) {
		jobID := c.Param("id")
		action := c.Query("action")

		switch action {
		case job.ActionJobStart:
			JobList[jobID].Start()
		case job.ActionJobStop:
			JobList[jobID].Stop()
		case job.ActionJobFinish:
			JobList[jobID].Finish()
		}
	})

	panic(router.Run(":26666"))
}
