package wallet

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type PsqlClient interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

type Storage interface {
	GetAll(ctx context.Context) ([]Wallet, error)
	GetOne(ctx context.Context, uuid string) (Wallet, error)
	UpdateOne(ctx context.Context, wallet Wallet) error
}

type Service interface {
	GetWallets(ctx context.Context) ([]Wallet, error)
	GetWalletByUUID(ctx context.Context, uuid string) (Wallet, error)
	ChangeBalance(ctx context.Context, update UpdateWalletRequest) (Wallet, error)
}
