// TODO Вы можете редактировать этот файл по вашему усмотрению

package pkg

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// TODO реализовать только нужные

type Database struct {
}

func (d *Database) Begin(ctx context.Context) (pgx.Tx, error) {
	//TODO implement me
	panic("implement me")
}

func (d *Database) Commit(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (d *Database) Rollback(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (d *Database) Exec(ctx context.Context, sql string, arguments ...any) (commandTag pgconn.CommandTag, err error) {
	//TODO implement me
	panic("implement me")
}

func (d *Database) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	//TODO implement me
	panic("implement me")
}

func (d *Database) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	//TODO implement me
	panic("implement me")
}
