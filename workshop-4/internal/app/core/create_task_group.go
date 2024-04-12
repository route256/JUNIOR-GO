package core

import "context"

type CreateTasksGroupRequest struct {
	Name            string
	Description     string
	Price           uint32
	SecondsToDecide uint32
}

func (s *Service) CreateTaskGroup(ctx context.Context, request CreateTasksGroupRequest) (uint32, error) {
	taskGroup := buildTaskGroupFromCreateRequest(request)
	return s.taskGroupService.Create(ctx, taskGroup)
}
