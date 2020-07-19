package routers

import (
	. "ginlearn/controllers"
	_ "ginlearn/docs"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/hello", GetDataList)
	router.POST("/user/add", InsertUser)
	router.POST("/job/sub", SubmitJob)

	return router
}
