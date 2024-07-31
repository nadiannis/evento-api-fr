package main

import (
	"fmt"
	"time"

	"github.com/nadiannis/evento-api-fr/internal/domain"
	"github.com/nadiannis/evento-api-fr/internal/domain/request"
	"github.com/nadiannis/evento-api-fr/internal/usecase"
)

var ticketTypeInputs = []*request.TicketTypeRequest{
	{
		Name:  domain.TicketTypeVIP,
		Price: 5000,
	},
	{
		Name:  domain.TicketTypeCAT1,
		Price: 250,
	},
}

var eventInputs = []*request.EventRequest{
	{
		Name: "Event 1",
		Date: time.Now().AddDate(0, 0, 14),
	},
	{
		Name: "Event 2",
		Date: time.Now().AddDate(0, 1, 0),
	},
	{
		Name: "Event 3",
		Date: time.Now().AddDate(0, 1, 14),
	},
	{
		Name: "Event 4",
		Date: time.Now().AddDate(0, 2, 0),
	},
	{
		Name: "Event 5",
		Date: time.Now().AddDate(0, 2, 14),
	},
}

func prepopulateTicketTypes(usecase usecase.ITicketTypeUsecase) {
	ticketTypes, _ := usecase.GetAll()
	if len(ticketTypes) != 0 {
		return
	}

	for _, ticketTypeInput := range ticketTypeInputs {
		usecase.Add(ticketTypeInput)
	}
}

func prepopulateEventsAndTickets(eventUsecase usecase.IEventUsecase, ticketUsecase usecase.ITicketUsecase) {
	events, _ := eventUsecase.GetAll()
	if len(events) != 0 {
		return
	}

	for _, eventInput := range eventInputs {
		event, err := eventUsecase.Add(eventInput)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		vipTicket := &request.TicketRequest{
			EventID:  event.ID,
			Type:     domain.TicketTypeVIP,
			Quantity: 10,
		}
		cat1Ticket := &request.TicketRequest{
			EventID:  event.ID,
			Type:     domain.TicketTypeCAT1,
			Quantity: 100,
		}
		ticketUsecase.Add(vipTicket)
		ticketUsecase.Add(cat1Ticket)
	}
}
