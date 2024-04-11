//go:build integration
// +build integration

package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/workshop3/workshop-3/internal/pkg/repository/postgresql"
	"gitlab.ozon.dev/workshop3/workshop-3/tests/fixtures"
)

func TestCreateArticle(t *testing.T) {
	var (
		ctx = context.Background()
	)
	t.Run("success", func(t *testing.T) {
		db.SetUp(t)
		defer db.TearDown()
		//arrange
		repo := postgresql.NewArticles(db.DB)

		//act
		resp, err := repo.Add(ctx, fixtures.Article().Valid().P())

		//assert
		require.NoError(t, err)
		assert.NotZero(t, resp)
	})
}

func TestGetArticle(t *testing.T) {
	var (
		ctx          = context.Background()
		articleValid = fixtures.Article().Valid().P()
	)
	t.Run("success", func(t *testing.T) {
		db.SetUp(t, articleValid)
		defer db.TearDown()
		//arrange
		repo := postgresql.NewArticles(db.DB)
		resp, err := repo.Add(ctx, articleValid)

		require.NoError(t, err)
		assert.NotZero(t, resp)
		//act
		respGet, err := repo.GetByID(ctx, resp)

		//assert
		require.NoError(t, err)
		assert.Equal(t, articleValid.Name, respGet.Name)
		assert.Equal(t, articleValid.Rating, respGet.Rating)
	})
}
