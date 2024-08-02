package router

import "github.com/gin-gonic/gin"



func InitRouter(router *gin.Engine) {
	
	r := router.Group("/api")
	{
		// the route for the tasks
		r.POST("/tasks", )
		
	}
	
}
