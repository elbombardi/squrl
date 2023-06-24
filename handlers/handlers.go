package handlers

import (
	"github.com/ashley-nygaard/platform/orders_service/api/operations"
	"github.com/ashley-nygaard/platform/orders_service/core"
)

type Handlers struct {
	OrderSubmitter core.OrderSubmitter
	OrderRetriever core.OrderRetriever
	OrderUpdater   core.OrderUpdater
}

func InstallHandlers(api *operations.OrdersAPI, hdlrs *Handlers) {
	api.GetOrderHandler = operations.GetOrderHandlerFunc(hdlrs.GetOrderHandler)
	api.ListOrdersHandler = operations.ListOrdersHandlerFunc(hdlrs.ListOrdersHandler)
	api.SubmitOrderHandler = operations.SubmitOrderHandlerFunc(hdlrs.SubmitOrderHandler)
	api.CancelOrderHandler = operations.CancelOrderHandlerFunc(hdlrs.CancelOrder)
	api.UpdateOrderHandler = operations.UpdateOrderHandlerFunc(hdlrs.UpdateOrder)
}
