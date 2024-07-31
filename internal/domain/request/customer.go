package request

type CustomerRequest struct {
	Username string  `json:"username"`
	Balance  float64 `json:"balance"`
}

type CustomerBalanceRequest struct {
	Balance float64 `json:"balance"`
}
