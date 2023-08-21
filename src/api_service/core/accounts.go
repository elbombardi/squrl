package core

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"

	"github.com/elbombardi/squrl/src/api_service/util"
	"github.com/elbombardi/squrl/src/db"
)

type AccountsService struct {
	db.AccountRepository
	*util.Config
	*slog.Logger
}

func (s *AccountsService) Create(params *CreateAccountParams, user *User) (*CreateAccountResponse, error) {
	// Check if the user is authenticated
	if user == nil {
		return nil, CoreError{
			Code:    ErrUnauthorized,
			Message: "Unauthorized access",
		}
	}

	//This service is only accessible by the admin user
	if !user.IsAdmin {
		s.Error("Unauthorized attempt to access CreateAccount by a non admin user", "User", user)
		return nil, CoreError{
			Code:    ErrUnauthorized,
			Message: "Unauthorized access",
		}
	}

	//Check if the request is valid
	err := validateCreateAccountParams(params)
	if err != nil {
		s.Error("Invalid parameters in create account request", "Details", err)
		return nil, CoreError{
			Code:    ErrBadParams,
			Message: err.Error(),
		}
	}

	//Check if the username is unique
	exists, err := s.CheckUsernameExists(context.Background(), params.Username)
	if err != nil {
		s.Error("Unexpected error while checking if username exists", "Details", err)
		return nil, err
	}
	if exists {
		s.Error("Username already exists", "Username", params.Username)
		return nil, CoreError{
			Code:    ErrBadParams,
			Message: "Username already exists",
		}
	}

	// Generate a unique prefix for the account
	var prefix string
	tryAgain := true
	for tryAgain {
		tryAgain, prefix, err = s.generatePrefix()
		if err != nil {
			s.Error("Unexpected error while generating prefix for the new account", "Details", err)
			return nil, err
		}
	}

	// Generate a password
	password, hashedPassword := util.GeneratePassword()

	// Insert the new account
	err = s.InsertNewAccount(context.Background(), db.InsertNewAccountParams{
		Prefix:         prefix,
		HashedPassword: hashedPassword,
		Username:       params.Username,
		Email:          params.Email,
	})
	if err != nil {
		s.Error("Unexpected error while inserting new account in DB", "Details", err)
		return nil, err
	}

	// Return response
	s.Info("New account created successfully", "Params", *params)
	return &CreateAccountResponse{
		Password: password,
		Prefix:   prefix,
	}, nil
}

func validateCreateAccountParams(params *CreateAccountParams) error {

	if params == nil {
		return errors.New("missing parameters")
	}

	if params.Username == "" {
		return errors.New("missing username")
	}

	err := util.ValidateUsername(params.Username)
	if err != nil {
		return err
	}

	if params.Email == "" {
		return errors.New("missing email")
	}

	err = util.ValidateEmail(params.Email)
	if err != nil {
		return err
	}

	return nil
}

func (s *AccountsService) generatePrefix() (bool, string, error) {
	prefix := util.GenerateRandomString(3)
	// Check if the prefix is unique
	exists, err := s.CheckPrefixExists(context.Background(), prefix)
	if err != nil {
		return false, "", err
	}
	return exists, prefix, nil
}

func (s *AccountsService) Update(params *UpdateAccountParams, user *User) (*UpdateAccountResponse, error) {

	// Check if the user is authenticated
	if user == nil {
		return nil, CoreError{
			Code:    ErrUnauthorized,
			Message: "unauthorized access",
		}
	}

	// This service is only accessible by the admin
	if !user.IsAdmin {
		s.Error("Unauthorized attempt to access UpdateAccount by a non admin user", "User", user)
		return nil, CoreError{
			Code:    ErrUnauthorized,
			Message: "unauthorized access",
		}
	}

	// Validate params
	err := validateUpdateAccountParams(params)
	if err != nil {
		s.Error("Bad UpdateAccount params", "Details", err)
		return nil, CoreError{
			Code:    ErrBadParams,
			Message: err.Error(),
		}
	}

	// Check if the account exists
	_, err = s.GetAccountByUsername(context.Background(), params.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			s.Error("Account not found for this username", "Username", params.Username)
			return nil, CoreError{
				Code:    ErrAccountNotFound,
				Message: "account not found",
			}
		}
		s.Error("Unexpected error while retrieving account by username", "Details", err)
		return nil, err
	}

	//Update account
	err = s.UpdateAccountStatusByUsername(context.Background(), db.UpdateAccountStatusByUsernameParams{
		Username: params.Username,
		Enabled:  params.Enabled,
	})
	if err != nil {
		s.Error("Unexpected error while updating account", "Details", err)
		return nil, err
	}

	s.Info("Account updated successfully", "Params", *params)
	return &UpdateAccountResponse{
		Enabled: params.Enabled,
	}, nil
}

func validateUpdateAccountParams(params *UpdateAccountParams) error {
	if params == nil {
		return errors.New("missing parameters")
	}
	if params.Username == "" {
		return errors.New("missing username")
	}
	return nil
}
