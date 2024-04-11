package task_group

import "context"

type Repository interface {
	create(ctx context.Context, taskGroup taskGroupRow) (uint32, error)
	getByID(ctx context.Context, taskGroupId uint32) (taskGroupRow, error)
	update(ctx context.Context, taskGroup taskGroupRow) error
}

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) Create(ctx context.Context, taskGroup TaskGroup) (uint32, error) {
	return s.repository.create(ctx, taskGroup.mapToModel())
}

func (s *Service) GetTaskGroup(ctx context.Context, taskGroupId uint32) (TaskGroup, error) {
	taskGroup, err := s.repository.getByID(ctx, taskGroupId)
	resultTaskGroup := TaskGroup{}
	if err != nil {
		return resultTaskGroup, err
	}

	return *resultTaskGroup.mapFromModel(taskGroup), nil
}

func (s *Service) Update(ctx context.Context, taskGroup TaskGroup) error {
	return s.repository.update(ctx, taskGroup.mapToModel())
}
