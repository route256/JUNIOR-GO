package task

import (
	"context"
	"fmt"
	"time"
)

type Connector interface {
}

type postgresRepository struct {
	connector Connector
}

func NewRepository(connector Connector) *postgresRepository {
	return &postgresRepository{
		connector: connector,
	}
}

func (r postgresRepository) create(ctx context.Context, tasks []taskRow, taskGroupId uint32) ([]uint32, error) {
	var ids []uint32
	tasksLen := uint32(len(tasks))
	for i := uint32(1); i <= tasksLen; i++ {
		ids = append(ids, i)
	}

	return ids, nil
}

func (r postgresRepository) update(ctx context.Context, tasks []taskRow, taskGroupId uint32) error {
	fmt.Println(tasks)
	return nil
}

func (r postgresRepository) getForCustomer(ctx context.Context, taskGroupId uint32) ([]taskRow, error) {
	var taskRows []taskRow
	tasksLen := uint32(5)
	for i := uint32(1); i <= tasksLen; i++ {
		taskRows = append(taskRows, taskRow{
			ID:           i,
			Name:         "test name",
			Description:  "test description",
			TaskGroupID:  taskGroupId,
			CustomerId:   i + 15,
			Overlap:      5,
			FirstOverlap: 5,
			StartedAt:    time.Now(),
			FinishedAt:   time.Now(),
		})
	}

	return taskRows, nil
}

func (r postgresRepository) getByID(ctx context.Context, ID uint32) (taskRow, error) {
	return taskRow{
		ID:           ID,
		Name:         "test name",
		Description:  "test description",
		TaskGroupID:  56,
		CustomerId:   15,
		Overlap:      5,
		FirstOverlap: 5,
		StartedAt:    time.Now(),
		FinishedAt:   time.Now(),
	}, nil
}

func (r postgresRepository) getForSolve(ctx context.Context, taskGroupId uint32) (taskRow, error) {
	return taskRow{
		ID:           123,
		Name:         "test name",
		Description:  "test description",
		TaskGroupID:  taskGroupId,
		CustomerId:   15,
		Overlap:      5,
		FirstOverlap: 5,
		StartedAt:    time.Now(),
		FinishedAt:   time.Now(),
	}, nil
}

func (r postgresRepository) solve(ctx context.Context, task taskRow) error {
	fmt.Println(task)
	return nil
}
