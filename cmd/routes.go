package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nadiannis/evento-api-fr/internal/utils"
)

func (app *application) routes() *gin.Engine {
	r := gin.Default()

	r.Use(requestLogger())

	r.GET("/api", func(c *gin.Context) {
		message := fmt.Sprintf("API is running on port %d", app.config.port)
		c.String(http.StatusOK, message)
		utils.SetLogMessage(c, message)
	})

	r.GET("/api/customers", app.handlers.Customers.GetAll)
	r.POST("/api/customers", app.handlers.Customers.Add)
	r.GET("/api/customers/:id", app.handlers.Customers.GetByID)
	r.PATCH("/api/customers/:id/balances", app.handlers.Customers.UpdateBalance)

	r.GET("/api/events", app.handlers.Events.GetAll)
	r.GET("/api/events/:id", app.handlers.Events.GetByID)

	r.GET("/api/tickets", app.handlers.Tickets.GetAll)
	r.GET("/api/tickets/:id", app.handlers.Tickets.GetByID)
	r.PATCH("/api/tickets/:id/quantities", app.handlers.Tickets.UpdateQuantity) // Intended solely for concurrency testing purpose

	r.GET("/api/orders", app.handlers.Orders.GetAll)
	r.POST("/api/orders", app.handlers.Orders.Add)
	r.DELETE("/api/orders", app.handlers.Orders.DeleteAll) // Intended solely for concurrency testing purpose

	return r
}
