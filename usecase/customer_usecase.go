package usecase

import (
	"fmt"
	"github.com/jutionck/go-laundry-app-core/model"
	"github.com/jutionck/go-laundry-app-core/repository"
)

type CustomerUseCase interface {
	RegisterNewCustomer(customer *model.Customer) error
	UpdateCustomer(customer *model.Customer) error
	DeleteCustomer(customerId string) error
	FindByCustomerId(customerId string) (*model.Customer, error)
	FindAllCustomer(page int, itemPerPage int, by string, values ...interface{}) ([]model.Customer, error)
}

type customerUseCase struct {
	customerRepo repository.CustomerRepository
}

func (c *customerUseCase) RegisterNewCustomer(customer *model.Customer) error {
	cst, err := c.customerRepo.FindByPhoneNumber(customer.PhoneNumber)
	if err != nil {
		return fmt.Errorf("RegisterNewCustomer : %w", err)
	}

	if cst != nil {
		return fmt.Errorf("customer with phone number %s already exists", customer.PhoneNumber)
	}
	err = c.customerRepo.Create(customer)
	if err != nil {
		return fmt.Errorf("RegisterNewCustomer : %w", err)
	}
	return nil
}

func (c *customerUseCase) UpdateCustomer(customer *model.Customer) error {
	cst, err := c.customerRepo.FindById(customer.Id)
	if err != nil {
		return fmt.Errorf("RegisterNewCustomer : %w", err)
	}

	if cst == nil {
		return fmt.Errorf("customer with ID %s not exists", customer.Id)
	}

	cst, err = c.customerRepo.FindByPhoneNumber(customer.PhoneNumber)
	if err != nil {
		return fmt.Errorf("RegisterNewCustomer : %w", err)
	}

	if cst != nil && cst.Id != customer.Id {
		return fmt.Errorf("customer with phone number %s already exists", customer.PhoneNumber)
	}
	err = c.customerRepo.Update(customer)
	if err != nil {
		return fmt.Errorf("RegisterNewCustomer : %w", err)
	}
	return nil
}

func (c *customerUseCase) DeleteCustomer(customerId string) error {
	cst, err := c.customerRepo.FindById(customerId)
	if err != nil {
		return fmt.Errorf("DeleteCustomer : %w", err)
	}

	if cst == nil {
		return fmt.Errorf("customer with ID %s not exists", customerId)
	}

	err = c.customerRepo.Delete(customerId)
	if err != nil {
		return fmt.Errorf("DeleteCustomer : %w", err)
	}
	return nil
}

func (c *customerUseCase) FindByCustomerId(customerId string) (*model.Customer, error) {
	customer, err := c.customerRepo.FindById(customerId)
	if err != nil {
		return nil, fmt.Errorf("FindByCustomerId : %w", err)
	}
	return customer, nil
}

func (c *customerUseCase) FindAllCustomer(page int, itemPerPage int, by string, values ...interface{}) ([]model.Customer, error) {
	return c.customerRepo.FindAll(page, itemPerPage, by, values)
}

func NewCustomerUseCase(customerRepository repository.CustomerRepository) CustomerUseCase {
	return &customerUseCase{customerRepo: customerRepository}
}
