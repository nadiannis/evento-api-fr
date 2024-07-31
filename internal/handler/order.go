package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nadiannis/evento-api-fr/internal/domain/request"
	"github.com/nadiannis/evento-api-fr/internal/domain/response"
	"github.com/nadiannis/evento-api-fr/internal/usecase"
	"github.com/nadiannis/evento-api-fr/internal/utils"
)

type OrderHandler struct {
	usecase usecase.IOrderUsecase
}

func NewOrderHandler(usecase usecase.IOrderUsecase) IOrderHandler {
	return &OrderHandler{
		usecase: usecase,
	}
}

func (h *OrderHandler) GetAll(c *gin.Context) {
	orders, err := h.usecase.GetAll()
	if err != nil {
		utils.ServerErrorResponse(c, err)
		return
	}

	res := response.SuccessResponse{
		Status:  response.Success,
		Message: "orders retrieved successfully",
		Data:    orders,
	}

	utils.WriteJSON(c, http.StatusOK, res)
}

func (h *OrderHandler) Add(c *gin.Context) {
	var input request.OrderRequest

	err := utils.ReadJSON(c, &input)
	if err != nil {
		utils.BadRequestResponse(c, err)
		return
	}

	v := utils.NewValidator()

	v.Check(input.CustomerID != 0, "customer_id", "customer_id is required")
	v.Check(input.TicketID != 0, "ticket_id", "ticket_id is required")
	v.Check(input.Quantity != 0, "quantity", "quantity is required")
	v.Check(input.Quantity > 0, "quantity", "quantity should not be a negative number")

	if !v.Valid() {
		utils.FailedValidationResponse(c, v.Errors)
		return
	}

	order, err := h.usecase.Add(&input)
	if err != nil {
		switch {
		case errors.Is(err, utils.ErrCustomerNotFound) || errors.Is(err, utils.ErrTicketNotFound) || errors.Is(err, utils.ErrTicketTypeNotFound):
			utils.NotFoundResponse(c, err)
		case errors.Is(err, utils.ErrInsufficientTicketQuantity) || errors.Is(err, utils.ErrInsufficientBalance):
			utils.BadRequestResponse(c, err)
		default:
			utils.ServerErrorResponse(c, err)
		}
		return
	}

	res := response.SuccessResponse{
		Status:  response.Success,
		Message: "order added successfully",
		Data:    order,
	}

	utils.WriteJSON(c, http.StatusCreated, res)
}

func (h *OrderHandler) DeleteAll(c *gin.Context) {
	h.usecase.DeleteAll()

	res := response.SuccessResponse{
		Status:  response.Success,
		Message: "orders deleted successfully",
		Data:    nil,
	}

	utils.WriteJSON(c, http.StatusOK, res)
}
