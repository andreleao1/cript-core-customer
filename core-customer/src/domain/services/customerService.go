package services

import (
	"core-customer/api/dto/in"
	"core-customer/api/infra/repositories"
	"core-customer/domain/entities"
	"log/slog"
)

type CustomerService struct {
	customerRepository repositories.CustomerRepository
}

func NewcustomerService(customerRepository repositories.CustomerRepository) CustomerService {
	return CustomerService{customerRepository}
}

func (c *CustomerService) CreateCustomer(customerIn in.CustomerInDTO) error {
	slog.Info("Initiating customer creation", "", "")

	customerCreated, err := entities.NewCustomer(customerIn.Name, customerIn.Email, customerIn.Password)

	if err != nil {
		slog.Error("Error to create customer: %v", err.Error(), "")
		return err
	}

	c.customerRepository.Create(&customerCreated)

	slog.Info("Customer created successfully id: %s", customerCreated.Id.String(), "")

	return nil
}
