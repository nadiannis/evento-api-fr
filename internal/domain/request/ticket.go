package request

import "github.com/nadiannis/evento-api-fr/internal/domain"

type TicketRequest struct {
	EventID  int64                 `json:"event_id"`
	Type     domain.TicketTypeName `json:"type"`
	Quantity int                   `json:"quantity"`
}

type TicketQuantityRequest struct {
	Action   UpdateNumberAction `json:"action"`
	Quantity int                `json:"quantity"`
}
