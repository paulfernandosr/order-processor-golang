package service

import (
	"context"

	"github.com/paulfernandosr/order-processor-golang/customer-service/internal/model"
	"github.com/paulfernandosr/order-processor-golang/customer-service/internal/repository"
)

type CustomerService interface {
	CreateCustomer(context.Context, *model.Customer) error
	GetAllCustomers(context.Context) ([]model.Customer, error)
	GetCustomersByIds(context.Context, []string) ([]model.Customer, error)
	GetCustomerById(context.Context, string) (*model.Customer, error)
	UpdateCustomerById(context.Context, string, *model.Customer) error
	DeleteCustomerById(context.Context, string) error
}

type CustomerServiceImpl struct {
	repository repository.CustomerRepository
}

func NewCustomerService(repository repository.CustomerRepository) CustomerService {
	return &CustomerServiceImpl{repository}
}

func (service *CustomerServiceImpl) CreateCustomer(ctx context.Context, customer *model.Customer) error {
	return service.repository.InsertOneCustomer(ctx, customer)
}

func (service *CustomerServiceImpl) GetAllCustomers(ctx context.Context) ([]model.Customer, error) {
	return service.repository.FindAllCustomers(ctx)
}

func (service *CustomerServiceImpl) GetCustomersByIds(ctx context.Context, customerIds []string) ([]model.Customer, error) {
	return service.repository.FindCustomersByIds(ctx, customerIds)
}

func (service *CustomerServiceImpl) GetCustomerById(ctx context.Context, customerId string) (*model.Customer, error) {
	return service.repository.FindOneCustomerById(ctx, customerId)
}

func (service *CustomerServiceImpl) UpdateCustomerById(ctx context.Context, customerId string, customer *model.Customer) error {
	return service.repository.UpdateOneCustomerById(ctx, customerId, customer)
}

func (service *CustomerServiceImpl) DeleteCustomerById(ctx context.Context, customerId string) error {
	return service.repository.DeleteOneCustomerById(ctx, customerId)
}
