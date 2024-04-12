package core

import "context"

func (s *Service) CreateTasksFromJSON(ctx context.Context, jsonString string, taskGroupId uint32) ([]uint32, error) {
	return s.taskService.CreateFromJSON(ctx, jsonString, taskGroupId)
}
