package core

import "test_jr_6/internal/app/task"

func buildTasksFromCreateRequest(request CreateTasksRequest) []task.Task {
	tasks := make([]task.Task, len(request.Tasks))

	for _, requestTask := range request.Tasks {
		tasks = append(tasks, task.Task{
			Name:        requestTask.Name,
			Description: requestTask.Description,
			TaskGroupID: request.TaskGroupID,
			CustomerID:  request.CustomerId,
			Overlap:     requestTask.Overlap,
		})
	}

	return tasks
}

func buildTasksFromUpdateRequest(request UpdateTasksRequest) []task.Task {
	tasks := make([]task.Task, len(request.Tasks))

	for _, requestTask := range request.Tasks {
		tasks = append(tasks, task.Task{
			ID:          requestTask.ID,
			Name:        requestTask.Name,
			Description: requestTask.Description,
			TaskGroupID: request.TaskGroupID,
			CustomerID:  request.CustomerID,
			Overlap:     requestTask.Overlap,
		})
	}

	return tasks
}
