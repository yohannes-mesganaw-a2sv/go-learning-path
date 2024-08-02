package main

import (
	db "task-management-api-mongo/db"
	router "task-management-api-mongo/routers"

	"github.com/gin-gonic/gin"
)

func main(){
	db.Connect()
	
	// start the gin server
	r := gin.Default()
	router.InitRouter(r)
	
	r.Run("localhost:8080")
}