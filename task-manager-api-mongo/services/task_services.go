package services

import (
	"context"
	dtos "task-management-api-mongo/Dtos"
	db "task-management-api-mongo/db"
	models "task-management-api-mongo/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// get access to the collection
var taskCollection = db.DB.Collection("tasks")

func CreateTask(task *dtos.CreateTaskDto) error {
	// create a Task
	taskModel := models.Task{
		Title: task.Title,
		Description: task.Description,
		Status: task.Status,
		ID: primitive.NewObjectID(),
	}

	// insert the Task
	taskCollection.InsertOne(context.TODO(), taskModel)
	
	return nil
}