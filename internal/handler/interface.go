package handler

import (
	"github.com/gin-gonic/gin"
)

type CustomerReader interface {
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
}

type CustomerWriter interface {
	Add(c *gin.Context)
	UpdateBalance(c *gin.Context)
}

type ICustomerHandler interface {
	CustomerReader
	CustomerWriter
}

type EventReader interface {
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
}

type IEventHandler interface {
	EventReader
}

type TicketReader interface {
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
}

type TicketWriter interface {
	UpdateQuantity(c *gin.Context)
}

type ITicketHandler interface {
	TicketReader
	TicketWriter
}

type OrderReader interface {
	GetAll(c *gin.Context)
}

type OrderWriter interface {
	Add(c *gin.Context)
	DeleteAll(c *gin.Context)
}

type IOrderHandler interface {
	OrderReader
	OrderWriter
}
