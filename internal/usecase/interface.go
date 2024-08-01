package usecase

import (
	"github.com/nadiannis/evento-api-fr/internal/domain"
	"github.com/nadiannis/evento-api-fr/internal/domain/request"
	"github.com/nadiannis/evento-api-fr/internal/domain/response"
)

type CustomerReader interface {
	GetAll() ([]*response.CustomerResponse, error)
	GetByID(customerID int64) (*response.CustomerResponse, error)
}

type CustomerWriter interface {
	Add(input *request.CustomerRequest) (*domain.Customer, error)
	UpdateBalance(customerID int64, input *request.CustomerBalanceRequest) (*domain.Customer, error)
}

type ICustomerUsecase interface {
	CustomerReader
	CustomerWriter
}

type EventReader interface {
	GetAll() ([]*response.EventResponse, error)
	GetByID(eventID int64) (*response.EventResponse, error)
}

type EventWriter interface {
	Add(input *request.EventRequest) (*domain.Event, error)
}

type IEventUsecase interface {
	EventReader
	EventWriter
}

type TicketTypeReader interface {
	GetAll() ([]*domain.TicketType, error)
}

type TicketTypeWriter interface {
	Add(input *request.TicketTypeRequest) (*domain.TicketType, error)
}

type ITicketTypeUsecase interface {
	TicketTypeReader
	TicketTypeWriter
}

type TicketReader interface {
	GetAll() ([]*domain.TicketDetail, error)
	GetByID(ticketID int64) (*domain.TicketDetail, error)
}

type TicketWriter interface {
	Add(input *request.TicketRequest) (*domain.Ticket, error)
	AddQuantity(ticketID int64, input *request.TicketQuantityRequest) (*domain.Ticket, error)
}

type ITicketUsecase interface {
	TicketReader
	TicketWriter
}

type OrderReader interface {
	GetAll() ([]*domain.Order, error)
}

type OrderWriter interface {
	Add(input *request.OrderRequest) (*domain.Order, error)
	DeleteAll()
}

type IOrderUsecase interface {
	OrderReader
	OrderWriter
}
