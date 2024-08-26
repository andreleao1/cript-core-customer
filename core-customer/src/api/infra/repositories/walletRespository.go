package repositories

import (
	"core-customer/domain/entities"
)

type WalletRepository interface {
	CreateWallet(wallet *entities.Wallet) error
	GetWalletByCustomerId(customerId string) (entities.Wallet, error)
}
