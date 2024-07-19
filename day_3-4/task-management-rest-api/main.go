package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID          string `json: "id"`
	Title       string `json: "title" `
	Description string `json: "description"`
	Status      string `json: "status"`
	DueDate     string `json: "dueDate"`
}

// Mock tasks
var Tasks = []Task{
	{ID: "1", Title: "Task 1", Description: "This is a task", Status: "In Progress", DueDate: "2021-10-01"},
	{ID: "2", Title: "Task 2", Description: "This is a task", Status: "In Progress", DueDate: "2021-10-01"},
	{ID: "3", Title: "Task 3", Description: "This is a task", Status: "In Progress", DueDate: "2021-10-01"},
	{ID: "4", Title: "Task 4", Description: "This is a task", Status: "In Progress", DueDate: "2021-10-01"},
	{ID: "5", Title: "Task 5", Description: "This is a task", Status: "In Progress", DueDate: "2021-10-01"},
}

func getAllTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Tasks)
}

func getSingleTask(c *gin.Context) {
	id := c.Param("id")

	for _, task := range Tasks {
		if task.ID == id {
			c.IndentedJSON(http.StatusOK, task)
			return
		}

	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task Not Found"})

}

func updateTask(c *gin.Context) {
	id := c.Param("id")
	// var task Task
	var task Task

	// update task from the given input
	for i := 0; i < len(Tasks); i++ {
		if Tasks[i].ID == id {
			err := c.BindJSON(&task)

			if err != nil {
				c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
				return
			}
			Tasks[i] = task
			c.IndentedJSON(http.StatusOK, Tasks[i])
			return
		}

	}
	c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Task Not Found"})
}

func postTask(c *gin.Context) {
	var task Task

	if err := c.BindJSON(&task); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Request"})
		return

	}

	Tasks = append(Tasks, task)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "success"})
}

func main() {
	router := gin.Default()

	router.GET("/tasks", getAllTasks)
	router.GET("/tasks/:id", getSingleTask)
	router.POST("/tasks", postTask)
	router.PUT("/tasks/:id", updateTask)

	router.Run("localhost:9000")

}
