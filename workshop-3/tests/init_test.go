//go:build integration
// +build integration

package tests

import "gitlab.ozon.dev/workshop3/workshop-3/tests/postgres"

var (
	db *postgres.TDB
)

func init() {
	//тут мы запрашиваем тестовые креды  для бд из енв
	// cfg,err := config.FromEnv()

	db = postgres.NewFromEnv()
}
