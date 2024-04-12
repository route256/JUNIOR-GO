package answer

type Answer struct {
	ID     int
	userID int
	sum    int
}

type PgRepository struct {
	connector string
}

func NewRepository(connector string) *PgRepository {
	return &PgRepository{
		connector: connector,
	}
}

func (r *PgRepository) getAnswer(ID int) *Answer {
	return &Answer{
		ID,
		300,
		500,
	}
}

func (r *PgRepository) verifyAnswer(_ int, _ bool) error {
	return nil
}
