package services

import (
	"core-customer/api/infra/repositories"
	"core-customer/domain/entities"
	"log/slog"
)

type WalletService struct {
	WalletRepository *repositories.WalletRepository
}

func NewWalletService(walletRepository *repositories.WalletRepository) WalletService {
	return WalletService{WalletRepository: walletRepository}
}

func (w *WalletService) CreateWallet(wallet entities.Wallet) error {
	slog.Info("Initiating wallet creation to user %s", wallet.CustomerId.String(), "")

	err := (*w.WalletRepository).CreateWallet(&wallet)

	if err != nil {
		slog.Error("Error to create wallet: %v", err.Error(), "")
		return err
	}

	return nil
}
