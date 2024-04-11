package datasource

import (
	"context"

	"to-do-list/internal/tasks"
)

type Datasource interface {
	Create(ctx context.Context, task *tasks.Task) (*tasks.Task, error)
	Read(ctx context.Context, id int) (*tasks.Task, error)
	Update(ctx context.Context, id int, task *tasks.Task) (*tasks.Task, error)
	Delete(ctx context.Context, id int) error

	List(ctx context.Context) (tasks.TaskList, error)
}
