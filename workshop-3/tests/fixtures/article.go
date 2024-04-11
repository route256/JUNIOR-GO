package fixtures

import (
	"time"

	"gitlab.ozon.dev/workshop3/workshop-3/internal/pkg/repository"
	"gitlab.ozon.dev/workshop3/workshop-3/tests/states"
)

type ArticleBuilder struct {
	instance *repository.Article
}

func Article() *ArticleBuilder {
	return &ArticleBuilder{instance: &repository.Article{}}
}

func (b *ArticleBuilder) ID(v int64) *ArticleBuilder {
	b.instance.ID = v
	return b
}
func (b *ArticleBuilder) Name(v string) *ArticleBuilder {
	b.instance.Name = v
	return b
}

func (b *ArticleBuilder) Rating(v int64) *ArticleBuilder {
	b.instance.Rating = v
	return b
}

func (b *ArticleBuilder) CreatedAt(v time.Time) *ArticleBuilder {
	b.instance.CreatedAt = v
	return b
}

func (b *ArticleBuilder) P() *repository.Article {
	return b.instance
}

func (b *ArticleBuilder) V() repository.Article {
	return *b.instance
}

func (b *ArticleBuilder) Valid() *ArticleBuilder {
	return Article().ID(states.Article1ID).Name(states.ArticleName1).Rating(1).CreatedAt(time.Time{})
}
