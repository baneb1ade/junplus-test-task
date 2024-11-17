package db

import (
	"context"
	"javacode-test-task/app/internal/wallet"
	"log/slog"
)

type Repository struct {
	Client wallet.PsqlClient
	logger *slog.Logger
}

func NewRepository(client wallet.PsqlClient, logger *slog.Logger) wallet.Storage {
	return Repository{client, logger}
}

func (r Repository) GetAll(ctx context.Context) ([]wallet.Wallet, error) {
	const op = "db.psql.GetAll"

	q := `SELECT id, balance
          FROM wallet`
	rows, err := r.Client.Query(ctx, q)
	if err != nil {
		r.logger.Error(op, "error", err)
		return nil, err
	}
	wallets := make([]wallet.Wallet, 0)
	for rows.Next() {
		var w wallet.Wallet
		rows.Scan(&w.UUID, &w.Balance)
		wallets = append(wallets, w)
	}

	if err = rows.Err(); err != nil {
		r.logger.Error(op, "error", err)
		return nil, err
	}
	return wallets, nil
}

func (r Repository) GetOne(ctx context.Context, uuid string) (wallet.Wallet, error) {
	const op = "db.psql.GetOneByUUID"

	q := `SELECT id, balance
          FROM wallet
          WHERE id = $1`
	var w wallet.Wallet
	if err := r.Client.QueryRow(ctx, q, uuid).Scan(&w.UUID, &w.Balance); err != nil {
		r.logger.Error(op, "error", err)
		return w, err
	}
	return w, nil
}

func (r Repository) UpdateOne(ctx context.Context, wallet wallet.Wallet) error {
	const op = "db.psql.UpdateOne"

	q := `UPDATE wallet
		  SET balance = $1
		  WHERE id = $2`
	_, err := r.Client.Exec(ctx, q, wallet.Balance, wallet.UUID)
	if err != nil {
		r.logger.Error(op, "error", err)
		return err
	}
	return nil
}
