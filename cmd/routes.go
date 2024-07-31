package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) routes() http.Handler {
	r := gin.Default()

	r.GET("/api", func(c *gin.Context) {
		c.String(http.StatusOK, "API is running on port %d\n", app.config.port)
	})

	r.GET("/api/customers", app.handlers.Customers.GetAll)
	r.POST("/api/customers", app.handlers.Customers.Add)
	r.GET("/api/customers/{id}", app.handlers.Customers.GetByID)
	r.PATCH("/api/customers/{id}/balances", app.handlers.Customers.AddBalance)

	r.GET("/api/events", app.handlers.Events.GetAll)
	r.GET("/api/events/{id}", app.handlers.Events.GetByID)

	r.GET("/api/tickets", app.handlers.Tickets.GetAll)
	r.GET("/api/tickets/{id}", app.handlers.Tickets.GetByID)
	r.PATCH("/api/tickets/{id}/quantities", app.handlers.Tickets.AddQuantity) // Intended solely for concurrency testing purpose

	r.GET("/api/orders", app.handlers.Orders.GetAll)
	r.POST("/api/orders", app.handlers.Orders.Add)
	r.DELETE("/api/orders", app.handlers.Orders.DeleteAll) // Intended solely for concurrency testing purpose

	return requestLogger(r)
}
