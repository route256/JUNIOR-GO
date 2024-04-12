package core

import "context"

type UpdateTaskGroupRequest struct {
	ID              uint32
	Name            string
	Description     string
	Price           uint32
	SecondsToDecide uint32
}

func (s *Service) UpdateTaskGroup(ctx context.Context, request UpdateTaskGroupRequest) error {
	taskGroup := buildTaskGroupFromUpdateRequest(request)
	return s.taskGroupService.Update(ctx, taskGroup)
}
