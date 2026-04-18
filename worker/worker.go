package worker

import (
	"fmt"

	"github.com/allantabilog/orchestrator/task"
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

/**
- Run tasks as Docker containers
- Accept tasks to run from a manager
- Provide relevant statistics to the manager for the purpose of scheduling tasks
- Keep track of its tasks and their state
**/

type Worker struct {
	Name      string
	Queue     queue.Queue
	Db        map[uuid.UUID]*task.Task
	TaskCount int
}

func (w *Worker) CollectStats() {
	fmt.Println("Collecting stats")
}

func (w *Worker) RunTask() {
	fmt.Println("Starting or stopping a task")
}

func (w *Worker) StartTask() {
	fmt.Println("Starting a task")
}
func (w *Worker) StopTask() {
	fmt.Println("Stopping a task")
}
