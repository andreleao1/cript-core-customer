package services

import (
	repositories "core-customer/api/infra/repositories/impl"
	"core-customer/domain/entities"
	"log/slog"
	"strconv"
)

type BalanceReserveService struct {
	BalanceReserveRepository repositories.BalanceReserveRepository
}

func NewBalanceReservationService(balanceReserveRepository repositories.BalanceReserveRepository) BalanceReserveService {
	return BalanceReserveService{BalanceReserveRepository: balanceReserveRepository}
}

func (b *BalanceReserveService) ReserveBalance(reserve *entities.BalanceReserve) error {
	slog.Info("Initiating balance reserve to wallet " + reserve.WalletId)
	var currentBalanceFloat float64
	var reservedBalanceFloat float64

	walletRespository := repositories.NewWalletRepository(b.BalanceReserveRepository.Db)

	currentBalance, err := walletRespository.GetBalance(reserve.WalletId)

	if err != nil {
		return err
	}

	walletRespository.ApplyExclusiveLock(reserve.WalletId)

	err = b.BalanceReserveRepository.ReserveBalance(reserve)

	if err != nil {
		return err
	}

	reservedBalanceFloat, err = parseToFloat(reserve.Amount)

	if err != nil {
		return err
	}

	currentBalanceFloat, err = parseToFloat(currentBalance)

	if err != nil {
		return err
	}

	newWalletBalance := currentBalanceFloat - reservedBalanceFloat

	err = walletRespository.UpdateBalance(reserve.WalletId, strconv.FormatFloat(newWalletBalance, 'f', -1, 64))

	if err != nil {
		return err
	}

	return nil
}

func parseToFloat(value string) (float64, error) {
	newValue, err := strconv.ParseFloat(value, 64)

	if err != nil {
		slog.Error("Error parsing value to float.")
		return 0, err
	}

	return newValue, nil
}
