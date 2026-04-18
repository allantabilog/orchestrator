package task

import (
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/google/uuid"
)

/**
A type-safe enum of Task states
**/

type State int

const (
	Pending State = iota
	Scheduled
	Running
	Completed
	Failed
)

type Task struct {
	ID            uuid.UUID
	Name          string
	State         State
	Image         string
	CPU           float64
	Memory        int
	Disk          int
	ExposedPorts  nat.PortSet
	PortBindings  map[string]string
	RestartPolicy string
	StartTime     time.Time
	FinishTime    time.Time
}

// This is an internal object that the system
// will use to transition tasks from one state to the next
type TaskEvent struct {
	ID        uuid.UUID
	State     State // this is the target state
	Timestamp time.Time
	Task      Task
}
