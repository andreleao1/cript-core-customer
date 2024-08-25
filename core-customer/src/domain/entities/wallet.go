package entities

import (
	"time"

	"github.com/google/uuid"
)

type Wallet struct {
	Id              uuid.UUID `json:"id"`
	CustomerId      uuid.UUID `json:"customer_id"`
	Balance         string    `json:"balance"`
	BalanceInvested string    `json:"balance_invested"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
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
