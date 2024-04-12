package datasource

import (
	"context"
	"time"
)

type Datasource interface {
	Set(ctx context.Context, key string, value any, expiration time.Duration) error
	Get(ctx context.Context, key string) (any, error)
}
