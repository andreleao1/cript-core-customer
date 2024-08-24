package repositories

import "core-customer/domain/entities"

type UserRepository interface {
	Create(user *entities.User)
}
