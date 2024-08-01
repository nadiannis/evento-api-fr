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

type TicketHandler struct {
	usecase usecase.ITicketUsecase
}

func NewTicketHandler(usecase usecase.ITicketUsecase) ITicketHandler {
	return &TicketHandler{
		usecase: usecase,
	}
}

func (h *TicketHandler) GetAll(c *gin.Context) {
	tickets, err := h.usecase.GetAll()
	if err != nil {
		utils.ServerErrorResponse(c, err)
		return
	}

	res := response.SuccessResponse{
		Status:  response.Success,
		Message: "tickets retrieved successfully",
		Data:    tickets,
	}

	utils.WriteJSON(c, http.StatusOK, res)
}

func (h *TicketHandler) GetByID(c *gin.Context) {
	id, err := utils.ReadIDParam(c)
	if err != nil {
		utils.BadRequestResponse(c, utils.ErrInvalidID)
		return
	}

	ticket, err := h.usecase.GetByID(id)
	if err != nil {
		switch {
		case errors.Is(err, utils.ErrTicketNotFound):
			utils.NotFoundResponse(c, err)
		default:
			utils.ServerErrorResponse(c, err)
		}
		return
	}

	res := response.SuccessResponse{
		Status:  response.Success,
		Message: "ticket retrieved successfully",
		Data:    ticket,
	}

	utils.WriteJSON(c, http.StatusOK, res)
}

func (h *TicketHandler) UpdateQuantity(c *gin.Context) {
	id, err := utils.ReadIDParam(c)
	if err != nil {
		utils.BadRequestResponse(c, utils.ErrInvalidID)
		return
	}

	var input request.TicketQuantityRequest

	err = utils.ReadJSON(c, &input)
	if err != nil {
		utils.BadRequestResponse(c, err)
		return
	}

	v := utils.NewValidator()

	v.Check(input.Action != "", "action", "action is required")
	v.Check(utils.PermittedValue(input.Action, request.ActionAdd, request.ActionDeduct), "action", "action should be 'add' or 'deduct'")
	v.Check(input.Quantity != 0, "quantity", "quantity is required")
	v.Check(input.Quantity > 0, "quantity", "quantity should not be a negative number")

	if !v.Valid() {
		utils.FailedValidationResponse(c, v.Errors)
		return
	}

	ticket, err := h.usecase.UpdateQuantity(id, &input)
	if err != nil {
		switch {
		case errors.Is(err, utils.ErrTicketNotFound):
			utils.NotFoundResponse(c, err)
		case errors.Is(err, utils.ErrInsufficientTicketQuantity):
			utils.BadRequestResponse(c, err)
		default:
			utils.ServerErrorResponse(c, err)
		}
		return
	}

	res := response.SuccessResponse{
		Status:  response.Success,
		Message: "ticket quantity updated successfully",
		Data:    ticket,
	}

	utils.WriteJSON(c, http.StatusOK, res)
}
