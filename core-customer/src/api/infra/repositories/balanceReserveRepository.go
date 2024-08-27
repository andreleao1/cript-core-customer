package repositories

import "core-customer/domain/entities"

type BalanceReserveRepository interface {
	ReserveBalance(reserva *entities.BalanceReserve) error
	EffectReserve(reserveId string) error
	CancelReserve(reserveId string) error
}
