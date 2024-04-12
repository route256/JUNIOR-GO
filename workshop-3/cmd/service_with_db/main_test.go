package main

import (
	"context"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	mock_repository "gitlab.ozon.dev/workshop3/workshop-3/internal/pkg/repository/mocks"
	"gitlab.ozon.dev/workshop3/workshop-3/tests/fixtures"
)

func Test_Create(t *testing.T) {
	var (
		ctx context.Context
		id  = 1
	)
	t.Run("success", func(t *testing.T) {
		t.Parallel()

		// arrange
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		m := mock_repository.NewMockArticlesRepo(ctrl)
		s := server1{repo: m}
		m.EXPECT().GetByID(gomock.Any(), int64(id)).Return(fixtures.Article().Valid().P(), nil)
		//act
		result, code := s.Get(ctx, int64(id))
		// assert
		require.Equal(t, http.StatusOK, code)
		assert.Equal(t, "{\"ID\":1,\"Name\":\"someName\",\"Rating\":1,\"CreatedAt\":\"0001-01-01T00:00:00Z\"}", string(result))
	})
}

func Test_parseGetID(t *testing.T) {
	type args struct {
		req *http.Request
	}
	tests := []struct {
		name  string
		args  args
		want  int64
		want1 int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := parseGetID(tt.args.req)
			if got != tt.want {
				t.Errorf("parseGetID() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("parseGetID() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
