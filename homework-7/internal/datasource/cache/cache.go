// TODO Вы можете редактировать этот файл по вашему усмотрению

package cache

import (
	"context"
	"time"
)

type Client struct {
}

func (c *Client) Set(ctx context.Context, key string, value any, expiration time.Duration) error {
	//TODO implement me
	panic("implement me")
}

func (c *Client) Get(ctx context.Context, key string) (any, error) {
	//TODO implement me
	panic("implement me")
}
