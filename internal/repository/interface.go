package repository

import "github.com/nadiannis/evento-api-fr/internal/domain"

type CustomerReader interface {
	GetAll() ([]*domain.Customer, error)
	GetByID(customerID int64) (*domain.Customer, error)
}

type CustomerWriter interface {
	Add(customer *domain.Customer) error
	AddBalance(customerID int64, amount float64) (*domain.Customer, error)
	DeductBalance(customerID int64, amount float64) (*domain.Customer, error)
}

type ICustomerRepository interface {
	CustomerReader
	CustomerWriter
}

type EventReader interface {
	GetAll() ([]*domain.Event, error)
	GetByID(eventID int64) (*domain.Event, error)
}

type EventWriter interface {
	Add(event *domain.Event) error
}

type IEventRepository interface {
	EventReader
	EventWriter
}

type TicketTypeReader interface {
	GetAll() ([]*domain.TicketType, error)
	GetByName(ticketTypeName domain.TicketTypeName) (*domain.TicketType, error)
}

type TicketTypeWriter interface {
	Add(ticketType *domain.TicketType) error
}

type ITicketTypeRepository interface {
	TicketTypeReader
	TicketTypeWriter
}

type TicketReader interface {
	GetAll() ([]*domain.TicketDetail, error)
	GetByID(ticketID int64) (*domain.TicketDetail, error)
	GetByEventID(eventID int64) ([]*domain.TicketDetail, error)
}

type TicketWriter interface {
	Add(ticket *domain.Ticket) error
	AddQuantity(ticketID int64, quantity int) (*domain.Ticket, error)
	DeductQuantity(ticketID int64, quantity int) error
}

type ITicketRepository interface {
	TicketReader
	TicketWriter
}

type OrderReader interface {
	GetAll() ([]*domain.Order, error)
	GetByCustomerID(customerID int64) ([]*domain.Order, error)
}

type OrderWriter interface {
	Add(order *domain.Order) error
	DeleteByID(orderID int64) error
	DeleteAll()
}

type IOrderRepository interface {
	OrderReader
	OrderWriter
}
