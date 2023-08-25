package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/elbombardi/squrl/src/api_service/api/models"
	"github.com/elbombardi/squrl/src/api_service/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandleCreateEmptyBody(t *testing.T) {
	helper, err := setup()
	assert.NoError(t, err, "Error while setting up the test")
	ts := httptest.NewServer(helper.Handler)
	defer ts.Close()

	helper.Authenticator.On("Validate", mock.Anything, mock.Anything).Return(&core.User{
		Username: "admin",
		IsAdmin:  true,
	}, nil)

	parsedUrl, _ := url.Parse("https://www.google.com")
	helper.LinksManager.On("Shorten", mock.Anything, mock.Anything).Return(&core.Link{
		LongUrl:         *parsedUrl,
		Enabled:         false,
		TrackingEnabled: false,
		ShortUrlKey:     "google",
	}, nil)

	reqBody := []byte{}
	req, err := http.NewRequest("POST", ts.URL+"/v1/links", bytes.NewReader(reqBody))
	assert.NoError(t, err, "Error while creating request")

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer token")

	res, err := http.DefaultClient.Do(req)
	assert.NoError(t, err, "Error while sending request")
	assert.Equal(t, http.StatusBadRequest, res.StatusCode, "Unexpected status code")

}

func TestHandleCreateLinkBadJson(t *testing.T) {
	helper, err := setup()
	assert.NoError(t, err, "Error while setting up the test")
	ts := httptest.NewServer(helper.Handler)
	defer ts.Close()

	helper.Authenticator.On("Validate", mock.Anything, mock.Anything).Return(&core.User{
		Username: "admin",
		IsAdmin:  true,
	}, nil)

	parsedUrl, _ := url.Parse("https://www.google.com")
	helper.LinksManager.On("Shorten", mock.Anything, mock.Anything).Return(&core.Link{
		LongUrl:         *parsedUrl,
		Enabled:         false,
		TrackingEnabled: false,
		ShortUrlKey:     "google",
	}, nil)

	reqBody := []byte(`{long_url`)
	req, err := http.NewRequest("POST", ts.URL+"/v1/links", bytes.NewReader(reqBody))
	assert.NoError(t, err, "Error while creating request")

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer token")

	res, err := http.DefaultClient.Do(req)
	assert.NoError(t, err, "Error while sending request")
	assert.Equal(t, http.StatusBadRequest, res.StatusCode, "Unexpected status code")
}

func TestHandleCreateLinkBadParams(t *testing.T) {
	helper, err := setup()
	assert.NoError(t, err, "Error while setting up the test")
	ts := httptest.NewServer(helper.Handler)
	defer ts.Close()

	helper.Authenticator.On("Validate", mock.Anything, mock.Anything).Return(&core.User{
		Username: "admin",
		IsAdmin:  true,
	}, nil)

	helper.LinksManager.On("Shorten", mock.Anything, mock.Anything).Return((*core.Link)(nil),
		core.CoreError{
			Code:    core.ErrBadParams,
			Message: "Bad params",
		})

	reqBody, err := json.Marshal(models.Link{
		LongURL: "https://www.google.com",
	})
	req, err := http.NewRequest("POST", ts.URL+"/v1/links", bytes.NewReader(reqBody))
	assert.NoError(t, err, "Error while creating request")

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer token")

	res, err := http.DefaultClient.Do(req)
	assert.NoError(t, err, "Error while sending request")
	assert.Equal(t, http.StatusBadRequest, res.StatusCode, "Unexpected status code")
}

func TestHandleCreateLinkUnauthorized(t *testing.T) {
	helper, err := setup()
	assert.NoError(t, err, "Error while setting up the test")
	ts := httptest.NewServer(helper.Handler)
	defer ts.Close()

	helper.Authenticator.On("Validate", mock.Anything, mock.Anything).Return(&core.User{
		Username: "admin",
		IsAdmin:  true,
	}, nil)

	helper.LinksManager.On("Shorten", mock.Anything, mock.Anything).Return((*core.Link)(nil),
		core.CoreError{
			Code:    core.ErrUnauthorized,
			Message: "Unauthorized",
		})

	reqBody, err := json.Marshal(models.Link{
		LongURL: "https://www.google.com",
	})
	req, err := http.NewRequest("POST", ts.URL+"/v1/links", bytes.NewReader(reqBody))
	assert.NoError(t, err, "Error while creating request")

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer token")

	res, err := http.DefaultClient.Do(req)
	assert.NoError(t, err, "Error while sending request")
	assert.Equal(t, http.StatusUnauthorized, res.StatusCode, "Unexpected status code")
}

