package usecase

import (
	"time"

	"github.com/nadiannis/evento-api-fr/internal/domain"
	"github.com/nadiannis/evento-api-fr/internal/domain/request"
	"github.com/nadiannis/evento-api-fr/internal/repository"
)

type OrderUsecase struct {
	orderRepository      repository.IOrderRepository
	customerRepository   repository.ICustomerRepository
	ticketRepository     repository.ITicketRepository
	ticketTypeRepository repository.ITicketTypeRepository
}

func NewOrderUsecase(
	orderRepository repository.IOrderRepository,
	customerRepository repository.ICustomerRepository,
	ticketRepository repository.ITicketRepository,
	ticketTypeRepository repository.ITicketTypeRepository,
) IOrderUsecase {
	return &OrderUsecase{
		orderRepository:      orderRepository,
		customerRepository:   customerRepository,
		ticketRepository:     ticketRepository,
		ticketTypeRepository: ticketTypeRepository,
	}
}

func (u *OrderUsecase) GetAll() ([]*domain.Order, error) {
	return u.orderRepository.GetAll()
}

func (u *OrderUsecase) Add(input *request.OrderRequest) (*domain.Order, error) {
	customer, err := u.customerRepository.GetByID(input.CustomerID)
	if err != nil {
		return nil, err
	}

	ticket, err := u.ticketRepository.GetByID(input.TicketID)
	if err != nil {
		return nil, err
	}

	ticketType, err := u.ticketTypeRepository.GetByName(ticket.Type)
	if err != nil {
		return nil, err
	}

	err = u.ticketRepository.DeductQuantity(ticket.ID, input.Quantity)
	if err != nil {
		return nil, err
	}

	totalPrice := float64(input.Quantity) * ticketType.Price
	err = u.customerRepository.DeductBalance(customer.ID, totalPrice)
	if err != nil {
		u.ticketRepository.AddQuantity(ticket.ID, input.Quantity)

		return nil, err
	}

	order := &domain.Order{
		CustomerID: customer.ID,
		TicketID:   ticket.ID,
		Quantity:   input.Quantity,
		TotalPrice: totalPrice,
		CreatedAt:  time.Now(),
	}

	err = u.orderRepository.Add(order)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (u *OrderUsecase) DeleteAll() {
}
