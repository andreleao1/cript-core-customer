package repositories

import (
	"context"
	"core-customer/domain/entities"
	"log/slog"

	"github.com/jmoiron/sqlx"
)

type CustomerRepository struct {
	Db sqlx.ExtContext
}

func NewCustomerRepository(db sqlx.ExtContext) *CustomerRepository {
	return &CustomerRepository{Db: db}
}

func (r *CustomerRepository) Create(customer *entities.Customer) error {
	_, err := r.Db.ExecContext(context.Background(), "INSERT INTO customers (id, name, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)",
		customer.Id,
		customer.Name,
		customer.Email,
		customer.Password,
		customer.CreatedAt,
		customer.UpdatedAt)

	if err != nil {
		slog.Error("Error creating customer: %v", err)
		return err
	}

	slog.Info("New customer created with id: " + customer.Id.String())
	return nil
}
