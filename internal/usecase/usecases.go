package usecase

import "github.com/nadiannis/evento-api-fr/internal/repository"

type Usecases struct {
	Customers   ICustomerUsecase
	Events      IEventUsecase
	TicketTypes ITicketTypeUsecase
	Tickets     ITicketUsecase
	Orders      IOrderUsecase
}

func NewUsecases(repositories repository.Repositories) Usecases {
	return Usecases{
		Customers:   NewCustomerUsecase(repositories.Customers, repositories.Orders),
		Events:      NewEventUsecase(repositories.Events, repositories.Tickets),
		TicketTypes: NewTicketTypeUsecase(repositories.TicketTypes),
		Tickets:     NewTicketUsecase(repositories.Tickets, repositories.Events),
		Orders: NewOrderUsecase(
			repositories.Orders,
			repositories.Customers,
			repositories.Tickets,
			repositories.TicketTypes,
		),
	}
}
