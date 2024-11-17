package wallet

import (
	"context"
	"fmt"
)

type service struct {
	storage Storage
}

func NewService(storage Storage) Service {
	return &service{storage: storage}
}

func (s *service) GetWallets(ctx context.Context) ([]Wallet, error) {
	wallets, err := s.storage.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return wallets, nil
}

func (s *service) GetWalletByUUID(ctx context.Context, uuid string) (Wallet, error) {
	w, err := s.storage.GetOne(ctx, uuid)
	if err != nil {
		return Wallet{}, err
	}
	return w, nil
}

func (s *service) ChangeBalance(ctx context.Context, update UpdateWalletRequest) (Wallet, error) {
	w, err := s.storage.GetOne(ctx, update.UUID)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return w, fmt.Errorf("wallet not found")
		}
		return w, err
	}
	if update.OperationType == withdraw {
		if w.Balance-update.Amount < 0 {
			return w, fmt.Errorf("not enough balance")
		}
		w.Balance -= update.Amount
		err = s.storage.UpdateOne(ctx, w)
		if err != nil {
			return w, err
		}
		return w, nil
	}

	w.Balance += update.Amount
	err = s.storage.UpdateOne(ctx, w)
	if err != nil {
		return w, err
	}
	return w, nil
}
