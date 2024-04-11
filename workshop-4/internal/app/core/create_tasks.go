package core

import "context"

type CreateTaskRequest struct {
	Name        string
	Description string
	Overlap     uint32
}

type CreateTasksRequest struct {
	Tasks       []CreateTaskRequest
	TaskGroupID uint32
	CustomerId  uint32
}

func (s *Service) CreateTasks(ctx context.Context, request CreateTasksRequest) ([]uint32, error) {
	tasks := buildTasksFromCreateRequest(request)
	return s.taskService.Create(ctx, tasks, request.TaskGroupID)
}
