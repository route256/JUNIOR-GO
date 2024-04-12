package task_group

type taskGroupRow struct {
	ID              uint32 `db:"id"`
	Name            string `db:"name"`
	Description     string `db:"description"`
	Price           uint32 `db:"price"`
	SecondsToDecide uint32 `db:"seconds_to_decide"`
}
