package handler

import "github.com/nadiannis/evento-api-fr/internal/usecase"

type Handlers struct {
	Customers ICustomerHandler
	Events    IEventHandler
	Tickets   ITicketHandler
	Orders    IOrderHandler
}

func NewHandlers(usecases usecase.Usecases) Handlers {
	return Handlers{
		Customers: NewCustomerHandler(usecases.Customers),
		Events:    NewEventHandler(usecases.Events),
		Tickets:   NewTicketHandler(usecases.Tickets),
		Orders:    NewOrderHandler(usecases.Orders),
	}
}
