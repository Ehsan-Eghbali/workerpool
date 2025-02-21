package main

import (
	"workerpool/internal/app"
	"workerpool/internal/infrastructure"
)

func main() {
	workerPool := infrastructure.NewWorkerPool(3, 10)
	taskService := app.NewTaskService(workerPool)
	server := infrastructure.NewHTTPServer(taskService)

	server.Start()
}
