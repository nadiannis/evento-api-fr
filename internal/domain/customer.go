package domain

type Customer struct {
	ID       int64   `json:"id"`
	Username string  `json:"username"`
	Balance  float64 `json:"balance"`
}
