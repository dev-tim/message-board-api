package store_test

import (
	"fmt"
	"os"
	"testing"
)

var (
	databaseUrl string
)

func TestMain(m *testing.M) {
	databaseUrl = os.Getenv("TEST_DATABASE_URL")

	if databaseUrl == "" {
		databaseUrl = fmt.Sprintf(
			"host=%s dbname=%s user=%s password=%s sslmode=disable",
			"localhost", "test-messages-db", "postgres", "guessme")
	}

	os.Exit(m.Run())
}
