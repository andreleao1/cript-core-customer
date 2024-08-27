package services

import (
	repositories "core-customer/api/infra/repositories/impl"
	"core-customer/domain/entities"
	"log/slog"
	"strconv"

	"github.com/jmoiron/sqlx"
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

	walletRespository, err := applyExclusiveLockToWallet(b.BalanceReserveRepository.Db, reserve.WalletId)
	if err != nil {
		return err
	}

	currentBalance, err := walletRespository.GetBalance(reserve.WalletId)
	if err != nil {
		return err
	}

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

func (b *BalanceReserveService) EffectReserve(reserveId string) error {
	slog.Info("Initiating reserve effect to reserve " + reserveId)
	walletRepository, err := applyExclusiveLockToWallet(b.BalanceReserveRepository.Db, reserveId)
	if err != nil {
		return err
	}

	err = b.BalanceReserveRepository.EffectReserve(reserveId)

	if err != nil {
		return err
	}

	walletId, reserveAmount, err := b.BalanceReserveRepository.GetWalletIdAndReserveAmount(reserveId)
	if err != nil {
		return err
	}

	currentBalanceInvested, err := walletRepository.GetBalanceInvested(walletId)
	if err != nil {
		return err
	}

	reserveAmountFloat, err := parseToFloat(reserveAmount)
	if err != nil {
		return err
	}

	currentBalanceInvestedFloat, err := parseToFloat(currentBalanceInvested)
	if err != nil {
		return err
	}

	newBalanceInvested := currentBalanceInvestedFloat + reserveAmountFloat

	err = walletRepository.UpdateBalanceInvested(walletId, strconv.FormatFloat(newBalanceInvested, 'f', -1, 64))
	if err != nil {
		return err
	}

	return nil
}

func applyExclusiveLockToWallet(transaction sqlx.ExtContext, walletId string) (*repositories.WalletRepository, error) {
	walletRepository := repositories.NewWalletRepository(transaction)
	err := walletRepository.ApplyExclusiveLock(walletId)
	if err != nil {
		return nil, err
	}

	return walletRepository, nil
}

func parseToFloat(value string) (float64, error) {
	newValue, err := strconv.ParseFloat(value, 64)

	if err != nil {
		slog.Error("Error parsing value to float.")
		return 0, err
	}

	return newValue, nil
}
