package entities

import (
	"strconv"
	"time"

	"github.com/google/uuid"
)

type ReserveStatus string

const (
	ReserveStatusPending ReserveStatus = "PENDING"
	ReserveStatusDone    ReserveStatus = "DONE"
	ReserveStatusCancel  ReserveStatus = "CANCEL"
)

type BalanceReserve struct {
	Id        uuid.UUID     `json:"id" db:"id"`
	WalletId  string        `json:"walletId" db:"wallet_id"`
	Amount    string        `json:"balance" db:"amount"`
	Status    ReserveStatus `json:"status" db:"status"`
	CreatedAt time.Time     `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time     `json:"updatedAt" db:"updated_at"`
}

func NewBalanceReserve(walletId string, amount float64) BalanceReserve {
	return BalanceReserve{
		Id:        uuid.New(),
		WalletId:  walletId,
		Amount:    strconv.FormatFloat(amount, 'f', 2, 64),
		Status:    ReserveStatusPending,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
