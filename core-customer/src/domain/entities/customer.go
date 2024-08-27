package entities

import (
	passwordUtil "core-customer/core-customer/src/utils/password"
	"log/slog"
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	Id        uuid.UUID `db:"id" json:"id"`
	Name      string    `db:"name" json:"name" binding:"required"`
	Email     string    `db:"email" json:"email" binding:"required"`
	Password  string    `db:"password" json:"password" binding:"required"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

func NewCustomer(name string, email string, password string) (Customer, error) {
	hashedPassword, err := hashPassword(password)
	if err != nil {
		slog.Error(err.Error())
		return Customer{}, err
	}

	return Customer{
		Id:        uuid.New(),
		Name:      name,
		Email:     email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := passwordUtil.HashPassword(password)
	if err != nil {
		slog.Error("Error hashing password.")
		return "", err
	}
	return hashedPassword, nil
}
