package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"to-do-list/internal/tasks"

	_ "github.com/mattn/go-sqlite3"
)

type Client struct {
	conn *sql.DB
}

func NewClient(conn *sql.DB) *Client {
	return &Client{conn: conn}
}

const sqlCreateTask = "INSERT INTO tasks(name, description, deadline, status) values (?, ?, ?, ?) RETURNING *"

func (c *Client) Create(ctx context.Context, task *tasks.Task) (
	*tasks.Task,
	error,
) {
	log.Println("database Create start")

	tx, err := c.conn.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare(sqlCreateTask)
	if err != nil {
		return nil, err
	}

	result, err := stmt.ExecContext(
		ctx,
		task.Name,
		task.Description,
		task.Deadline,
		task.Status,
	)
	if err != nil {
		return nil, fmt.Errorf("err: %v, rollback: %v", err, tx.Rollback())
	}

	if _err := stmt.Close(); _err != nil {
		return nil, _err
	}

	if _err := tx.Commit(); _err != nil {
		return nil, _err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	log.Println("database Create end")
	return c.Read(ctx, int(id))
}

const sqlReadTask = "SELECT * FROM tasks WHERE id = ? LIMIT 1"

func (c *Client) Read(ctx context.Context, id int) (*tasks.Task, error) {
	selectStmt, err := c.conn.Prepare(sqlReadTask)
	if err != nil {
		return nil, err
	}

	row := selectStmt.QueryRowContext(ctx, id)

	var newTask tasks.Task
	if _err := row.Scan(
		&newTask.ID,
		&newTask.Name,
		&newTask.Description,
		&newTask.Deadline,
		&newTask.Status,
	); _err != nil {
		return nil, _err
	}

	return &newTask, nil
}

const sqlUpdateTask = "UPDATE tasks SET name = ?, description = ?, deadline = ?, status = ? WHERE id = ?"

func (c *Client) Update(
	ctx context.Context,
	id int,
	task *tasks.Task,
) (*tasks.Task, error) {
	tx, err := c.conn.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare(sqlUpdateTask)
	if err != nil {
		return nil, err
	}

	result, err := stmt.ExecContext(
		ctx,
		&task.Name,
		&task.Description,
		&task.Deadline,
		&task.Status,
		id,
	)
	if err != nil {
		return nil, err
	}

	if _err := stmt.Close(); _err != nil {
		return nil, _err
	}

	if _err := tx.Commit(); _err != nil {
		return nil, _err
	}

	if cnt, _err := result.RowsAffected(); cnt == 0 {
		return nil, fmt.Errorf("no row affected %v", _err)
	}

	return c.Read(ctx, id)
}

const sqlDeleteTask = "DELETE FROM tasks WHERE id = ?"

func (c *Client) Delete(ctx context.Context, id int) error {
	tx, err := c.conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(sqlDeleteTask)
	if err != nil {
		return err
	}

	result, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}

	if cnt, _err := result.RowsAffected(); cnt == 0 {
		return fmt.Errorf("no row affected %v", _err)
	}

	return nil
}

const sqlListTasks = "SELECT * FROM tasks"

func (c *Client) List(ctx context.Context) (tasks.TaskList, error) {
	log.Println("database List start")
	var list tasks.TaskList

	rows, err := c.conn.QueryContext(ctx, sqlListTasks)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var value tasks.Task
		if _err := rows.Scan(
			&value.ID,
			&value.Name,
			&value.Description,
			&value.Deadline,
			&value.Status,
		); _err != nil {
			return nil, _err
		}

		list = append(list, &value)
	}

	log.Println("database List end")
	return list, nil
}
