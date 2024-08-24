package repositories

import (
	"core-customer/domain/entities"
	"log/slog"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *entities.User) {
	slog.Info("Creating user %v", user)
	r.db.MustExec("INSERT INTO customer (id, name) VALUES ($1, $2)", user.Id, user.Name)
	slog.Info("User created %v", user)
}
