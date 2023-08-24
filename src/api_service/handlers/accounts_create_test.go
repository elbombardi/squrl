package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/elbombardi/squrl/src/api_service/api"
	"github.com/elbombardi/squrl/src/api_service/api/models"
	"github.com/elbombardi/squrl/src/api_service/api/operations"
	"github.com/elbombardi/squrl/src/api_service/core"
	mocks_core "github.com/elbombardi/squrl/src/api_service/mocks/core"
	mocks_util "github.com/elbombardi/squrl/src/api_service/mocks/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/elbombardi/squrl/src/api_service/util"
	"github.com/go-openapi/loads"
)

type TestHelper struct {
	Handlers        *Handlers
	AccountsManager *mocks_core.MockAccountsManager
	Authenticator   *mocks_core.MockAuthenticator
	LinksManager    *mocks_core.MockLinksManager
	Config          *util.Config
	Handler         http.Handler
}

func setup() (*TestHelper, error) {
	swaggerSpec, err := loads.Analyzed(api.SwaggerJSON, "")
	if err != nil {
		return nil, err
	}
	adminAPI := operations.NewAdminAPI(swaggerSpec)
	if err != nil {
		return nil, err
	}
	server := api.NewServer(adminAPI)

	helper := &TestHelper{
		AccountsManager: new(mocks_core.MockAccountsManager),
		Authenticator:   new(mocks_core.MockAuthenticator),
		LinksManager:    new(mocks_core.MockLinksManager),
		Config:          mocks_util.MockConfig(),
	}

	helper.Handlers = &Handlers{
		Authenticator:   helper.Authenticator,
		AccountsManager: helper.AccountsManager,
		LinksManager:    helper.LinksManager,
		Config:          helper.Config,
	}
	helper.Handlers.InstallHandlers(adminAPI)

	server.ConfigureAPI()
	err = adminAPI.Validate()
	if err != nil {
		return nil, err
	}
	helper.Handler = server.GetHandler()

	return helper, nil
}

func TestHandleCreateAccountWithUnexpectedError(t *testing.T) {
	helper, err := setup()
	require.NoError(t, err, "Error while setting up the test")
	ts := httptest.NewServer(helper.Handler)
	defer ts.Close()

	helper.Authenticator.On("Validate", mock.Anything, mock.Anything).Return(&core.User{
		Username: "admin",
		IsAdmin:  true,
	}, nil)

	helper.AccountsManager.On("Create", mock.Anything, mock.Anything).Return((*core.CreateAccountResponse)(nil), errors.New("unexpected error"))

	reqBody, err := json.Marshal(models.Account{
		Username: "test",
		Email:    "test@gmail.com",
	})
	req, err := http.NewRequest("POST", ts.URL+"/v1/accounts", bytes.NewReader(reqBody))
	if err != nil {
		fmt.Println(err)
	}
	require.NoError(t, err, "Error while creating request")

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer token")

	res, err := http.DefaultClient.Do(req)
	assert.NoError(t, err, "Expected error while sending request")
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode, "Unexpected status code")
}

func TestHandleCreateAccountWithAuthorizationError(t *testing.T) {
	helper, err := setup()
	require.NoError(t, err, "Error while setting up the test")
	ts := httptest.NewServer(helper.Handler)
	defer ts.Close()

	helper.Authenticator.On("Validate", mock.Anything, mock.Anything).Return((*core.User)(nil), nil)

	helper.AccountsManager.On("Create", mock.Anything, mock.Anything).Return((*core.CreateAccountResponse)(nil), core.CoreError{
		Code:    core.ErrUnauthorized,
		Message: "Unauthorized access",
	})

	reqBody, err := json.Marshal(models.Account{
		Username: "test",
		Email:    "test@gmail.com",
	})
	req, err := http.NewRequest("POST", ts.URL+"/v1/accounts", bytes.NewReader(reqBody))
	require.NoError(t, err, "Error while creating request")

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer token")

	res, err := http.DefaultClient.Do(req)
	fmt.Println(err)
	require.NoError(t, err, "Error while sending request")
	require.Equal(t, http.StatusUnauthorized, res.StatusCode, "Unexpected status code")
}

func TestHandleCreateAccountWithoutToken(t *testing.T) {
	helper, err := setup()
	require.NoError(t, err, "Error while setting up the test")
	ts := httptest.NewServer(helper.Handler)
	defer ts.Close()

	helper.Authenticator.On("Validate", mock.Anything, mock.Anything).Return(&core.User{
		Username: "admin",
		IsAdmin:  true,
	}, nil)

	helper.AccountsManager.On("Create", mock.Anything, mock.Anything).Return(&core.CreateAccountResponse{
		Password: "password",
		Prefix:   "prefix",
	}, nil)

	reqBody, err := json.Marshal(models.Account{
		Username: "test",
		Email:    "test@gmail.com",
	})
	req, err := http.NewRequest("POST", ts.URL+"/v1/accounts", bytes.NewReader(reqBody))
	require.NoError(t, err, "Error while creating request")

	req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("Authorization", "Bearer token")

	res, err := http.DefaultClient.Do(req)
	require.NoError(t, err, "Error while sending request")
	require.Equal(t, http.StatusUnauthorized, res.StatusCode, "Unexpected status code")
}

