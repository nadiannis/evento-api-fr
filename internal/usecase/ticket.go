package usecase

import (
	"github.com/nadiannis/evento-api-fr/internal/domain"
	"github.com/nadiannis/evento-api-fr/internal/domain/request"
	"github.com/nadiannis/evento-api-fr/internal/repository"
)

type TicketUsecase struct {
	ticketRepository repository.ITicketRepository
	eventRepository  repository.IEventRepository
}

func NewTicketUsecase(ticketRepository repository.ITicketRepository, eventRepository repository.IEventRepository) ITicketUsecase {
	return &TicketUsecase{
		ticketRepository: ticketRepository,
		eventRepository:  eventRepository,
	}
}

func (u *TicketUsecase) GetAll() ([]*domain.Ticket, error) {
	return u.ticketRepository.GetAll()
}

func (u *TicketUsecase) Add(input *request.TicketRequest) (*domain.Ticket, error) {
	ticket := &domain.Ticket{
		EventID:  input.EventID,
		Type:     input.Type,
		Quantity: input.Quantity,
	}

	err := u.ticketRepository.Add(ticket)
	if err != nil {
		return nil, err
	}

	_, err = u.eventRepository.AddTicket(ticket.EventID, ticket)
	if err != nil {
		return nil, err
	}

	return ticket, nil
}

func (u *TicketUsecase) GetByID(ticketID int64) (*domain.Ticket, error) {
	ticket, err := u.ticketRepository.GetByID(ticketID)
	if err != nil {
		return nil, err
	}

	return ticket, nil
}

func (u *TicketUsecase) AddQuantity(ticketID int64, input *request.TicketQuantityRequest) (*domain.Ticket, error) {
	_, err := u.ticketRepository.GetByID(ticketID)
	if err != nil {
		return nil, err
	}

	ticket, err := u.ticketRepository.AddQuantity(ticketID, input.Quantity)
	if err != nil {
		return nil, err
	}

	return ticket, nil
}
