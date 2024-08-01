package domain

type Ticket struct {
	ID           int64 `json:"id"`
	EventID      int64 `json:"event_id"`
	TicketTypeID int64 `json:"ticket_type_id"`
	Quantity     int   `json:"quantity"`
}

type TicketDetail struct {
	ID       int64      `json:"id"`
	EventID  int64      `json:"event_id"`
	Quantity int        `json:"quantity"`
	Type     TicketType `json:"type"`
}
