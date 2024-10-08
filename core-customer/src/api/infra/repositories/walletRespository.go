package repositories

import (
	"core-customer/core-customer/src/domain/entities"
)

type WalletRepository interface {
	CreateWallet(wallet *entities.Wallet) error
	GetWalletByCustomerId(customerId string) (entities.Wallet, error)
	ApplyExclusiveLock(walletId string) error
	UpdateWalletBalance(walletId string, balance string) error
	GetBalance(walletId string) (string, error)
	GetBalanceInvested(walletId string) (string, error)
	UpdateBalanceInvested(walletId string, balanceInvested string) error
}
