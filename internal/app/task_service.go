package app

import (
	"workerpool/internal/domain"
	"workerpool/internal/infrastructure"
)

type TaskService struct {
	workerPool *infrastructure.WorkerPool
}

func NewTaskService(workerPool *infrastructure.WorkerPool) *TaskService {
	return &TaskService{workerPool: workerPool}
}

func (s *TaskService) ProcessTask(id int, data string) {
	task := domain.Task{ID: id, Data: data}
	s.workerPool.AddTask(task)
}
