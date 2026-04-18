package manager

import (
	"cube/task"
	"fmt"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

/**
/**
- Accept requests from users to start and stop tasks
- Schedule tasks onto worker machines
- Keep track of tasks, their states, and the machine on which they run
**/

type Manager struct {
	Pending       queue.Queue
	TaskDb        map[string][]*task.Task
	EventDb       map[string][]*task.TaskEvent
	Workers       []string
	WorkerTaskMap map[string][]uuid.UUID
	TaskWorkermap map[uuid.UUID]string
}

func (*Manager) SelectWorker() {
	fmt.Println("Selecting an appropriate worker for a task")
}

func (m *Manager) UpdateTasks() {
	fmt.Println("Updating tasks")
}

func (m *Manager) SendWork() {
	fmt.Println("Sending work to workers")
}
