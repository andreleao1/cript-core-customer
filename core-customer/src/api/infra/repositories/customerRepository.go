package repositories

import "core-customer/domain/entities"

type CustomerRepository interface {
	Create(customer *entities.Customer)
}
