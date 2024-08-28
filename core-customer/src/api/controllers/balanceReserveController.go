package controllers

import (
	"core-customer/core-customer/src/domain/entities"
	"core-customer/core-customer/src/domain/services"
)

type BalanceReserveController struct {
	balanceReserveService services.BalanceReserveService
}

func NewBalanceReserveController(balanceReserceService services.BalanceReserveService) BalanceReserveController {
	return BalanceReserveController{balanceReserceService}
}

func (b *BalanceReserveController) ReserveBalance(reserve *entities.BalanceReserve) (string, error) {
	reserveId, err := b.balanceReserveService.ReserveBalance(reserve)

	if err != nil {
		return "", err
	}

	return reserveId, nil
}

func (b *BalanceReserveController) EffectReserve(reserveId string) error {
	err := b.balanceReserveService.EffectReserve(reserveId)

	if err != nil {
		return err
	}

	return nil
}

func (b *BalanceReserveController) CancelReserve(reserveId string) error {
	err := b.balanceReserveService.CancelReserve(reserveId)

	if err != nil {
		return err
	}

	return nil
}
