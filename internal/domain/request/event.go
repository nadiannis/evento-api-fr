package request

import "time"

type EventRequest struct {
	Name string    `json:"name"`
	Date time.Time `json:"date"`
}
