package main

import (
	"workerpool/internal/app"
	"workerpool/internal/infrastructure"
	"workerpool/internal/presentation"
)

func main() {
	workerPool := infrastructure.NewWorkerPool(3, 10)
	taskService := app.NewTaskService(workerPool)
	server := presentation.NewHTTPServer(taskService)

	server.Start()
}