func TestHandleCreateLinkAccountNotFound(t *testing.T) {
	helper, err := setup()
	assert.NoError(t, err, "Error while setting up the test")
	ts := httptest.NewServer(helper.Handler)
	defer ts.Close()

	helper.Authenticator.On("Validate", mock.Anything, mock.Anything).Return(&core.User{
		Username: "admin",
		IsAdmin:  true,
	}, nil)

	helper.LinksManager.On("Shorten", mock.Anything, mock.Anything).Return((*core.Link)(nil),
		core.CoreError{
			Code:    core.ErrAccountNotFound,
			Message: "Account not found",
		})

	reqBody, err := json.Marshal(models.Link{
		LongURL: "https://www.google.com",
	})
	req, err := http.NewRequest("POST", ts.URL+"/v1/links", bytes.NewReader(reqBody))
	assert.NoError(t, err, "Error while creating request")

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer token")

	res, err := http.DefaultClient.Do(req)
	assert.NoError(t, err, "Error while sending request")
	assert.Equal(t, http.StatusUnauthorized, res.StatusCode, "Unexpected status code")
}

func TestHandleCreateLinkAccountDisabled(t *testing.T) {
	helper, err := setup()
	assert.NoError(t, err, "Error while setting up the test")
	ts := httptest.NewServer(helper.Handler)
	defer ts.Close()

	helper.Authenticator.On("Validate", mock.Anything, mock.Anything).Return(&core.User{
		Username: "admin",
		IsAdmin:  true,
	}, nil)

	helper.LinksManager.On("Shorten", mock.Anything, mock.Anything).Return((*core.Link)(nil),
		core.CoreError{
			Code:    core.ErrAccountDisabled,
			Message: "Account not found",
		})

	reqBody, err := json.Marshal(models.Link{
		LongURL: "https://www.google.com",
	})
	req, err := http.NewRequest("POST", ts.URL+"/v1/links", bytes.NewReader(reqBody))
	assert.NoError(t, err, "Error while creating request")

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer token")

	res, err := http.DefaultClient.Do(req)
	assert.NoError(t, err, "Error while sending request")
	assert.Equal(t, http.StatusUnauthorized, res.StatusCode, "Unexpected status code")
}

func TestHandleCreateLinkUnexpectedError(t *testing.T) {
	helper, err := setup()
	assert.NoError(t, err, "Error while setting up the test")
	ts := httptest.NewServer(helper.Handler)
	defer ts.Close()

	helper.Authenticator.On("Validate", mock.Anything, mock.Anything).Return(&core.User{
		Username: "admin",
		IsAdmin:  true,
	}, nil)

	helper.LinksManager.On("Shorten", mock.Anything, mock.Anything).Return((*core.Link)(nil),
		errors.New("Unexpected error"))

	reqBody, err := json.Marshal(models.Link{
		LongURL: "https://www.google.com",
	})
	req, err := http.NewRequest("POST", ts.URL+"/v1/links", bytes.NewReader(reqBody))
	assert.NoError(t, err, "Error while creating request")

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer token")

	res, err := http.DefaultClient.Do(req)
	assert.NoError(t, err, "Error while sending request")
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode, "Unexpected status code")
}

func TestHandleCreateLinkOK(t *testing.T) {
	helper, err := setup()
	assert.NoError(t, err, "Error while setting up the test")
	ts := httptest.NewServer(helper.Handler)
	defer ts.Close()

	helper.Authenticator.On("Validate", mock.Anything, mock.Anything).Return(&core.User{
		Username: "admin",
		IsAdmin:  true,
	}, nil)

	parsedUrl, _ := url.Parse("https://www.google.com")
	helper.LinksManager.On("Shorten", mock.Anything, mock.Anything).Return(&core.Link{
		LongUrl:         *parsedUrl,
		Enabled:         false,
		TrackingEnabled: false,
		ShortUrlKey:     "google",
	}, nil)

	reqBody, err := json.Marshal(models.Link{
		LongURL: "https://www.google.com",
	})
	req, err := http.NewRequest("POST", ts.URL+"/v1/links", bytes.NewReader(reqBody))
	assert.NoError(t, err, "Error while creating request")

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer token")

	res, err := http.DefaultClient.Do(req)
	assert.NoError(t, err, "Error while sending request")
	assert.Equal(t, http.StatusOK, res.StatusCode, "Unexpected status code")

	respBody, _ := io.ReadAll(res.Body)

	linkCreated := &models.LinkCreated{}
	err = json.Unmarshal(respBody, linkCreated)
	assert.NoError(t, err, "Error while unmarshalling response body")

	assert.Equal(t, "google", linkCreated.ShortURLKey)
}
