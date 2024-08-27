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

func (r *BalanceReserveRepository) EffectReserve(reserveId string) error {
	slog.Info("Executing query to effect reserve.")
	query :=
		`UPDATE
		balance_reserves
	SET
		status = $1
	WHERE
		id = $2`
	_, err := r.Db.ExecContext(
		context.Background(),
		query,
		entities.ReserveStatusDone,
		reserveId,
	)

	if err != nil {
		slog.Error("Error effecting balance reserve.")
		return err
	}

	slog.Info("Balance reserve effected with id: " + reserveId)

	return nil
}

func (r *BalanceReserveRepository) GetWalletIdAndReserveAmount(reserveId string) (string, string, error) {
	var walletId string
	var reserveAmount string
	slog.Info("Executing query to get wallet id by reserve id.")

	query := `
	SELECT
		wallet_id,
		amount
	FROM
		balance_reserves
	WHERE
		id = $1`

	row := r.Db.QueryRowxContext(
		context.Background(),
		query,
		reserveId,
	)

	err := row.Scan(&walletId, &reserveAmount)

	if err != nil {
		slog.Error("Error getting wallet id by reserve id.")
		return "", "", err
	}

	slog.Info("Wallet id by reserve id: " + walletId)

	return walletId, reserveAmount, nil
}
