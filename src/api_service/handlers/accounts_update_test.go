package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/elbombardi/squrl/src/api_service/api/models"
	"github.com/elbombardi/squrl/src/api_service/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

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
