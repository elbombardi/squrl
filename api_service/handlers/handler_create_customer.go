package handlers

import (
	"context"

	"github.com/elbombardi/squrl/api_service/api/operations"
	db "github.com/elbombardi/squrl/db/sqlc"
	"github.com/elbombardi/squrl/util"
	"github.com/go-openapi/runtime/middleware"
)

func (h *Handlers) CreateCustomerHandler(params operations.PostCustomerParams) middleware.Responder {
	prefix := util.GenerateRandomString(3)
	apiKey := util.GenerateRandomString(32)
	err := h.CustomersRepository.InsertNewCustomer(context.Background(), db.InsertNewCustomerParams{
		Prefix:   prefix,
		ApiKey:   apiKey,
		Username: *params.Customer.Username,
		Email:    *params.Customer.Email,
	})
	if err != nil {
		return operations.NewPostCustomerInternalServerError().WithPayload(&operations.PostCustomerInternalServerErrorBody{
			Error: err.Error()})
	}
	return operations.NewPostCustomerOK().WithPayload(&operations.PostCustomerOKBody{
		APIKey: apiKey,
		Prefix: prefix,
	})
}
