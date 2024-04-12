package task

import (
	"time"
)

type userAnswerRow struct {
	UserId uint32 `db:"user_id"`
	Answer string `db:"answer"`
}

type taskRow struct {
	ID           uint32          `db:"id"`
	Name         string          `db:"name"`
	Description  string          `db:"description"`
	TaskGroupID  uint32          `db:"task_group_id"`
	CustomerId   uint32          `db:"user_id"`
	Answers      []userAnswerRow `db:"answers"`
	Answer       string          `db:"answer"`
	Overlap      uint32          `db:"overlap"`
	FirstOverlap uint32          `db:"first_overlap"`
	StartedAt    time.Time       `db:"started_at"`
	FinishedAt   time.Time       `db:"finished_at"`
}
