package repositories

import "core-customer/core-customer/src/domain/entities"

type CustomerRepository interface {
	Create(customer *entities.Customer) error
}
