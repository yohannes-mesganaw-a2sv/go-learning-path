package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// create a Enum for the Status of the Task
var Status int

const (
	Inprogress = iota
	Completed
	NotStarted
	
)


// Task struct
type Task struct {
	ID primitive.ObjectID `bson:"_id"`
	Title string `bson:"title"`
	Status int `bson:"status"`
	Description string `bson:"description"`
}