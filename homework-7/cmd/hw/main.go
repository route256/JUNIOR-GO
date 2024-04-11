package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"gitlab.ozon.dev/go/classroom-9/students/homework-7/internal/controller"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	client := controller.NewClient() // TODO

	var bestUser = "best user, expired after 30 seconds"

	// Создаём запись
	err := client.Set(ctx, "user:12345:profile", bestUser, 1)
	if err != nil {
		panic(err)
	}

	// Получаем запись из кэша
	got, err := client.Get(ctx, "user:12345:profile")
	if err != nil {
		panic(err)
	}

	if got != bestUser {
		panic("invalid value")
	}

	select {
	case <-time.After(2 * time.Second):
	case <-ctx.Done():
	}

	// Получаем запись из базы данных и обновляем кэщ
	gotAgain, err := client.Get(ctx, "user:12345:profile")
	if err != nil {
		panic(err)
	}

	if gotAgain != bestUser {
		panic("invalid value")
	}
}
