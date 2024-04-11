package repository

import "time"

type Article struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	Rating    int64     `db:"rating"`
	CreatedAt time.Time `db:"created_at"`
}
