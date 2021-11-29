package taskservice

import (
	Task "taskapi.com/m/v2/src/models/task"
	taskrepo "taskapi.com/m/v2/src/persistence/taskRepo"
)

type taskService struct {
	repository taskrepo.TaskRepositer
}

type TaskServicer interface {
	GetTasks() ([]Task.Tasker, error)
	GetTask(id string) (Task.Tasker, error)
	AddTask(taskDto Task.NewTaskInputDto) error
	DeleteTask(id string) error
	EditTask(id string, newTaskInfo Task.EditTaskDto) error
	CompleteTask(id string) error
}

func CreateTaskService(tr taskrepo.TaskRepositer) TaskServicer {
	ts := &taskService{}
	ts.repository = tr
	return ts
}

func (ts taskService) GetTasks() ([]Task.Tasker, error) {
	tasks, err := ts.repository.GetAllTasks()
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (ts taskService) GetTask(id string) (Task.Tasker, error) {
	task, err := ts.repository.GetTask(id)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (ts taskService) AddTask(taskDto Task.NewTaskInputDto) error {
	err := ts.repository.AddTask(taskDto)
	if err != nil {
		return err
	}
	return nil
}

func (ts taskService) DeleteTask(id string) error {
	err := ts.repository.DeleteTask(id)
	if err != nil {
		return err
	}
	return nil
}

func (ts taskService) EditTask(id string, newTaskInfo Task.EditTaskDto) error {
	err := ts.repository.EditTask(id, newTaskInfo)
	if err != nil {
		return err
	}
	return nil
}

func (ts taskService) CompleteTask(id string) error {
	err := ts.repository.CompleteTask(id)
	if err != nil {
		return err
	}
	return nil
}
