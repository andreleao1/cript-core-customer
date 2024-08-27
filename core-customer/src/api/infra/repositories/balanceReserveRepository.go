package repositories

import "core-customer/core-customer/src/domain/entities"

type BalanceReserveRepository interface {
	ReserveBalance(reserva *entities.BalanceReserve) error
	EffectReserve(reserveId string) error
	CancelReserve(reserveId string) error
	GetWalletIdAndReserveAmount(reserveId string) (string, error)
}
