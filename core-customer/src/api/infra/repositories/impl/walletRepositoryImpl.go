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
		slog.Error("Error creating wallet.")
		return err
	}

	slog.Info("New customer wallet with id: " + wallet.Id.String())
	return nil
}

func (w *WalletRepository) GetWalletByCustomerId(customerId string) (entities.Wallet, error) {
	var wallet entities.Wallet
	slog.Info("Executing query")

	query := `
	SELECT 
		id,
		customer_id, 
		balance, 
		balance_invested, 
		created_at, 
		updated_at 
	FROM 
		wallets 
	WHERE 
		customer_id = $1`

	row := w.db.QueryRowxContext(
		context.Background(),
		query,
		customerId)

	err := row.Scan(
		&wallet.Id,
		&wallet.CustomerId,
		&wallet.Balance,
		&wallet.BalanceInvested,
		&wallet.CreatedAt,
		&wallet.UpdatedAt,
	)

	if err != nil {
		slog.Error("Error getting wallet.")
		return entities.Wallet{}, err
	}

	slog.Info("Query executed successfully")

	return wallet, nil
}
