package entities

import (
	"time"

	"github.com/google/uuid"
)

type Wallet struct {
	Id              uuid.UUID `json:"id"`
	CustomerId      uuid.UUID `json:"customerId"`
	Balance         string    `json:"balance"`
	BalanceInvested string    `json:"balanceInvested"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

func NewWallet(customerId uuid.UUID) Wallet {
	return Wallet{
		Id:              uuid.New(),
		CustomerId:      customerId,
		Balance:         "0.0",
		BalanceInvested: "0.0",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
}
