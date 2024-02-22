// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"database/sql"
	"github.com/uptrace/bun"
)

type Accounts struct {
	bun.BaseModel `bun:"table:accounts,alias:a"`

	ID        int64        `bun:"id,pk,autoincrement"`
	Owner     string       `bun:"owner,notnull"`
	Balance   int64        `bun:"balance,notnull"`
	Currency  string       `bun:"currency,notnull"`
	CreatedAt sql.NullTime `bun:"created at,default:current_timestamp"`
}
type Entries struct {
	bun.BaseModel `bun:"table:entries,alias:e"`

	ID        int64  `bun:"id,pk"`
	AccountID string `bun:"account_id,notnull"`
	//AccountID string       `bun:"owner,notnull"`
	//Accounts *Accounts `bun:"rel:belongs-to,join:account_id=id"`
	//Accounts *Accounts `bun:"rel:has-one,join:account_id=id"`

	Amount    int64        `bun:"amount,notnull"`
	CreatedAt sql.NullTime `bun:"created at,default:current_timestamp"`
}

type Transfers struct {
	bun.BaseModel `bun:"table:transfers,alias:t"`

	ID 				int64 		`bun:"id,pk"`
	FromAccountID 	int8 		`bun:"from_account_id,notnull"`
	ToAccountId 	int8 		`bun:"to_account_id,notnull"`
	Amount 			int8 		`bun:"to_account_id,notnull"`
	CreatedAt 	sql.NullTime 	`bun:"created at,default:current_timestamp"`
}
