package repositories

import (
	"context"
	"core-customer/domain/entities"
	"log/slog"

	"github.com/jmoiron/sqlx"
)

type WalletRepository struct {
	db sqlx.ExtContext
}

func NewWalletRepository(db sqlx.ExtContext) *WalletRepository {
	return &WalletRepository{db: db}
}

func (w *WalletRepository) CreateWallet(wallet entities.Wallet) error {
	_, err := w.db.ExecContext(context.Background(), "INSERT INTO wallets (id, customer_id, balance, balance_invested, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)",
		wallet.Id,
		wallet.CustomerId,
		wallet.Balance,
		wallet.BalanceInvested,
		wallet.CreatedAt,
		wallet.UpdatedAt)

	if err != nil {
		slog.Error("Error creating wallet: %v", err)
		return err
	}

	slog.Info("New customer wallet with id: %s", wallet.Id.String(), "")
	return nil
}
