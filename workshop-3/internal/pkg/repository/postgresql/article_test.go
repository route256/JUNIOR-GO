package postgresql

import (
	"context"
	"database/sql"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestArticles_GetByID(t *testing.T) {
	t.Parallel()
	var (
		ctx = context.Background()
		id  = 1
	)
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		// arrange

		s := setUp(t)
		defer s.tearDown()
		s.mockDb.EXPECT().Get(gomock.Any(), gomock.Any(), "SELECT id,name,rating,created_at FROM articles WHERE id=$1", gomock.Any()).Return(nil)

		// act
		user, err := s.repo.GetByID(ctx, int64(id))
		// assert

		require.NoError(t, err)
		assert.Equal(t, int64(0), user.ID)
	})
	t.Run("fail", func(t *testing.T) {
		t.Parallel()
		t.Run("not found", func(t *testing.T) {
			t.Parallel()
			// arrange
			s := setUp(t)
			defer s.tearDown()

			s.mockDb.EXPECT().Get(gomock.Any(), gomock.Any(), "SELECT id,name,rating,created_at FROM articles WHERE id=$1", gomock.Any()).Return(sql.ErrNoRows)

			// act
			user, err := s.repo.GetByID(ctx, int64(id))
			// assert
			require.EqualError(t, err, "object not found")

			assert.Nil(t, user)
		})
		t.Run("internal error", func(t *testing.T) {
			t.Parallel()
			// arrange
			s := setUp(t)
			defer s.tearDown()

			s.mockDb.EXPECT().Get(gomock.Any(), gomock.Any(), "SELECT id,name,rating,created_at FROM articles WHERE id=$1", gomock.Any()).Return(assert.AnError)

			// act
			user, err := s.repo.GetByID(ctx, int64(id))
			// assert
			require.EqualError(t, err, "assert.AnError general error for testing")

			assert.Nil(t, user)
		})
	})

}
