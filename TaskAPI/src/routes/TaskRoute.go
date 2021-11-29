package routes

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	models "taskapi.com/m/v2/src/models/task"
	taskservice "taskapi.com/m/v2/src/services/taskService"
)

var service taskservice.TaskServicer

func SetupTaskRouter(router *httprouter.Router, taskService taskservice.TaskServicer) {
	setupRoutes(router)
	service = taskService
}

func setupRoutes(router *httprouter.Router) {
	router.GET("/task/", getAllTasks)
	router.POST("/task/", addTask)
	router.GET("/task/:id", getTask)
	router.DELETE("/task/:id", deleteTask)
	router.PATCH("/task/:id", updateTask)
}

func getAllTasks(wr http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tasks, err := service.GetTasks()
	if err != nil {
		wr.WriteHeader(http.StatusInternalServerError)
	}
	if len(tasks) <= 0 {
		wr.WriteHeader(http.StatusNotFound)
	}
	output, err := json.MarshalIndent(tasks, "", "\t")
	if err != nil {
		wr.WriteHeader(http.StatusInternalServerError)
	}
	wr.Header().Set("Content-Type", "application/json")
	wr.WriteHeader(http.StatusOK)
	wr.Write(output)
}

func addTask(wr http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var task models.NewTaskInputDto
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		wr.WriteHeader(http.StatusInternalServerError)
	}
	service.AddTask(task)
	wr.WriteHeader(http.StatusCreated)
	wr.Write([]byte("Add Task Route"))
}

func getTask(wr http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	task, err := service.GetTask(id)
	if err != nil {
		wr.WriteHeader(http.StatusInternalServerError)
	}
	if task == nil {
		wr.WriteHeader(http.StatusNotFound)
	}
	output, err := json.Marshal(task)
	if err != nil {
		wr.WriteHeader(http.StatusInternalServerError)
	}
	// Only valid order for customizing headers
	// and status responses
	wr.Header().Set("Content-Type", "application/json")
	wr.WriteHeader(http.StatusOK)
	wr.Write(output)
}

func deleteTask(wr http.ResponseWriter, r *http.Request, params httprouter.Params) {
	wr.WriteHeader(http.StatusOK)
	wr.Write([]byte(params.ByName("id")))
}

func updateTask(wr http.ResponseWriter, r *http.Request, params httprouter.Params) {
	wr.WriteHeader(http.StatusOK)
	wr.Write([]byte(params.ByName("id")))
}
