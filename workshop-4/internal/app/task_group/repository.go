package task_group

import (
	"context"
	"fmt"
)

type Connector interface {
}

type postgresRepository struct {
	connector Connector
}

func NewRepository(connector Connector) *postgresRepository {
	return &postgresRepository{
		connector:      connector,
	}
}

func (r postgresRepository) create(ctx context.Context, taskGroup taskGroupRow) (uint32, error) {
	return 1, nil
}

func (r postgresRepository) getByID(ctx context.Context, taskGroupId uint32) (taskGroupRow, error) {
	return taskGroupRow{
		ID:              taskGroupId,
		Name:            "test name",
		Description:     "test description",
		Price:           50,
		SecondsToDecide: 15000,
	}, nil
}

func (r postgresRepository) update(ctx context.Context, taskGroup taskGroupRow) error {
	fmt.Println(taskGroup)
	return nil
}

