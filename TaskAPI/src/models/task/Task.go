package models

import "github.com/google/uuid"

// TODO: Work with TASKDTO instead of task model itself
// have to be uppercase for marshaller to work
type task struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      bool      `json:"status"`
}

type Tasker interface {
	GetId() string
	GetName() string
	GetDescription() string
	GetStatus() bool
}

func CreateTask(name, description string) *task {
	return &task{Id: uuid.New(), Name: name, Description: description, Status: false}
}

func (t task) GetId() string {
	return string(t.Id.String())
}

func (t task) GetName() string {
	return t.Name
}

func (t task) GetDescription() string {
	return t.Description
}

func (t task) GetStatus() bool {
	return t.Status
}
