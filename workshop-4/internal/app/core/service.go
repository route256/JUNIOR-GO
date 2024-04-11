package core

import (
	"context"
	"test_jr_6/internal/app/task"
	"test_jr_6/internal/app/task_group"
)

type Service struct {
	taskService      TaskService
	taskGroupService TaskGroupService
	ratingService    RatingService
}

type TaskService interface {
	Create(ctx context.Context, tasks []task.Task, taskGroupId uint32) ([]uint32, error)
	Update(ctx context.Context, tasks []task.Task, taskGroupId uint32) error
	CreateFromJSON(ctx context.Context, jsonString string, taskGroupId uint32) ([]uint32, error)
	GetTasksForCustomer(ctx context.Context, taskGroupId uint32) ([]task.Task, error)
	GetTaskForSolve(ctx context.Context, taskGroupId uint32) (task.Task, error)
	Solve(ctx context.Context, ID uint32, userID uint32, answer string) error
}

type TaskGroupService interface {
	Create(ctx context.Context, taskGroup task_group.TaskGroup) (uint32, error)
	GetTaskGroup(ctx context.Context, taskGroupId uint32) (task_group.TaskGroup, error)
	Update(ctx context.Context, taskGroup task_group.TaskGroup) error
}

type RatingService interface {
	Calc(ctx context.Context, userID uint32) error
}

func NewCoreService(taskService TaskService, taskGroupService TaskGroupService, ratingService RatingService) *Service {
	return &Service{
		taskService:      taskService,
		taskGroupService: taskGroupService,
		ratingService:    ratingService,
	}
}
