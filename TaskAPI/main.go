package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/julienschmidt/httprouter"
	"taskapi.com/m/v2/src/persistence/db"
	taskrepo "taskapi.com/m/v2/src/persistence/taskRepo"
	Routes "taskapi.com/m/v2/src/routes"
	taskservice "taskapi.com/m/v2/src/services/taskService"
)

const PORT = 3000

func main() {

	router := httprouter.New()
	db := db.CreateDb()
	db.Connect()
	db.GetDbConnection().Ping()
	defer db.Disconnect()
	HandleRoutes(router, db)
	server := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", PORT),
		Handler: router,
	}
	fmt.Fprintf(os.Stdout, "Server started listening at: %d\n", PORT)

	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		// We received an interrupt signal, shut down.
		if err := server.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("HTTP server Shutdown: %v\n", err)
		}
		close(idleConnsClosed)
	}()

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Printf("HTTP server ListenAndServe: %v\n", err)
	}

	<-idleConnsClosed
}

func HandleRoutes(router *httprouter.Router, db *db.Database) {
	SetupTaskRoutes(router, db)
}

func SetupTaskRoutes(router *httprouter.Router, db *db.Database) {
	taskRepository := taskrepo.CreateTaskRepository(db)
	taskService := taskservice.CreateTaskService(taskRepository)
	Routes.SetupTaskRouter(router, taskService)
}
