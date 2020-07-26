package api

import "github.com/gin-gonic/gin"

var (
	Router        *gin.Engine
	JobRouterList []gin.IRouter
)

func init() {
	Router = gin.Default()

}
