package migrations_test

import (
	"context"
	"database/sql"
	"fmt"
	"sort"
	"testing"
	"time"

	"github.com/dhui/dktest"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/zikaeroh/strawrank/internal/db/migrations"
	"gotest.tools/assert"
	"gotest.tools/assert/cmp"
)

func TestUp(t *testing.T) {
	t.Parallel()

	withDatabase(t, func(t *testing.T, db *sql.DB) {
		assert.NilError(t, migrations.Up(db, t.Logf))
		assertTableNames(t, db, "polls", "ballots", "schema_migrations")
	})
}

func TestUpDown(t *testing.T) {
	t.Parallel()

	withDatabase(t, func(t *testing.T, db *sql.DB) {
		assert.NilError(t, migrations.Up(db, t.Logf))
		assert.NilError(t, migrations.Down(db, t.Logf))
		assertTableNames(t, db, "schema_migrations")
	})
}

func TestReset(t *testing.T) {
	t.Parallel()

	withDatabase(t, func(t *testing.T, db *sql.DB) {
		assert.NilError(t, migrations.Up(db, t.Logf))
		assertTableNames(t, db, "polls", "ballots", "schema_migrations")
		assert.NilError(t, migrations.Reset(db, t.Logf))
		assertTableNames(t, db, "polls", "ballots", "schema_migrations")
	})
}

func withDatabase(t *testing.T, fn func(t *testing.T, db *sql.DB)) {
	if testing.Short() {
		t.Skip("requires starting a docker container")
	}

	opts := dktest.Options{
		PortRequired:   true,
		ReadyFunc:      postgresReady,
		CleanupTimeout: time.Second,
	}

	dktest.Run(t, "zikaeroh/postgres-initialized", opts,
		func(t *testing.T, c dktest.ContainerInfo) {
			ip, port, err := c.FirstPort()
			assert.NilError(t, err)

			db, err := sql.Open("postgres", connStr(ip, port))
			assert.NilError(t, err)
			defer db.Close()

			assert.NilError(t, db.Ping())

			assertTableNames(t, db)
			fn(t, db)
		})
}

func connStr(ip, port string) string {
	return fmt.Sprintf(`postgres://postgres:mysecretpassword@%s:%s/postgres?sslmode=disable`, ip, port)
}

func postgresReady(ctx context.Context, c dktest.ContainerInfo) bool {
	ip, port, err := c.FirstPort()
	if err != nil {
		return false
	}

	db, err := sql.Open("postgres", connStr(ip, port))
	if err != nil {
		return false
	}

	if err := db.PingContext(ctx); err != nil {
		return false
	}

	return true
}

func assertTableNames(t *testing.T, db *sql.DB, names ...string) {
	t.Helper()
	sort.Strings(names)

	tables := tableNames(t, db)
	sort.Strings(tables)

	assert.Check(t, cmp.DeepEqual(names, tables, cmpopts.EquateEmpty()))
}

func tableNames(t *testing.T, db *sql.DB) []string {
	t.Helper()

	query := `SELECT table_name FROM information_schema.tables WHERE table_schema=(SELECT current_schema()) AND table_type='BASE TABLE'`
	rows, err := db.Query(query)
	assert.NilError(t, err)
	defer rows.Close()

	var names []string
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		assert.NilError(t, err)
		if len(name) > 0 {
			names = append(names, name)
		}
	}

	return names
}
