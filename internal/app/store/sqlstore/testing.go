package sqlstore

import (
	"database/sql"
	"fmt"
	"github.com/dev-tim/message-board-api/internal/app/db/sqldb"
	"github.com/sirupsen/logrus/hooks/test"
	"strings"
	"testing"
)

func TestDB(t *testing.T, databaseUrl string) (*sql.DB, func(...string)) {
	t.Helper()

	db, err := sql.Open("postgres", databaseUrl)
	if err != nil {
		t.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}

	logger, _ := test.NewNullLogger()
	if err := sqldb.Migrate(db, &sqldb.Config{
		DbUrl:          databaseUrl,
		CurrentVersion: 1,
	}, logger); err != nil {
		t.Fatal(err)
	}

	return db, func(tables ...string) {
		if len(tables) > 0 {
			if _, err := db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE ", strings.Join(tables, ", "))); err != nil {
				t.Fatal(err)
			}
		}
	}
}
