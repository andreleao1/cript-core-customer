package repositories

import (
	"core-customer/domain/entities"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *entities.User) error {
	_, err := r.db.NamedExec("INSERT INTO customers (id, name) VALUES (:id, :name)", user)
	return err
}
