package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"to-do-list/internal/datasource"
	"to-do-list/internal/tasks"

	"github.com/allegro/bigcache/v3"
)

const keyTaskAll = "task:all"

type Client struct {
	conn   *bigcache.BigCache
	source datasource.Datasource
}

func NewClient(conn *bigcache.BigCache, source datasource.Datasource) *Client {
	return &Client{conn: conn, source: source}
}

func (c *Client) Create(ctx context.Context, task *tasks.Task) (
	*tasks.Task,
	error,
) {
	log.Println("cache Create start")

	result, err := c.source.Create(ctx, task)
	if err != nil {
		return nil, err
	}

	taskData, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}

	if _err := c.conn.Set(
		fmt.Sprintf("task:%d", result.ID),
		taskData,
	); _err != nil {
		return nil, _err
	}

	if _err := c.conn.Delete(keyTaskAll); _err != nil {
		log.Printf("error delete all tasks: %v", _err)
	}

	log.Println("cache Create end")
	return result, nil
}

func (c *Client) Read(ctx context.Context, id int) (*tasks.Task, error) {
	key := fmt.Sprintf("task:%d", id)

	taskData, err := c.conn.Get(key)
	if errors.Is(err, bigcache.ErrEntryNotFound) {
		task, _err := c.source.Read(ctx, id)
		if _err != nil {
			return nil, _err
		}

		marshalData, marshalErr := json.Marshal(task)
		if marshalErr != nil {
			return nil, marshalErr
		}

		if cErr := c.conn.Set(key, marshalData); cErr != nil {
			return nil, cErr
		}

		return task, nil
	}
	if err != nil {
		return nil, err
	}

	var result tasks.Task
	if _err := json.Unmarshal(taskData, &result); _err != nil {
		return nil, _err
	}

	return &result, nil
}

func (c *Client) Update(
	ctx context.Context,
	id int,
	task *tasks.Task,
) (*tasks.Task, error) {
	result, err := c.source.Update(ctx, id, task)
	if err != nil {
		return nil, err
	}

	if _err := c.conn.Delete(fmt.Sprintf("task:%d", id)); _err != nil {
		log.Printf("error delete task:%d: %v", id, _err)
	}
	if _err := c.conn.Delete(keyTaskAll); _err != nil {
		log.Printf("error delete all tasks: %v", _err)
	}

	return result, nil
}

func (c *Client) Delete(ctx context.Context, id int) error {
	err := c.source.Delete(ctx, id)
	if err != nil {
		return err
	}

	return c.conn.Delete(fmt.Sprintf("task:%d", id))
}

func (c *Client) List(ctx context.Context) (tasks.TaskList, error) {
	log.Println("cache List start")

	allData, err := c.conn.Get(keyTaskAll)
	if errors.Is(err, bigcache.ErrEntryNotFound) {
		list, _err := c.source.List(ctx)
		if _err != nil {
			return nil, _err
		}

		marshalData, marshalErr := json.Marshal(list)
		if marshalErr != nil {
			return nil, marshalErr
		}

		if cErr := c.conn.Set(keyTaskAll, marshalData); cErr != nil {
			return nil, cErr
		}

		log.Println("cache List end")
		return list, nil
	}
	if err != nil {
		return nil, err
	}

	var result tasks.TaskList
	if _err := json.Unmarshal(allData, &result); _err != nil {
		return nil, _err
	}

	log.Println("cache List end")

	return result, err
}
