package rating

type ratingRow struct {
	UserID        uint32 `db:"user_id"`
	SimpleValue   uint32 `db:"simple_value"`
	MultiplyValue uint32 `db:"multiply_value"`
	AdditionValue uint32 `db:"addition_value"`
}
