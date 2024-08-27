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

func (w *WalletRepository) ApplyExclusiveLock(walletId string) error {
	slog.Info("Applying exclusive lock to wallet with id: " + walletId)

	query := `
	SELECT
		*
	FROM
		wallets
	WHERE
		id = $1
	FOR UPDATE
	`

	_, err := w.db.ExecContext(
		context.Background(),
		query,
		walletId,
	)

	if err != nil {
		slog.Error("Error applying exclusive lock to wallet.")
		return err
	}

	slog.Info("Exclusive lock applied to wallet with id: " + walletId)

	return nil
}

func (w *WalletRepository) UpdateBalance(walletId string, balance string) error {
	slog.Info("Updating wallet balance with id: " + walletId)

	query := `
	UPDATE
		wallets
	SET
		balance = $1,
		updated_at = now()
	WHERE
		id = $2

	`

	_, err := w.db.ExecContext(
		context.Background(),
		query,
		balance,
		walletId,
	)

	if err != nil {
		slog.Error("Error updating wallet balance.")
		return err
	}

	slog.Info("Wallet balance updated with id: " + walletId)

	return nil
}

func (w *WalletRepository) GetBalance(walletId string) (string, error) {
	var balance string
	slog.Info("Getting balance from wallet with id: " + walletId)

	query := `
	SELECT 
		balance
	FROM 
		wallets 
	WHERE 
		id = $1`

	row := w.db.QueryRowxContext(
		context.Background(),
		query,
		walletId)

	err := row.Scan(
		&balance,
	)

	if err != nil {
		slog.Error("Error getting wallet balance.")
		return "", err
	}

	slog.Info("Query executed successfully")

	return balance, nil
}
