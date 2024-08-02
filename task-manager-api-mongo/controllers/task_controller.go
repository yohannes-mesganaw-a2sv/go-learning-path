package controllers

import (
	dtos "task-management-api-mongo/dtos"
	service "task-management-api-mongo/services"

	"github.com/gin-gonic/gin"
)

func createTask(c *gin.Context) {
	var task dtos.CreateTaskDto

	err := c.BindJSON(&task)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request",
		})
		return
	}

	// save the task to the database using service
	service.SaveTask(task)

	// return a success message
	c.JSON(200, gin.H{
		"message": "Task created successfully",
	})


}
