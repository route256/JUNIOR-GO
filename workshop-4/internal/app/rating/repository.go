package rating

import (
	"context"
	"fmt"
)

type Connector interface {
}

type postgresRepository struct {
	connector Connector
}

func NewRepository(connector Connector) *postgresRepository {
	return &postgresRepository{
		connector: connector,
	}
}

func (r postgresRepository) get(ctx context.Context, userID uint32) (uint32, error) {
	return 51, nil
}

func (r postgresRepository) save(ctx context.Context, ratingRow ratingRow) error {
	fmt.Println(ratingRow)
	return nil
}
