package handler

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nadiannis/evento-api-fr/internal/domain/request"
	"github.com/nadiannis/evento-api-fr/internal/domain/response"
	"github.com/nadiannis/evento-api-fr/internal/usecase"
	"github.com/nadiannis/evento-api-fr/internal/utils"
)

type CustomerHandler struct {
	usecase usecase.ICustomerUsecase
}

func NewCustomerHandler(usecase usecase.ICustomerUsecase) ICustomerHandler {
	return &CustomerHandler{
		usecase: usecase,
	}
}

func (h *CustomerHandler) GetAll(c *gin.Context) {
	customers, err := h.usecase.GetAll()
	if err != nil {
		utils.ServerErrorResponse(c, err)
		return
	}

	time.Sleep(2 * time.Second) // Simulate real processing time

	res := response.SuccessResponse{
		Status:  response.Success,
		Message: "customers retrieved successfully",
		Data:    customers,
	}

	utils.WriteJSON(c, http.StatusOK, res)
}

func (h *CustomerHandler) Add(c *gin.Context) {
	var input request.CustomerRequest

	err := utils.ReadJSON(c, &input)
	if err != nil {
		utils.BadRequestResponse(c, err)
		return
	}

	v := utils.NewValidator()

	v.Check(input.Username != "", "username", "username is required")
	v.Check(input.Balance >= 0, "balance", "balance should not be a negative number")

	if !v.Valid() {
		utils.FailedValidationResponse(c, v.Errors)
		return
	}

	customer, err := h.usecase.Add(&input)
	if err != nil {
		switch {
		case errors.Is(err, utils.ErrCustomerAlreadyExists):
			utils.BadRequestResponse(c, err)
		default:
			utils.ServerErrorResponse(c, err)
		}
		return
	}

	res := response.SuccessResponse{
		Status:  response.Success,
		Message: "customer added successfully",
		Data:    customer,
	}

	utils.WriteJSON(c, http.StatusCreated, res)
}

func (h *CustomerHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	customerID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		utils.BadRequestResponse(c, utils.ErrInvalidID)
		return
	}

	customer, err := h.usecase.GetByID(customerID)
	if err != nil {
		switch {
		case errors.Is(err, utils.ErrCustomerNotFound):
			utils.NotFoundResponse(c, err)
		default:
			utils.ServerErrorResponse(c, err)
		}
		return
	}

	res := response.SuccessResponse{
		Status:  response.Success,
		Message: "customer retrieved successfully",
		Data:    customer,
	}

	utils.WriteJSON(c, http.StatusOK, res)
}

func (h *CustomerHandler) AddBalance(c *gin.Context) {
	id := c.Param("id")
	customerID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		utils.BadRequestResponse(c, utils.ErrInvalidID)
		return
	}

	var input request.CustomerBalanceRequest

	err = utils.ReadJSON(c, &input)
	if err != nil {
		utils.BadRequestResponse(c, err)
		return
	}

	v := utils.NewValidator()

	v.Check(input.Balance != 0, "balance", "balance is required")
	v.Check(input.Balance > 0, "balance", "balance should not be a negative number")

	if !v.Valid() {
		utils.FailedValidationResponse(c, v.Errors)
		return
	}

	customer, err := h.usecase.AddBalance(customerID, &input)
	if err != nil {
		switch {
		case errors.Is(err, utils.ErrCustomerNotFound):
			utils.NotFoundResponse(c, err)
		default:
			utils.ServerErrorResponse(c, err)
		}
		return
	}

	res := response.SuccessResponse{
		Status:  response.Success,
		Message: "customer balance added successfully",
		Data:    customer,
	}

	utils.WriteJSON(c, http.StatusOK, res)
}
