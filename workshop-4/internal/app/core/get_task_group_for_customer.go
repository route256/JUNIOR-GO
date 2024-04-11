package core

import (
	"context"
	"test_jr_6/internal/app/task_group"
)

func (s *Service) GetTaskGroupForCustomer(ctx context.Context, taskGroupId uint32) (task_group.TaskGroup, error) {
	return s.taskGroupService.GetTaskGroup(ctx, taskGroupId)
}
