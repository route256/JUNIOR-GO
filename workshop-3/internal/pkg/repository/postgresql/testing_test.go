package postgresql

import (
	"testing"

	"github.com/golang/mock/gomock"
	mock_database "gitlab.ozon.dev/workshop3/workshop-3/internal/pkg/db/mocks"
	"gitlab.ozon.dev/workshop3/workshop-3/internal/pkg/repository"
)

type articlesRepoFixture struct {
	ctrl   *gomock.Controller
	repo   repository.ArticlesRepo
	mockDb *mock_database.MockDBops
}

func setUp(t *testing.T) articlesRepoFixture {
	ctrl := gomock.NewController(t)
	mockDb := mock_database.NewMockDBops(ctrl)
	repo := NewArticles(mockDb)
	return articlesRepoFixture{
		ctrl:   ctrl,
		repo:   repo,
		mockDb: mockDb,
	}
}

func (a *articlesRepoFixture) tearDown() {
	a.ctrl.Finish()
}
