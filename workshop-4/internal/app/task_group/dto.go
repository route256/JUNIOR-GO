package task_group

type TaskGroup struct {
	ID              uint32 `json:"id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	Price           uint32 `json:"price"`
	SecondsToDecide uint32 `json:"secondsToDecide"`
}

func (t *TaskGroup) mapFromModel(row taskGroupRow) *TaskGroup {
	t.ID = row.ID
	t.Name = row.Name
	t.Description = row.Description
	t.Price = row.Price
	t.SecondsToDecide = row.SecondsToDecide

	return t
}

func (t *TaskGroup) mapToModel() taskGroupRow {
	return taskGroupRow{
		ID:              t.ID,
		Name:            t.Name,
		Description:     t.Description,
		Price:           t.Price,
		SecondsToDecide: t.SecondsToDecide,
	}
}
