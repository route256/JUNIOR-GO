package core

import (
	"test_jr_6/internal/app/task_group"
)

func buildTaskGroupFromCreateRequest(request CreateTasksGroupRequest) task_group.TaskGroup {
	return task_group.TaskGroup{
		Name:            request.Name,
		Description:     request.Description,
		Price:           request.Price,
		SecondsToDecide: request.SecondsToDecide,
	}
}

func buildTaskGroupFromUpdateRequest(request UpdateTaskGroupRequest) task_group.TaskGroup {
	return task_group.TaskGroup{
		ID:              request.ID,
		Name:            request.Name,
		Description:     request.Description,
		Price:           request.Price,
		SecondsToDecide: request.SecondsToDecide,
	}
}
