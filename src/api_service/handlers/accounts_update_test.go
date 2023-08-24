package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/elbombardi/squrl/src/api_service/api/models"
	"github.com/elbombardi/squrl/src/api_service/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandleUpdateAccountNoAuthorizationHeader(t *testing.T) {
	helper, err := setup()
	assert.NoError(t, err, "Error while setting up the test")
	ts := httptest.NewServer(helper.Handler)
	defer ts.Close()

	helper.Authenticator.On("Validate", mock.Anything, mock.Anything).Return(&core.User{
		Username: "admin",
		IsAdmin:  true,
	}, nil)

	helper.AccountsManager.On("Update", mock.Anything, mock.Anything).Return((*core.UpdateAccountResponse)(nil), errors.New("Unexpected error"))

	reqBody, err := json.Marshal(models.AccountUpdate{
		Username: "test",
		Status:   "inactive",
	})
	req, err := http.NewRequest("PUT", ts.URL+"/v1/accounts", bytes.NewReader(reqBody))
	assert.NoError(t, err, "Error while creating request")

	req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("Authorization", "Bearer token")

	res, err := http.DefaultClient.Do(req)
	assert.NoError(t, err, "Error while sending request")
	assert.Equal(t, http.StatusUnauthorized, res.StatusCode, "Unexpected status code")
}

func TestHandleUpdateAccountEmptyInput(t *testing.T) {
	helper, err := setup()
	assert.NoError(t, err, "Error while setting up the test")
	ts := httptest.NewServer(helper.Handler)
	defer ts.Close()

	helper.Authenticator.On("Validate", mock.Anything, mock.Anything).Return(&core.User{
		Username: "admin",
		IsAdmin:  true,
	}, nil)

	helper.AccountsManager.On("Update", mock.Anything, mock.Anything).Return((*core.UpdateAccountResponse)(nil), errors.New("Unexpected error"))

	reqBody := []byte{}
	req, err := http.NewRequest("PUT", ts.URL+"/v1/accounts", bytes.NewReader(reqBody))
	assert.NoError(t, err, "Error while creating request")

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer token")

	res, err := http.DefaultClient.Do(req)
	assert.NoError(t, err, "Error while sending request")
	assert.Equal(t, http.StatusBadRequest, res.StatusCode, "Unexpected status code")
}

func TestHandleUpdateAccountUnexpectedError(t *testing.T) {
	helper, err := setup()
	assert.NoError(t, err, "Error while setting up the test")
	ts := httptest.NewServer(helper.Handler)
	defer ts.Close()

	helper.Authenticator.On("Validate", mock.Anything, mock.Anything).Return(&core.User{
		Username: "admin",
		IsAdmin:  true,
	}, nil)

	helper.AccountsManager.On("Update", mock.Anything, mock.Anything).Return((*core.UpdateAccountResponse)(nil), errors.New("Unexpected error"))

	reqBody, err := json.Marshal(models.AccountUpdate{
		Username: "test",
		Status:   "inactive",
	})
	req, err := http.NewRequest("PUT", ts.URL+"/v1/accounts", bytes.NewReader(reqBody))
	assert.NoError(t, err, "Error while creating request")

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer token")

	res, err := http.DefaultClient.Do(req)
	assert.NoError(t, err, "Error while sending request")
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode, "Unexpected status code")
}

func TestHandleUpdateAccountUnauthorizedError(t *testing.T) {
	helper, err := setup()
	assert.NoError(t, err, "Error while setting up the test")
	ts := httptest.NewServer(helper.Handler)
	defer ts.Close()

	helper.Authenticator.On("Validate", mock.Anything, mock.Anything).Return(&core.User{
		Username: "admin",
		IsAdmin:  true,
	}, nil)

	helper.AccountsManager.On("Update", mock.Anything, mock.Anything).Return((*core.UpdateAccountResponse)(nil), core.CoreError{
		Code:    core.ErrUnauthorized,
		Message: "Unauthorized access",
	})

	reqBody, err := json.Marshal(models.AccountUpdate{
		Username: "test",
		Status:   "inactive",
	})
	req, err := http.NewRequest("PUT", ts.URL+"/v1/accounts", bytes.NewReader(reqBody))
	assert.NoError(t, err, "Error while creating request")

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer token")

	res, err := http.DefaultClient.Do(req)
	assert.NoError(t, err, "Error while sending request")
	assert.Equal(t, http.StatusUnauthorized, res.StatusCode, "Unexpected status code")
}

func TestHandleUpdateAccountOK(t *testing.T) {
	helper, err := setup()
	assert.NoError(t, err, "Error while setting up the test")
	ts := httptest.NewServer(helper.Handler)
	defer ts.Close()

	helper.Authenticator.On("Validate", mock.Anything, mock.Anything).Return(&core.User{
		Username: "admin",
		IsAdmin:  true,
	}, nil)

	helper.AccountsManager.On("Update", mock.Anything, mock.Anything).Return(&core.UpdateAccountResponse{
		Enabled: false,
	}, nil)

	reqBody, err := json.Marshal(models.AccountUpdate{
		Username: "test",
		Status:   "inactive",
	})
	req, err := http.NewRequest("PUT", ts.URL+"/v1/accounts", bytes.NewReader(reqBody))
	assert.NoError(t, err, "Error while creating request")

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer token")

	res, err := http.DefaultClient.Do(req)
	assert.NoError(t, err, "Error while sending request")
	assert.Equal(t, http.StatusOK, res.StatusCode, "Unexpected status code")

	respBody, _ := io.ReadAll(res.Body)

	accountUpdated := &models.AccountUpdated{}
	err = json.Unmarshal(respBody, accountUpdated)
	assert.NoError(t, err, "Error while unmarshalling response body")

	assert.Equal(t, "inactive", accountUpdated.Status, "Unexpected enabled status")
}
