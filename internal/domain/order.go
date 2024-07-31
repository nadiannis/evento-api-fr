package domain

import "time"

type Order struct {
	ID         int64     `json:"id"`
	CustomerID int64     `json:"customer_id"`
	TicketID   int64     `json:"ticket_id"`
	Quantity   int       `json:"quantity"`
	TotalPrice float64   `json:"total_price"`
	CreatedAt  time.Time `json:"created_at"`
}
