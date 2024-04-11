// TODO Вы можете редактировать этот файл по вашему усмотрению

package pkg

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type Cache struct {
}

func (c *Cache) Get(ctx context.Context, key string) *redis.StringCmd {
	//TODO implement me
	panic("implement me")
}

func (c *Cache) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}
