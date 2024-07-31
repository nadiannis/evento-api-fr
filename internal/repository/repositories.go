package repository

import "database/sql"

type Repositories struct {
	Customers   ICustomerRepository
	Events      IEventRepository
	TicketTypes ITicketTypeRepository
	Tickets     ITicketRepository
	Orders      IOrderRepository
}

func NewRepositories(db *sql.DB) Repositories {
	return Repositories{
		Customers:   NewCustomerRepository(db),
		Events:      NewEventRepository(db),
		TicketTypes: NewTicketTypeRepository(db),
		Tickets:     NewTicketRepository(db),
		Orders:      NewOrderRepository(db),
	}
}
