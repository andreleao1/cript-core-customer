package repositories

import (
	"core-customer/domain/entities"
)

type WalletRepository interface {
	CreateWallet(wallet *entities.Wallet) error
	GetWalletByCustomerId(customerId string) (entities.Wallet, error)
	ApplyExclusiveLock(walletId string) error
	UpdateWalletBalance(walletId string, balance string) error
	GetBalance(walletId string) (string, error)
}
