package repositories

import (
	"core-customer/domain/entities"
	"log/slog"

	"github.com/jmoiron/sqlx"
)

type CustomerRepository struct {
	db *sqlx.DB
}

func NewCustomerRepository(db *sqlx.DB) *CustomerRepository {
	return &CustomerRepository{db: db}
}

func (r *CustomerRepository) Create(customer *entities.Customer) {
	r.db.MustExec("INSERT INTO customers (id, name, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)",
		customer.Id,
		customer.Name,
		customer.Email,
		customer.Password,
		customer.CreatedAt,
		customer.UpdatedAt)

	slog.Info("New customer created with id: %s", customer.Id.String(), "")
}
