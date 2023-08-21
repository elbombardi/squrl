package handlers

import (
	"bytes"
	"encoding/json"
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

func TestHandleCreateAccountOK(t *testing.T) {
	helper, err := setup()
	assert.NoError(t, err, "Error while setting up the test")
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
	assert.NoError(t, err, "Error while creating request")

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer token")

	res, err := http.DefaultClient.Do(req)
	assert.NoError(t, err, "Error while sending request")
	assert.Equal(t, http.StatusOK, res.StatusCode, "Unexpected status code")

	respBody, _ := io.ReadAll(res.Body)

	accountCreated := &models.AccountCreated{}
	err = json.Unmarshal(respBody, accountCreated)
	assert.NoError(t, err, "Error while unmarshalling response body")

	assert.Equal(t, "password", accountCreated.Password, "Unexpected password")
	assert.Equal(t, "prefix", accountCreated.Prefix, "Unexpected prefix")
}
