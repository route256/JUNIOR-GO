package core

import "context"

type UpdateTaskRequest struct {
	ID          uint32
	Name        string
	Description string
	Overlap     uint32
}

type UpdateTasksRequest struct {
	Tasks       []UpdateTaskRequest
	TaskGroupID uint32
	CustomerID  uint32
}

func (s *Service) UpdateTasks(ctx context.Context, request UpdateTasksRequest) error {
	tasks := buildTasksFromUpdateRequest(request)
	return s.taskService.Update(ctx, tasks, request.TaskGroupID)
}
