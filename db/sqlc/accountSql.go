package db

import (
	"context"
	"database/sql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

const (
	DbSource = "postgres://postgres:ncell@localhost:5430/postgres?sslmode=disable"
)

func connectToDatabase() *sql.DB {
	return sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(DbSource)))
}

type CreateAccountParams struct {
	ID       int64  `bun:"id,pk"`
	Owner    string `bun:"owner,notnull"`
	Balance  int64  `bun:"balance,notnull"`
	Currency string `bun:"currency,notnull"`
}

type UpdateAccountParams struct {
	ID      int64
	Balance int64
}

func CreateAccount(ctx context.Context, arg CreateAccountParams) (*Accounts, error) {
	db, _ := connectAndCreateAccount()
	accounts := new(Accounts)
	accounts.Owner = arg.Owner
	accounts.Balance = arg.Balance
	accounts.Currency = arg.Currency
	_, err := db.NewInsert().Model(accounts).Exec(ctx)
	return accounts, err
}

func connectAndCreateAccount() (*bun.DB, *Accounts) {
	pgdb := connectToDatabase()
	db := bun.NewDB(pgdb, pgdialect.New())
	accounts := new(Accounts)
	return db, accounts
}

func GetAccount(ctx context.Context, id int64) (Accounts, error) {
	db, accounts := connectAndCreateAccount()
	err := db.NewSelect().Model(accounts).Where("id = ?", id).Scan(ctx, accounts)
	return *accounts, err
}

func DeleteAccount(ctx context.Context, id int64) error {
	db, accounts := connectAndCreateAccount()
	_, err := db.NewDelete().Model(accounts).Where("id = ?", id).Exec(ctx)
	return err
}

func UpdateAccount(ctx context.Context, arg UpdateAccountParams, argRand *Accounts) (*Accounts, error) {
	db, _ := connectAndCreateAccount()
	accounts := new(Accounts)
	accounts.ID = arg.ID
	accounts.Owner = argRand.Owner
	accounts.Balance = arg.Balance
	accounts.Currency = argRand.Currency
	_, err := db.NewUpdate().Model(accounts).WherePK().Exec(ctx)
	return accounts, err
}
