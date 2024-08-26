package services

import (
	repositories "core-customer/api/infra/repositories/impl"
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
	slog.Info("Initiating wallet creation to user " + wallet.CustomerId.String())

	err := (*w.WalletRepository).CreateWallet(wallet)

	if err != nil {
		slog.Error("Error to create wallet: " + err.Error())
		return err
	}

	return nil
}

func (w *WalletService) GetWalletByCustomerId(customerId string) (entities.Wallet, error) {
	slog.Info("Initiating wallet search to user " + customerId)

	wallet, err := (*w.WalletRepository).GetWalletByCustomerId(customerId)

	if err != nil {
		slog.Error("Error to get wallet: " + err.Error())
		return entities.Wallet{}, err
	}

	slog.Info("Returning wallet to user " + customerId)

	return wallet, nil
}
