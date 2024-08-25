package services

import (
	"core-customer/api/dto/in"
	"core-customer/api/infra/db"
	"core-customer/api/infra/repositories"
	repositoriesImpl "core-customer/api/infra/repositories/impl"
	"core-customer/domain/entities"
	"errors"
	"log/slog"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type CustomerService struct {
	CustomerRepository repositories.CustomerRepository
}

func NewcustomerService(customerRepository repositories.CustomerRepository) CustomerService {
	return CustomerService{customerRepository}
}

func (c *CustomerService) CreateCustomer(customerIn in.CustomerInDTO) error {
	slog.Info("Initiating customer creation")
	newDbConnection := db.OpenConnection()
	slog.Info("Creating new transaction")
	newTransaction, err := newDbConnection.Beginx()

	if err != nil {
		slog.Error("Error to create transaction: %v", err.Error(), "")
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			newTransaction.Rollback()
			panic(p)
		} else if err != nil {
			newTransaction.Rollback()
		} else {
			err = newTransaction.Commit()
			if err != nil {
				slog.Error("Failed to commit transaction: %v", err.Error(), "")
			}
		}
	}()

	c.CustomerRepository = repositoriesImpl.NewCustomerRepository(newTransaction)

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

	err = c.CustomerRepository.Create(&customerCreated)

	if err != nil {
		slog.Error("Error to create customer: %v", err.Error(), "")
		return err

	}

	err = createWalletToNewCustomer(customerCreated.Id.String(), newTransaction)

	if err != nil {
		slog.Error("Error to create wallet: %v", err.Error(), "")
		return err
	}

	slog.Info("Customer created successfully id: %s", customerCreated.Id.String(), "")

	return nil
}

func verifyTypedPassword(typedPassword, reTypedPassword string) bool {
	return typedPassword == reTypedPassword
}

func createWalletToNewCustomer(customerId string, transaction *sqlx.Tx) error {
	walletRepo := repositoriesImpl.NewWalletRepository(transaction)

	err := walletRepo.CreateWallet(entities.NewWallet(uuid.MustParse(customerId)))

	if err != nil {
		slog.Error("Error to create wallet: %v", err.Error(), "")
		return errors.New("error to create wallet")
	}

	return nil
}
