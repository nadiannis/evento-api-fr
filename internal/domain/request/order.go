package request

type OrderRequest struct {
	CustomerID int64 `json:"customer_id"`
	TicketID   int64 `json:"ticket_id"`
	Quantity   int   `json:"quantity"`
}
