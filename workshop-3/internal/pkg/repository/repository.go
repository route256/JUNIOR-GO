//go:generate mockgen -source ./repository.go -destination=./mocks/repository.go -package=mock_repository
package repository

import "context"

type ArticlesRepo interface {
	Add(ctx context.Context, article *Article) (int64, error)
	GetByID(ctx context.Context, id int64) (*Article, error)
}
