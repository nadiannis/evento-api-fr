package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nadiannis/evento-api-fr/internal/domain/response"
	"github.com/nadiannis/evento-api-fr/internal/usecase"
	"github.com/nadiannis/evento-api-fr/internal/utils"
)

type EventHandler struct {
	usecase usecase.IEventUsecase
}

func NewEventHandler(usecase usecase.IEventUsecase) IEventHandler {
	return &EventHandler{
		usecase: usecase,
	}
}

func (h *EventHandler) GetAll(c *gin.Context) {
	events, err := h.usecase.GetAll()
	if err != nil {
		utils.ServerErrorResponse(c, err)
		return
	}

	res := response.SuccessResponse{
		Status:  response.Success,
		Message: "events retrieved successfully",
		Data:    events,
	}

	utils.WriteJSON(c, http.StatusOK, res)
}

func (h *EventHandler) GetByID(c *gin.Context) {
	id, err := utils.ReadIDParam(c)
	if err != nil {
		utils.BadRequestResponse(c, utils.ErrInvalidID)
		return
	}

	event, err := h.usecase.GetByID(id)
	if err != nil {
		switch {
		case errors.Is(err, utils.ErrEventNotFound):
			utils.NotFoundResponse(c, err)
		default:
			utils.ServerErrorResponse(c, err)
		}
		return
	}

	res := response.SuccessResponse{
		Status:  response.Success,
		Message: "event retrieved successfully",
		Data:    event,
	}

	utils.WriteJSON(c, http.StatusOK, res)
}
