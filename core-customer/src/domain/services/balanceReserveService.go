package services

import (
	repositories "core-customer/api/infra/repositories/impl"
	"core-customer/domain/entities"
	"log/slog"
)

type BalanceReserveService struct {
	BalanceReserveRepository repositories.BalanceReserveRepository
}

func NewBalanceReservationService(balanceReserveRepository repositories.BalanceReserveRepository) BalanceReserveService {
	return BalanceReserveService{BalanceReserveRepository: balanceReserveRepository}
}

func (b *BalanceReserveService) ReserveBalance(reserve *entities.BalanceReserve) error {
	slog.Info("Initiating balance reserve to wallet " + reserve.WalletId)
	err := b.BalanceReserveRepository.ReserveBalance(reserve)

	if err != nil {
		return err
	}

	return nil
}
