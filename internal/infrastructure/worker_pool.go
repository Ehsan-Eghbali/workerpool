package infrastructure

import (
	"log"
	"sync"
	"time"
	"workerpool/internal/domain"
)

type WorkerPool struct {
	taskQueue chan domain.Task
	wg        sync.WaitGroup
}

func NewWorkerPool(workerCount, queueSize int) *WorkerPool {
	pool := &WorkerPool{
		taskQueue: make(chan domain.Task, queueSize),
	}

	for i := 0; i < workerCount; i++ {
		go pool.worker(i)
	}

	return pool
}

func (wp *WorkerPool) worker(workerID int) {
	for task := range wp.taskQueue {
		log.Printf("Worker %d processing task %d\n", workerID, task.ID)
		time.Sleep(2 * time.Second)
	}
}

func (wp *WorkerPool) AddTask(task domain.Task) {
	wp.wg.Add(1)
	go func() {
		defer wp.wg.Done()
		wp.taskQueue <- task
	}()
}

func (wp *WorkerPool) Wait() {
	wp.wg.Wait()
	close(wp.taskQueue)
}
