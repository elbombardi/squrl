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

func TestHandleUpdateLinkEmptyBody(t *testing.T) {
	helper, err := setup()
	assert.NoError(t, err, "Error while setting up the test")
	ts := httptest.NewServer(helper.Handler)
	defer ts.Close()

	helper.Authenticator.On("Validate", mock.Anything, mock.Anything).Return(&core.User{
		Username: "admin",
		IsAdmin:  true,
	}, nil)

	parsedUrl, _ := url.Parse("https://www.google.com")
	helper.LinksManager.On("Update", mock.Anything, mock.Anything).Return(&core.Link{
		Enabled:         false,
		TrackingEnabled: false,
		LongUrl:         *parsedUrl,
	}, nil)

	reqBody := []byte("")
	req, err := http.NewRequest("PUT", ts.URL+"/v1/links", bytes.NewReader(reqBody))
	assert.NoError(t, err, "Error while creating request")

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer token")

	res, err := http.DefaultClient.Do(req)
	assert.NoError(t, err, "Error while sending request")
	assert.Equal(t, http.StatusBadRequest, res.StatusCode, "Unexpected status code")
}

func TestHandleUpdateLinkBadJson(t *testing.T) {
	helper, err := setup()
	assert.NoError(t, err, "Error while setting up the test")
	ts := httptest.NewServer(helper.Handler)
	defer ts.Close()

	helper.Authenticator.On("Validate", mock.Anything, mock.Anything).Return(&core.User{
		Username: "admin",
		IsAdmin:  true,
	}, nil)

	parsedUrl, _ := url.Parse("https://www.google.com")
	helper.LinksManager.On("Update", mock.Anything, mock.Anything).Return(&core.Link{
		Enabled:         false,
		TrackingEnabled: false,
		LongUrl:         *parsedUrl,
	}, nil)

	reqBody := []byte("{badjson}")
	req, err := http.NewRequest("PUT", ts.URL+"/v1/links", bytes.NewReader(reqBody))
	assert.NoError(t, err, "Error while creating request")

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer token")

	res, err := http.DefaultClient.Do(req)
	assert.NoError(t, err, "Error while sending request")
	assert.Equal(t, http.StatusBadRequest, res.StatusCode, "Unexpected status code")

}

func TestHandleUpdateLinkBadStatus(t *testing.T) {
	helper, err := setup()
	assert.NoError(t, err, "Error while setting up the test")
	ts := httptest.NewServer(helper.Handler)
	defer ts.Close()

	helper.Authenticator.On("Validate", mock.Anything, mock.Anything).Return(&core.User{
		Username: "admin",
		IsAdmin:  true,
	}, nil)

	parsedUrl, _ := url.Parse("https://www.google.com")
	helper.LinksManager.On("Update", mock.Anything, mock.Anything).Return(&core.Link{
		Enabled:         false,
		TrackingEnabled: false,
		LongUrl:         *parsedUrl,
	}, nil)

	reqBody := []byte(`{
		"new_long_url": "https://www.google.com",
		"status": "badstatus",
		"short_url_key": "google",
		"tracking_status": "inactive"
	}`)
	req, err := http.NewRequest("PUT", ts.URL+"/v1/links", bytes.NewReader(reqBody))
	assert.NoError(t, err, "Error while creating request")

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer token")

	res, err := http.DefaultClient.Do(req)
	assert.NoError(t, err, "Error while sending request")
	assert.Equal(t, http.StatusBadRequest, res.StatusCode, "Unexpected status code")

}

func TestHandleUpdateLinkBadTrackingStatus(t *testing.T) {
	helper, err := setup()
	assert.NoError(t, err, "Error while setting up the test")
	ts := httptest.NewServer(helper.Handler)
	defer ts.Close()

	helper.Authenticator.On("Validate", mock.Anything, mock.Anything).Return(&core.User{
		Username: "admin",
		IsAdmin:  true,
	}, nil)

	parsedUrl, _ := url.Parse("https://www.google.com")
	helper.LinksManager.On("Update", mock.Anything, mock.Anything).Return(&core.Link{
		Enabled:         false,
		TrackingEnabled: false,
		LongUrl:         *parsedUrl,
	}, nil)

	reqBody := []byte(`{
		"new_long_url": "https://www.google.com",
		"status": "inactive",
		"short_url_key": "google",
		"tracking_status": "badstatus"
	}`)
	req, err := http.NewRequest("PUT", ts.URL+"/v1/links", bytes.NewReader(reqBody))
	assert.NoError(t, err, "Error while creating request")

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer token")

	res, err := http.DefaultClient.Do(req)
	assert.NoError(t, err, "Error while sending request")
	assert.Equal(t, http.StatusBadRequest, res.StatusCode, "Unexpected status code")

}

func TestHandleUpdateLinkAccountNotFound(t *testing.T) {
	helper, err := setup()
	assert.NoError(t, err, "Error while setting up the test")
	ts := httptest.NewServer(helper.Handler)
	defer ts.Close()

	helper.Authenticator.On("Validate", mock.Anything, mock.Anything).Return(&core.User{
		Username: "admin",
		IsAdmin:  true,
	}, nil)

	helper.LinksManager.On("Update", mock.Anything, mock.Anything).Return((*core.Link)(nil), core.CoreError{
		Code:    core.ErrAccountNotFound,
		Message: "Account not found",
	})

	reqBody, err := json.Marshal(models.LinkUpdate{
		NewLongURL:     "https://www.google.com",
		Status:         "inactive",
		ShortURLKey:    "google",
		TrackingStatus: "inactive",
	})
	req, err := http.NewRequest("PUT", ts.URL+"/v1/links", bytes.NewReader(reqBody))
	assert.NoError(t, err, "Error while creating request")

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer token")

	res, err := http.DefaultClient.Do(req)
	assert.NoError(t, err, "Error while sending request")
	assert.Equal(t, http.StatusUnauthorized, res.StatusCode, "Unexpected status code")
}

