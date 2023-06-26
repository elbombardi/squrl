package handlers

import (
	"context"
	"database/sql"
	"errors"

	"github.com/elbombardi/squrl/api_service/api/operations"
	db "github.com/elbombardi/squrl/db/sqlc"
	"github.com/elbombardi/squrl/util"
	"github.com/go-openapi/runtime/middleware"
)

func (h *Handlers) UpdateCustomerHandler(params operations.PutCustomerParams) middleware.Responder {
	//Validate params
	err := validateUpdateCustomerParams(params)
	if err != nil {
		return operations.NewPutCustomerBadRequest().WithPayload(&operations.PutCustomerBadRequestBody{
			Error: err.Error()})
	}

	//Check if the Admin API key is valid
	if params.XAPIKEY != *util.ConfigAdminAPIKey() {
		return operations.NewPostCustomerUnauthorized().WithPayload(&operations.PostCustomerUnauthorizedBody{
			Error: "invalid x-api-key header"})
	}

	//Check if the customer exists
	_, err = h.CustomersRepository.GetCustomerByUsername(context.Background(), *params.Body.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return operations.NewPutCustomerNotFound().WithPayload(&operations.PutCustomerNotFoundBody{
				Error: "customer not found for this username: " + *params.Body.Username})
		}
		return internalErrorInUpdateCustomer(err)
	}

	//Update customer
	err = h.CustomersRepository.UpdateCustomerStatusByUsername(context.Background(), db.UpdateCustomerStatusByUsernameParams{
		Username: *params.Body.Username,
		Status:   encodeStatus(*params.Body.Status),
	})
	if err != nil {
		return internalErrorInUpdateCustomer(err)
	}
	return operations.NewPutCustomerOK().WithPayload(&operations.PutCustomerOKBody{
		Status: *params.Body.Status,
	})
}

func validateUpdateCustomerParams(params operations.PutCustomerParams) error {
	if params.Body.Username == nil {
		return errors.New("missing username")
	}
	if params.Body.Status == nil {
		return errors.New("missing status")
	}
	if *params.Body.Status != "active" && *params.Body.Status != "inactive" {
		return errors.New("invalid status, should be one of the two values: 'active', 'inactive'")
	}
	if params.XAPIKEY == "" {
		return errors.New("missing x-api-key header")
	}
	return nil
}

func internalErrorInUpdateCustomer(err error) middleware.Responder {
	return operations.NewPutCustomerInternalServerError().WithPayload(&operations.PutCustomerInternalServerErrorBody{
		Error: err.Error()})
}
