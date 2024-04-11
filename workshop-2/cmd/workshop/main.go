package main

import (
	"context"
	"os"
	"os/signal"
	"sync/atomic"
	"time"

	"github.com/intamyuto/workshop/logger"
)

type Counter struct {
	value atomic.Value
}

func (c *Counter) Get() int {
	<-time.After(1 * time.Second)

	return c.value.Load().(int)
}

func (c *Counter) Increment() {
	for {
		old := c.value.Load().(int)
		if c.value.CompareAndSwap(old, old+1) {
			break
		}
	}
}

func main() {
	ctx := context.Background()

	ctx = logger.ContextWithTimestamp(ctx, time.Now())

	if err := run(ctx); err != nil {
		logger.Errorf(ctx, "%v", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	var cnt Counter
	cnt.value.Store(0)

	go counter(ctx, &cnt)
	for i := 0; i < 10; i++ {
		go observe(ctx, &cnt)
	}

	<-ctx.Done()
	return ctx.Err()
}

func counter(ctx context.Context, c *Counter) {
	t := time.NewTicker(5 * time.Second)
	defer t.Stop()

	for {
		select {
		case <-t.C:
			c.Increment()
		case <-ctx.Done():
			return
		}
	}
}

func observe(ctx context.Context, c *Counter) {
	t := time.NewTicker(2 * time.Second)
	defer t.Stop()

	for {
		select {
		case <-t.C:
			logger.Infof(ctx, "counter: %d", c.Get())
		case <-ctx.Done():
			return
		}
	}
}
