package controllers

import (
	"core-customer/api/dto/in"
	"core-customer/domain/services"
)

type CustomerController struct {
	customerService services.CustomerService
}

func NewCustomerController(customerService services.CustomerService) CustomerController {
	return CustomerController{customerService}
}

func (u CustomerController) CreateCustomer(customerIn in.CustomerInDTO) error {
	err := u.customerService.CreateCustomer(customerIn)

	if err != nil {
		return err
	}

	return nil
}
