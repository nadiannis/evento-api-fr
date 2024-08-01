package request

type CustomerRequest struct {
	Username string  `json:"username"`
	Balance  float64 `json:"balance"`
}

type CustomerBalanceRequest struct {
	Action  UpdateNumberAction `json:"action"`
	Balance float64            `json:"balance"`
}
