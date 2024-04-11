package controller

import (
	"context"
	"time"

	"gitlab.ozon.dev/go/classroom-9/students/homework-7/internal/datasource"
)

type Client struct {
	source datasource.Datasource
}

func NewClient(source datasource.Datasource) *Client {
	return &Client{source: source}
}

func (c *Client) Set(
	ctx context.Context,
	key string,
	value any,
	expiration time.Duration,
) error {
	return c.source.Set(ctx, key, value, expiration)
}

func (c *Client) Get(ctx context.Context, key string) (any, error) {
	return c.source.Get(ctx, key)
}
