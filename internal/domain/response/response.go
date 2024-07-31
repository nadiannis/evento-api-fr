package response

type ResponseStatus string

var (
	Success ResponseStatus = "success"
	Error   ResponseStatus = "error"
)

type Response interface {
}

type SuccessResponse struct {
	Status  ResponseStatus `json:"status"`
	Message string         `json:"message"`
	Data    any            `json:"data"`
}

type ErrorResponse struct {
	Status  ResponseStatus `json:"status"`
	Message string         `json:"message"`
	Detail  any            `json:"detail"`
}
