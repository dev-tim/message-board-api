package store

import (
	"fmt"
	"github.com/dev-tim/message-board-api/internal/app/common"
	"strings"
	"testing"
)

func TestStore(t *testing.T, databaseUrl string) (*Store, func(...string)) {
	t.Helper()

	_, _ = common.NewLoggerFactory(common.NewLoggerConfig())

	config := NewConfig()
	config.DbUrl = databaseUrl
	config.CurrentVersion = 1

	s := New(config)
	if err := s.Open(); err != nil {
		t.Fatal("Could not open db connection", err)
	}

	if err := s.Migrate(); err != nil {
		t.Fatal("Could not apply migrations", err)
	}

	return s, func(tables ...string) {
		if len(tables) > 0 {
			if _, err := s.db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE ", strings.Join(tables, ","))); err != nil {
				t.Fatal(err)
			}
		}
	}
}
