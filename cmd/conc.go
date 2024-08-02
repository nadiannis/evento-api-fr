package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/nadiannis/evento-api-fr/internal/domain/request"
	"github.com/nadiannis/evento-api-fr/internal/usecase"
)

func ConcurrentOrderCreation(orderUsecase usecase.IOrderUsecase, ticketUsecase usecase.ITicketUsecase) {
	const numOrders = 300
	var wg sync.WaitGroup
	errors := make(chan error, numOrders)

	var customerID int64 = 1
	var ticketID int64 = 2

	ticket, err := ticketUsecase.GetByID(ticketID)
	if err != nil {
		fmt.Println("error getting tickets:", err.Error())
		return
	}

	if ticket.Quantity == 0 {
		ticketQuantityInput := &request.TicketQuantityRequest{
			Action:   "add",
			Quantity: 100,
		}
		ticketUsecase.UpdateQuantity(ticketID, ticketQuantityInput)
	}

	for i := 0; i < numOrders; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)

			// Call your API to create an order
			orderInput := &request.OrderRequest{
				CustomerID: customerID,
				TicketID:   ticketID,
				Quantity:   1,
			}
			_, err := orderUsecase.Add(orderInput)
			if err != nil {
				errors <- err
			}
		}()
	}

	wg.Wait()
	close(errors)

	for err := range errors {
		fmt.Printf("error creating order: %v\n", err)
	}
}