func TestHandleCreateAccountBadParams(t *testing.T) {
	helper, err := setup()
	require.NoError(t, err, "Error while setting up the test")
	ts := httptest.NewServer(helper.Handler)
	defer ts.Close()

	helper.Authenticator.On("Validate", mock.Anything, mock.Anything).Return(&core.User{
		Username: "admin",
		IsAdmin:  true,
	}, nil)

	helper.AccountsManager.On("Create", mock.Anything, mock.Anything).Return((*core.CreateAccountResponse)(nil), core.CoreError{
		Code:    core.ErrBadParams,
		Message: "Bad params",
	})

	reqBody, err := json.Marshal(models.Account{
		Username: "test",
		Email:    "test@gmail.com",
	})
	req, err := http.NewRequest("POST", ts.URL+"/v1/accounts", bytes.NewReader(reqBody))
	require.NoError(t, err, "Error while creating request")

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer token")

	res, err := http.DefaultClient.Do(req)
	require.NoError(t, err, "Error while sending request")
	require.Equal(t, http.StatusBadRequest, res.StatusCode, "Unexpected status code")
}

func TestHandleCreateAccountBadJson(t *testing.T) {
	helper, err := setup()
	require.NoError(t, err, "Error while setting up the test")
	ts := httptest.NewServer(helper.Handler)
	defer ts.Close()

	helper.Authenticator.On("Validate", mock.Anything, mock.Anything).Return(&core.User{
		Username: "admin",
		IsAdmin:  true,
	}, nil)

	helper.AccountsManager.On("Create", mock.Anything, mock.Anything).Return(&core.CreateAccountResponse{
		Password: "password",
		Prefix:   "prefix",
	}, nil)

	reqBody := []byte("bad json") // <==== Bad JSON
	req, err := http.NewRequest("POST", ts.URL+"/v1/accounts", bytes.NewReader(reqBody))
	require.NoError(t, err, "Error while creating request")

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer token")

	res, err := http.DefaultClient.Do(req)
	require.NoError(t, err, "Error while sending request")
	require.Equal(t, http.StatusBadRequest, res.StatusCode, "Unexpected status code")
}

func TestHandleCreateAccountWithNoParams(t *testing.T) {
	helper, err := setup()
	require.NoError(t, err, "Error while setting up the test")
	ts := httptest.NewServer(helper.Handler)
	defer ts.Close()

	helper.Authenticator.On("Validate", mock.Anything, mock.Anything).Return(&core.User{
		Username: "admin",
		IsAdmin:  true,
	}, nil)

	helper.AccountsManager.On("Create", mock.Anything, mock.Anything).Return(&core.CreateAccountResponse{
		Password: "password",
		Prefix:   "prefix",
	}, nil)

	reqBody := []byte("") // <==== empty body
	req, err := http.NewRequest("POST", ts.URL+"/v1/accounts", bytes.NewReader(reqBody))
	require.NoError(t, err, "Error while creating request")

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer token")

	res, err := http.DefaultClient.Do(req)
	require.NoError(t, err, "Error while sending request")
	require.Equal(t, http.StatusBadRequest, res.StatusCode, "Unexpected status code")
}

func TestHandleCreateAccountOK(t *testing.T) {
	helper, err := setup()
	require.NoError(t, err, "Error while setting up the test")
	ts := httptest.NewServer(helper.Handler)
	defer ts.Close()

	helper.Authenticator.On("Validate", mock.Anything, mock.Anything).Return(&core.User{
		Username: "admin",
		IsAdmin:  true,
	}, nil)

	helper.AccountsManager.On("Create", mock.Anything, mock.Anything).Return(&core.CreateAccountResponse{
		Password: "password",
		Prefix:   "prefix",
	}, nil)

	reqBody, err := json.Marshal(models.Account{
		Username: "test",
		Email:    "test@gmail.com",
	})
	req, err := http.NewRequest("POST", ts.URL+"/v1/accounts", bytes.NewReader(reqBody))
	require.NoError(t, err, "Error while creating request")

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer token")

	res, err := http.DefaultClient.Do(req)
	require.NoError(t, err, "Error while sending request")
	require.Equal(t, http.StatusOK, res.StatusCode, "Unexpected status code")

	respBody, _ := io.ReadAll(res.Body)

	accountCreated := &models.AccountCreated{}
	err = json.Unmarshal(respBody, accountCreated)
	require.NoError(t, err, "Error while unmarshalling response body")

	require.Equal(t, "password", accountCreated.Password, "Unexpected password")
	require.Equal(t, "prefix", accountCreated.Prefix, "Unexpected prefix")
}
