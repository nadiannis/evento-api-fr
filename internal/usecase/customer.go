package usecase

import (
	"github.com/nadiannis/evento-api-fr/internal/domain"
	"github.com/nadiannis/evento-api-fr/internal/domain/request"
	"github.com/nadiannis/evento-api-fr/internal/domain/response"
	"github.com/nadiannis/evento-api-fr/internal/repository"
)

type CustomerUsecase struct {
	customerRepository repository.ICustomerRepository
	orderRepository    repository.IOrderRepository
}

func NewCustomerUsecase(customerRepository repository.ICustomerRepository, orderRepository repository.IOrderRepository) ICustomerUsecase {
	return &CustomerUsecase{
		customerRepository: customerRepository,
		orderRepository:    orderRepository,
	}
}

func (u *CustomerUsecase) GetAll() ([]*response.CustomerResponse, error) {
	customers, err := u.customerRepository.GetAll()
	if err != nil {
		return nil, err
	}

	customerResponses := make([]*response.CustomerResponse, 0)

	for _, customer := range customers {
		orders, err := u.orderRepository.GetByCustomerID(customer.ID)
		if err != nil {
			return nil, err
		}

		customerResponse := &response.CustomerResponse{
			ID:       customer.ID,
			Username: customer.Username,
			Balance:  customer.Balance,
			Orders:   orders,
		}
		customerResponses = append(customerResponses, customerResponse)
	}

	return customerResponses, nil
}

func (u *CustomerUsecase) Add(input *request.CustomerRequest) (*domain.Customer, error) {
	customer := &domain.Customer{
		Username: input.Username,
		Balance:  input.Balance,
	}

	err := u.customerRepository.Add(customer)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (u *CustomerUsecase) GetByID(customerID int64) (*response.CustomerResponse, error) {
	customer, err := u.customerRepository.GetByID(customerID)
	if err != nil {
		return nil, err
	}

	orders, err := u.orderRepository.GetByCustomerID(customer.ID)
	if err != nil {
		return nil, err
	}

	customerResponse := &response.CustomerResponse{
		ID:       customer.ID,
		Username: customer.Username,
		Balance:  customer.Balance,
		Orders:   orders,
	}

	return customerResponse, nil
}

func (u *CustomerUsecase) AddBalance(customerID int64, input *request.CustomerBalanceRequest) (*domain.Customer, error) {
	_, err := u.customerRepository.GetByID(customerID)
	if err != nil {
		return nil, err
	}

	customer, err := u.customerRepository.AddBalance(customerID, input.Balance)
	if err != nil {
		return nil, err
	}

	return customer, nil
}
