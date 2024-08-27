package controllers

import (
	"core-customer/domain/entities"
	"core-customer/domain/services"
)

type BalanceReserveController struct {
	balanceReserveService services.BalanceReserveService
}

func NewBalanceReserveController(balanceReserceService services.BalanceReserveService) BalanceReserveController {
	return BalanceReserveController{balanceReserceService}
}

func (b *BalanceReserveController) ReserveBalance(reserve *entities.BalanceReserve) error {
	err := b.balanceReserveService.ReserveBalance(reserve)

	if err != nil {
		return err
	}

	return nil
}