func TestHandleUpdateLinkAccountDisabled(t *testing.T) {
	helper, err := setup()
	assert.NoError(t, err, "Error while setting up the test")
	ts := httptest.NewServer(helper.Handler)
	defer ts.Close()

	helper.Authenticator.On("Validate", mock.Anything, mock.Anything).Return(&core.User{
		Username: "admin",
		IsAdmin:  true,
	}, nil)

	helper.LinksManager.On("Update", mock.Anything, mock.Anything).Return((*core.Link)(nil), core.CoreError{
		Code:    core.ErrAccountDisabled,
		Message: "Account not found",
	})

	reqBody, err := json.Marshal(models.LinkUpdate{
		NewLongURL:     "https://www.google.com",
		Status:         "inactive",
		ShortURLKey:    "google",
		TrackingStatus: "inactive",
	})
	req, err := http.NewRequest("PUT", ts.URL+"/v1/links", bytes.NewReader(reqBody))
	assert.NoError(t, err, "Error while creating request")

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer token")

	res, err := http.DefaultClient.Do(req)
	assert.NoError(t, err, "Error while sending request")
	assert.Equal(t, http.StatusUnauthorized, res.StatusCode, "Unexpected status code")
}

func TestHandleUpdateLinkLinkNotFOund(t *testing.T) {
	helper, err := setup()
	assert.NoError(t, err, "Error while setting up the test")
	ts := httptest.NewServer(helper.Handler)
	defer ts.Close()

	helper.Authenticator.On("Validate", mock.Anything, mock.Anything).Return(&core.User{
		Username: "admin",
		IsAdmin:  true,
	}, nil)

	helper.LinksManager.On("Update", mock.Anything, mock.Anything).Return((*core.Link)(nil), core.CoreError{
		Code:    core.ErrLinkNotFound,
		Message: "Account not found",
	})

	reqBody, err := json.Marshal(models.LinkUpdate{
		NewLongURL:     "https://www.google.com",
		Status:         "inactive",
		ShortURLKey:    "google",
		TrackingStatus: "inactive",
	})
	req, err := http.NewRequest("PUT", ts.URL+"/v1/links", bytes.NewReader(reqBody))
	assert.NoError(t, err, "Error while creating request")

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer token")

	res, err := http.DefaultClient.Do(req)
	assert.NoError(t, err, "Error while sending request")
	assert.Equal(t, http.StatusNotFound, res.StatusCode, "Unexpected status code")
}

func TestHandleUpdateLinkUnexpectedError(t *testing.T) {
	helper, err := setup()
	assert.NoError(t, err, "Error while setting up the test")
	ts := httptest.NewServer(helper.Handler)
	defer ts.Close()

	helper.Authenticator.On("Validate", mock.Anything, mock.Anything).Return(&core.User{
		Username: "admin",
		IsAdmin:  true,
	}, nil)

	helper.LinksManager.On("Update", mock.Anything, mock.Anything).Return((*core.Link)(nil), errors.New("Unexpected error"))

	reqBody, err := json.Marshal(models.LinkUpdate{
		NewLongURL:     "https://www.google.com",
		Status:         "inactive",
		ShortURLKey:    "google",
		TrackingStatus: "inactive",
	})
	req, err := http.NewRequest("PUT", ts.URL+"/v1/links", bytes.NewReader(reqBody))
	assert.NoError(t, err, "Error while creating request")

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer token")

	res, err := http.DefaultClient.Do(req)
	assert.NoError(t, err, "Error while sending request")
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode, "Unexpected status code")
}

func TestHandleUpdateLinkOK(t *testing.T) {
	helper, err := setup()
	assert.NoError(t, err, "Error while setting up the test")
	ts := httptest.NewServer(helper.Handler)
	defer ts.Close()

	helper.Authenticator.On("Validate", mock.Anything, mock.Anything).Return(&core.User{
		Username: "admin",
		IsAdmin:  true,
	}, nil)

	parsedUrl, _ := url.Parse("https://www.google.com")
	helper.LinksManager.On("Update", mock.Anything, mock.Anything).Return(&core.Link{
		Enabled:         false,
		TrackingEnabled: false,
		LongUrl:         *parsedUrl,
	}, nil)

	reqBody, err := json.Marshal(models.LinkUpdate{
		NewLongURL:     "https://www.google.com",
		Status:         "inactive",
		ShortURLKey:    "google",
		TrackingStatus: "inactive",
	})
	req, err := http.NewRequest("PUT", ts.URL+"/v1/links", bytes.NewReader(reqBody))
	assert.NoError(t, err, "Error while creating request")

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer token")

	res, err := http.DefaultClient.Do(req)
	assert.NoError(t, err, "Error while sending request")
	assert.Equal(t, http.StatusOK, res.StatusCode, "Unexpected status code")

	respBody, _ := io.ReadAll(res.Body)

	linkUpdated := &models.LinkUpdated{}
	err = json.Unmarshal(respBody, linkUpdated)
	assert.NoError(t, err, "Error while unmarshalling response body")

	assert.Equal(t, "https://www.google.com", linkUpdated.LongURL, "Unexpected long URL")
	assert.Equal(t, "inactive", linkUpdated.Status, "Unexpected status")
	assert.Equal(t, "inactive", linkUpdated.TrackingStatus, "Unexpected tracking status")
}
