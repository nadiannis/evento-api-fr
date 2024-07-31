package response

import "github.com/nadiannis/evento-api-fr/internal/domain"

type CustomerResponse struct {
	ID       int64           `json:"id"`
	Username string          `json:"username"`
	Balance  float64         `json:"balance"`
	Orders   []*domain.Order `json:"orders"`
}
