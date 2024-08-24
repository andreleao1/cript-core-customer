package services

import (
	"core-customer/api/dto/in"
	"core-customer/api/infra/repositories"
	"core-customer/domain/entities"
	"errors"
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

	isValidTypedPassword := verifyTypedPassword(customerIn.TypedPassword, customerIn.ReTypedPassword)

	if !isValidTypedPassword {
		slog.Error("Typed password and re-typed password are different", "", "")
		return errors.New("typed password and re-typed password are different")
	}

	customerCreated, err := entities.NewCustomer(customerIn.Name, customerIn.Email, customerIn.TypedPassword)

	if err != nil {
		slog.Error("Error to create customer: %v", err.Error(), "")
		return err
	}

	c.customerRepository.Create(&customerCreated)

	slog.Info("Customer created successfully id: %s", customerCreated.Id.String(), "")

	return nil
}

func verifyTypedPassword(typedPassword, reTypedPassword string) bool {
	return typedPassword == reTypedPassword
}
