package controllers

import (
	"core-customer/domain/entities"
	"core-customer/domain/services"
)

type WalletController struct {
	WalletService services.WalletService
}

func NewWalletController(walletService services.WalletService) WalletController {
	return WalletController{WalletService: walletService}
}

func (w *WalletController) GetWalletByCustomerId(customerId string) (entities.Wallet, error) {
	waller, err := w.WalletService.GetWalletByCustomerId(customerId)

	if err != nil {
		return entities.Wallet{}, err
	}

	return waller, nil
}
