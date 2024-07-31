package usecase

import (
	"github.com/nadiannis/evento-api-fr/internal/domain"
	"github.com/nadiannis/evento-api-fr/internal/domain/request"
	"github.com/nadiannis/evento-api-fr/internal/domain/response"
	"github.com/nadiannis/evento-api-fr/internal/repository"
)

type EventUsecase struct {
	eventRepository  repository.IEventRepository
	ticketRepository repository.ITicketRepository
}

func NewEventUsecase(eventRepository repository.IEventRepository, ticketRepository repository.ITicketRepository) IEventUsecase {
	return &EventUsecase{
		eventRepository:  eventRepository,
		ticketRepository: ticketRepository,
	}
}

func (u *EventUsecase) GetAll() ([]*response.EventResponse, error) {
	events, err := u.eventRepository.GetAll()
	if err != nil {
		return nil, err
	}

	eventResponses := make([]*response.EventResponse, 0)

	for _, event := range events {
		tickets, err := u.ticketRepository.GetByEventID(event.ID)
		if err != nil {
			return nil, err
		}

		eventResponse := &response.EventResponse{
			ID:      event.ID,
			Name:    event.Name,
			Date:    event.Date,
			Tickets: tickets,
		}
		eventResponses = append(eventResponses, eventResponse)
	}

	return eventResponses, nil
}

func (u *EventUsecase) Add(input *request.EventRequest) (*domain.Event, error) {
	event := &domain.Event{
		Name: input.Name,
		Date: input.Date,
	}

	err := u.eventRepository.Add(event)
	if err != nil {
		return nil, err
	}

	return event, nil
}

func (u *EventUsecase) GetByID(eventID int64) (*response.EventResponse, error) {
	event, err := u.eventRepository.GetByID(eventID)
	if err != nil {
		return nil, err
	}

	tickets, err := u.ticketRepository.GetByEventID(event.ID)
	if err != nil {
		return nil, err
	}

	eventResponse := &response.EventResponse{
		ID:      event.ID,
		Name:    event.Name,
		Date:    event.Date,
		Tickets: tickets,
	}

	return eventResponse, nil
}
