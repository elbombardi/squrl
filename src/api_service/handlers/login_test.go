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

func TestHandleLoginEmptyParameter(t *testing.T) {
	helper, err := setup()
	assert.NoError(t, err, "Error while setting up the test")
	ts := httptest.NewServer(helper.Handler)
	defer ts.Close()

	helper.Authenticator.On("Authenticate", mock.Anything, mock.Anything).Return("token", nil)

	reqBody := []byte{} // <== empty request body
	req, err := http.NewRequest("POST", ts.URL+"/v1/login", bytes.NewReader(reqBody))
	assert.NoError(t, err, "Error while creating request")

	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	assert.NoError(t, err, "Error while sending request")
	assert.Equal(t, http.StatusBadRequest, res.StatusCode, "Unexpected status code")

}

func TestHandleLoginBadJson(t *testing.T) {
	helper, err := setup()
	assert.NoError(t, err, "Error while setting up the test")
	ts := httptest.NewServer(helper.Handler)
	defer ts.Close()

	helper.Authenticator.On("Authenticate", mock.Anything, mock.Anything).Return("token", nil)

	reqBody := []byte("{badjson}") // <== empty request body
	req, err := http.NewRequest("POST", ts.URL+"/v1/login", bytes.NewReader(reqBody))
	assert.NoError(t, err, "Error while creating request")

	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	assert.NoError(t, err, "Error while sending request")
	assert.Equal(t, http.StatusBadRequest, res.StatusCode, "Unexpected status code")

}

func TestHandleLoginBadParams(t *testing.T) {
	helper, err := setup()
	assert.NoError(t, err, "Error while setting up the test")
	ts := httptest.NewServer(helper.Handler)
	defer ts.Close()

	helper.Authenticator.On("Authenticate", mock.Anything, mock.Anything).Return("", core.CoreError{
		Code:    core.ErrBadParams,
		Message: "Bad params",
	})

	reqBody, err := json.Marshal(models.LoginInfo{
		Password: "password",
		Username: "admin",
	})
	req, err := http.NewRequest("POST", ts.URL+"/v1/login", bytes.NewReader(reqBody))
	assert.NoError(t, err, "Error while creating request")

	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	assert.NoError(t, err, "Error while sending request")
	assert.Equal(t, http.StatusBadRequest, res.StatusCode, "Unexpected status code")

}

func TestHandleLoginAccountNotFound(t *testing.T) {
	helper, err := setup()
	assert.NoError(t, err, "Error while setting up the test")
	ts := httptest.NewServer(helper.Handler)
	defer ts.Close()

	helper.Authenticator.On("Authenticate", mock.Anything, mock.Anything).Return("", core.CoreError{
		Code:    core.ErrAccountNotFound,
		Message: "Bad params",
	})

	reqBody, err := json.Marshal(models.LoginInfo{
		Password: "password",
		Username: "admin",
	})
	req, err := http.NewRequest("POST", ts.URL+"/v1/login", bytes.NewReader(reqBody))
	assert.NoError(t, err, "Error while creating request")

	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	assert.NoError(t, err, "Error while sending request")
	assert.Equal(t, http.StatusUnauthorized, res.StatusCode, "Unexpected status code")
}

func TestHandleLoginAccountDisabled(t *testing.T) {
	helper, err := setup()
	assert.NoError(t, err, "Error while setting up the test")
	ts := httptest.NewServer(helper.Handler)
	defer ts.Close()

	helper.Authenticator.On("Authenticate", mock.Anything, mock.Anything).Return("", core.CoreError{
		Code:    core.ErrAccountDisabled,
		Message: "Bad params",
	})

	reqBody, err := json.Marshal(models.LoginInfo{
		Password: "password",
		Username: "admin",
	})
	req, err := http.NewRequest("POST", ts.URL+"/v1/login", bytes.NewReader(reqBody))
	assert.NoError(t, err, "Error while creating request")

	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	assert.NoError(t, err, "Error while sending request")
	assert.Equal(t, http.StatusUnauthorized, res.StatusCode, "Unexpected status code")
}

func TestHandleLoginAccountUnexpectedError(t *testing.T) {
	helper, err := setup()
	assert.NoError(t, err, "Error while setting up the test")
	ts := httptest.NewServer(helper.Handler)
	defer ts.Close()

	helper.Authenticator.On("Authenticate", mock.Anything, mock.Anything).Return("", errors.New("Unexpected error"))

	reqBody, err := json.Marshal(models.LoginInfo{
		Password: "password",
		Username: "admin",
	})
	req, err := http.NewRequest("POST", ts.URL+"/v1/login", bytes.NewReader(reqBody))
	assert.NoError(t, err, "Error while creating request")

	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	assert.NoError(t, err, "Error while sending request")
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode, "Unexpected status code")
}

func TestHandleLoginOK(t *testing.T) {
	helper, err := setup()
	assert.NoError(t, err, "Error while setting up the test")
	ts := httptest.NewServer(helper.Handler)
	defer ts.Close()

	helper.Authenticator.On("Authenticate", mock.Anything, mock.Anything).Return("token", nil)

	reqBody, err := json.Marshal(models.LoginInfo{
		Password: "password",
		Username: "admin",
	})
	req, err := http.NewRequest("POST", ts.URL+"/v1/login", bytes.NewReader(reqBody))
	assert.NoError(t, err, "Error while creating request")

	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	assert.NoError(t, err, "Error while sending request")
	assert.Equal(t, http.StatusOK, res.StatusCode, "Unexpected status code")

	respBody, _ := io.ReadAll(res.Body)

	loginSuccess := &models.LoginSuccess{}
	err = json.Unmarshal(respBody, loginSuccess)
	assert.NoError(t, err, "Error while unmarshalling response body")

	assert.Equal(t, "token", loginSuccess.Token, "Unexpected token")
	assert.Equal(t, true, loginSuccess.Success, "Unexpected success value")

}
