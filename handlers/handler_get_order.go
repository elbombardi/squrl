package handlers

import (
	"log"
	"net/http"

	"github.com/ashley-nygaard/platform/orders_service/api/operations"
	"github.com/ashley-nygaard/platform/orders_service/errors"
	"github.com/ashley-nygaard/platform/orders_service/models"
	"github.com/go-openapi/runtime/middleware"
)

func (handlers *Handlers) GetOrderHandler(params operations.GetOrderParams) middleware.Responder {
	log.Printf("GET /order/%v\n", params.ID)
	order, err := handlers.OrderRetriever.GetOrder(params.ID)
	if err != nil {
		noSuchOrder, ok := err.(*errors.NoSuchOrder)
		if ok {
			msg := noSuchOrder.Error()
			return operations.NewGetOrderNotFound().WithPayload(&models.OrdersResponse{
				Status:  http.StatusNotFound,
				Message: msg,
			})
		}

		invalidInput, ok := err.(*errors.InvalidInput)
		if ok {
			msg := invalidInput.Error()
			return operations.NewGetOrderUnprocessableEntity().WithPayload(&models.OrdersResponse{
				Status:  http.StatusUnprocessableEntity,
				Message: msg,
			})
		}
		msg := errors.InternalErrorMessage
		log.Println("ERROR", err)
		return operations.NewGetOrderDefault(http.StatusInternalServerError).WithPayload(&models.OrdersResponse{
			Status:  http.StatusInternalServerError,
			Message: msg,
		})
	}
	return operations.NewGetOrderOK().WithPayload(&models.OrdersResponse{
		Status: http.StatusOK,
		Order:  order,
	})
}
