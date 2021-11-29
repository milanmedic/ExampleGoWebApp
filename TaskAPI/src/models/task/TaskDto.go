package models

import "github.com/google/uuid"

type NewTaskInputDto struct {
	Name        string
	Description string
}

type EditTaskDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
}

type TaskDto struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      bool      `json:"status"`
}

func (t TaskDto) GetId() string {
	return string(t.Id.String())
}

func (t TaskDto) GetName() string {
	return t.Name
}

func (t TaskDto) GetDescription() string {
	return t.Description
}

func (t TaskDto) GetStatus() bool {
	return t.Status
}
