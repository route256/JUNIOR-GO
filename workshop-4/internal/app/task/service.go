package task

import (
	"context"
	"encoding/json"
	"fmt"
)

type Repository interface {
	create(ctx context.Context, tasks []taskRow, taskGroupId uint32) ([]uint32, error)
	update(ctx context.Context, tasks []taskRow, taskGroupId uint32) error
	getForCustomer(ctx context.Context, taskGroupId uint32) ([]taskRow, error)
	getForSolve(ctx context.Context, taskGroupId uint32) (taskRow, error)
	solve(ctx context.Context, task taskRow) error
	getByID(ctx context.Context, ID uint32) (taskRow, error)
}

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) Create(ctx context.Context, tasks []Task, taskGroupId uint32) ([]uint32, error) {
	fmt.Println(tasks)
	taskRows := mapToModels(tasks)

	return s.repository.create(ctx, taskRows, taskGroupId)
}

func (s *Service) Update(ctx context.Context, tasks []Task, taskGroupId uint32) error {
	taskRows := mapToModels(tasks)

	return s.repository.update(ctx, taskRows, taskGroupId)
}

func (s *Service) CreateFromJSON(ctx context.Context, jsonString string, taskGroupId uint32) ([]uint32, error) {
	var tasks []Task
	err := json.Unmarshal([]byte(jsonString), &tasks)
	if err != nil {
		return []uint32{}, err
	}

	return s.Create(ctx, tasks, taskGroupId)
}

func (s *Service) GetTasksForCustomer(ctx context.Context, taskGroupId uint32) ([]Task, error) {
	tasks, err := s.repository.getForCustomer(ctx, taskGroupId)
	if err != nil {
		return []Task{}, err
	}

	responseTasks := make([]Task, len(tasks))
	for _, task := range tasks {
		responseTask := Task{}
		responseTasks = append(responseTasks, *responseTask.mapFromModel(task))
	}

	return responseTasks, nil
}

func (s *Service) GetTaskForSolve(ctx context.Context, taskGroupId uint32) (Task, error) {
	task, err := s.repository.getForSolve(ctx, taskGroupId)
	if err != nil {
		return Task{}, err
	}

	responseTask := Task{}
	return *responseTask.mapFromModel(task), nil
}

func (s *Service) Solve(ctx context.Context, ID uint32, userID uint32, answer string) error {
	taskRow, err := s.repository.getByID(ctx, ID)
	if err != nil {
		return err
	}

	task := Task{}
	task.mapFromModel(taskRow)
	task.Answers = append(task.Answers, UserAnswer{
		UserID: userID,
		Answer: answer,
	})
	task.Overlap--

	return s.repository.solve(ctx, task.mapToModel())
}
