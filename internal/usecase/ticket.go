package usecase

import (
	"github.com/nadiannis/evento-api-fr/internal/domain"
	"github.com/nadiannis/evento-api-fr/internal/domain/request"
	"github.com/nadiannis/evento-api-fr/internal/repository"
	"github.com/nadiannis/evento-api-fr/internal/utils"
)

type TicketUsecase struct {
	ticketRepository     repository.ITicketRepository
	ticketTypeRepository repository.ITicketTypeRepository
	eventRepository      repository.IEventRepository
}

func NewTicketUsecase(
	ticketRepository repository.ITicketRepository,
	ticketTypeRepository repository.ITicketTypeRepository,
	eventRepository repository.IEventRepository,
) ITicketUsecase {
	return &TicketUsecase{
		ticketRepository:     ticketRepository,
		ticketTypeRepository: ticketTypeRepository,
		eventRepository:      eventRepository,
	}
}

func (u *TicketUsecase) GetAll() ([]*domain.TicketDetail, error) {
	return u.ticketRepository.GetAll()
}

func (u *TicketUsecase) Add(input *request.TicketRequest) (*domain.Ticket, error) {
	ticketType, err := u.ticketTypeRepository.GetByName(input.Type)
	if err != nil {
		return nil, utils.ErrTicketTypeNotFound
	}

	ticket := &domain.Ticket{
		EventID:      input.EventID,
		TicketTypeID: ticketType.ID,
		Quantity:     input.Quantity,
	}

	err = u.ticketRepository.Add(ticket)
	if err != nil {
		return nil, err
	}

	_, err = u.eventRepository.AddTicket(ticket.EventID, ticket)
	if err != nil {
		return nil, err
	}

	return ticket, nil
}

func (u *TicketUsecase) GetByID(ticketID int64) (*domain.TicketDetail, error) {
	ticketDetail, err := u.ticketRepository.GetByID(ticketID)
	if err != nil {
		return nil, err
	}

	return ticketDetail, nil
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
