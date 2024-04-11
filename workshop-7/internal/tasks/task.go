package tasks

import "time"

type TaskList []*Task

type Task struct {
	ID          int
	Name        string
	Description string
	Deadline    time.Time
	Status      Status
}
