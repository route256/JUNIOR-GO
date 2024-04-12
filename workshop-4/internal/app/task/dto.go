package task

import "time"

type UserAnswer struct {
	UserID uint32 `json:"userId"`
	Answer string `json:"answer"`
}

type Task struct {
	ID           uint32       `json:"id"`
	Name         string       `json:"name"`
	Description  string       `json:"description"`
	TaskGroupID  uint32       `json:"taskGroupId"`
	CustomerID   uint32       `json:"userId"`
	Answers      []UserAnswer `json:"answers"`
	Answer       string       `json:"answer"`
	Overlap      uint32       `json:"overlap"`
	FirstOverlap uint32       `json:"firstOverlap"`
	StartedAt    time.Time    `json:"startedAt"`
	FinishedAt   time.Time    `json:"finishedAt"`
}

func (t *Task) mapFromModel(row taskRow) *Task {
	var answers []UserAnswer
	for _, answer := range row.Answers {
		answers = append(answers, UserAnswer{
			answer.UserId,
			answer.Answer,
		})
	}

	t.ID = row.ID
	t.Name = row.Name
	t.Description = row.Description
	t.TaskGroupID = row.TaskGroupID
	t.CustomerID = row.CustomerId
	t.Answers = answers
	t.Answer = row.Answer
	t.Overlap = row.Overlap
	t.FirstOverlap = row.FirstOverlap
	t.StartedAt = row.StartedAt
	t.FinishedAt = row.FinishedAt

	return t
}

func (t *Task) mapToModel() taskRow {
	var answers []userAnswerRow
	for _, answer := range t.Answers {
		answers = append(answers, userAnswerRow{
			answer.UserID,
			answer.Answer,
		})
	}

	return taskRow{
		ID:           t.ID,
		Name:         t.Name,
		Description:  t.Description,
		TaskGroupID:  t.TaskGroupID,
		CustomerId:   t.CustomerID,
		Answers:      answers,
		Answer:       t.Answer,
		Overlap:      t.Overlap,
		FirstOverlap: t.FirstOverlap,
		StartedAt:    t.StartedAt,
		FinishedAt:   t.FinishedAt,
	}
}

func mapToModels(tasks []Task) []taskRow {
	taskRows := make([]taskRow, len(tasks))
	for _, task := range tasks {
		taskRows = append(taskRows, task.mapToModel())
	}

	return taskRows
}
