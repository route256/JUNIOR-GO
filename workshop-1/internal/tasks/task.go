package tasks

import "time"

type TaskList []*Task

type Task struct {
	Name        string
	Description string
	Deadline    time.Time
	Status      Status
}
