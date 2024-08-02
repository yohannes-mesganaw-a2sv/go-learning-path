package dtos


// TaskDto struct
// create a task Dto

type TaskDto struct {
	ID string
	Title string
	Status int
	Description string
}


type CreateTaskDto struct {
	Title string
	Description string
	Status int
}

type UpdateTaskDto struct {
	ID string 
	Title string
	Status int
	Description string
}