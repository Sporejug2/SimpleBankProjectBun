package db

import (
	"database/sql"
	"fmt"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"os"
	"testing"
)

var testDb *bun.DB

func TestMain(m *testing.M) {
	pgdb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(DbSource)))
	if pgdb == nil {
		fmt.Println("cannot connect to db:", pgdb)
	}

	// Create a Bun db on top of it.
	testDb = bun.NewDB(pgdb, pgdialect.New())

	os.Exit(m.Run())
}
