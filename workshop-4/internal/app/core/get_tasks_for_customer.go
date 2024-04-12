package core

import (
	"context"
	"test_jr_6/internal/app/task"
)

func (s *Service) GetTasksForCustomer(ctx context.Context, taskGroupId uint32) ([]task.Task, error) {
	return s.taskService.GetTasksForCustomer(ctx, taskGroupId)
}
