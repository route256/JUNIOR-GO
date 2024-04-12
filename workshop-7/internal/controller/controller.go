package controller

import (
	"context"
	"fmt"
	"log"
	"time"

	"to-do-list/internal/datasource"
	"to-do-list/internal/tasks"
)

type Storage interface {
	Create(task tasks.Task) error
	List() (tasks.TaskList, error)
}

type Controller struct {
	storage datasource.Datasource
}

func NewController(storage datasource.Datasource) *Controller {
	return &Controller{storage: storage}
}

func (c *Controller) List() error {
	list, err := c.storage.List(context.Background())
	if err != nil {
		return err
	}

	list, err = c.storage.List(context.Background())
	if err != nil {
		return err
	}

	for i := range list {
		fmt.Printf(
			"status: %s, name: %s, description: %s, deadline: %s\n",
			list[i].Status,
			list[i].Name,
			list[i].Description,
			list[i].Deadline,
		)
	}

	return nil
}

func (c *Controller) Create(name, description, deadline string) error {
	t, err := time.Parse("2006-01-02", deadline)
	if err != nil {
		return err
	}

	task, err := c.storage.Create(
		context.Background(), // TODO
		&tasks.Task{
			Name:        name,
			Description: description,
			Deadline:    t,
		},
	)

	fmt.Println(task)

	return err
}

func (c *Controller) Update(
	id int,
	name, description, deadline string,
	status int,
) error {
	t, err := time.Parse("2006-01-02", deadline)
	if err != nil {
		return err
	}

	result, err := c.storage.Update(
		context.Background(),
		id,
		&tasks.Task{
			ID:          id,
			Name:        name,
			Description: description,
			Deadline:    t,
			Status:      tasks.Status(status),
		},
	)
	if err != nil {
		return err
	}

	log.Println(result)

	return nil
}
