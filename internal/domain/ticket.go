package domain

type Ticket struct {
	ID       int64          `json:"id"`
	EventID  int64          `json:"event_id"`
	Type     TicketTypeName `json:"type"`
	Quantity int            `json:"quantity"`
}
