package repositories

import (
	"context"
	"core-customer/domain/entities"
	"log/slog"

	"github.com/jmoiron/sqlx"
)

type BalanceReserveRepository struct {
	Db sqlx.ExtContext
}

func NewBalanceReserveRepository(db sqlx.ExtContext) *BalanceReserveRepository {
	return &BalanceReserveRepository{Db: db}
}

func (r *BalanceReserveRepository) ReserveBalance(reserve *entities.BalanceReserve) error {
	slog.Info("Executing query to reserve balance.")
	query :=
		`INSERT INTO
		balance_reserves (id, wallet_id, amount, status, created_at, updated_at)
	VALUES
		($1, $2, $3, $4, $5, $6)`
	_, err := r.Db.ExecContext(
		context.Background(),
		query,
		reserve.Id,
		reserve.WalletId,
		reserve.Amount,
		reserve.Status,
		reserve.CreatedAt,
		reserve.UpdatedAt,
	)

	if err != nil {
		slog.Error("Error reserving balance.")
		return err
	}

	slog.Info("Balance reserved with id: " + reserve.Id.String())

	return nil
}
