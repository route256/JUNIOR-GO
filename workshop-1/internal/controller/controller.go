package controller

import (
	"fmt"
	"time"

	"to-do-list/internal/tasks"
)

type Storage interface {
	Create(task tasks.Task) error
	List() (tasks.TaskList, error)
}

type Controller struct {
	storage Storage
}

func (c *Controller) List() error {
	list, err := c.storage.List()
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

	return c.storage.Create(
		tasks.Task{
			Name:        name,
			Description: description,
			Deadline:    t,
		},
	)
}

func NewController(storage Storage) *Controller {
	return &Controller{storage: storage}
}
