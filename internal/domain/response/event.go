package response

import (
	"time"

	"github.com/nadiannis/evento-api-fr/internal/domain"
)

type EventResponse struct {
	ID      int64            `json:"id"`
	Name    string           `json:"name"`
	Date    time.Time        `json:"date"`
	Tickets []*domain.Ticket `json:"tickets"`
}
