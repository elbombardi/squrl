package routes

import (
	"errors"
	"log/slog"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/elbombardi/squrl/src/redirection_service/core"
	core_mocks "github.com/elbombardi/squrl/src/redirection_service/mocks/core"
	util_mock "github.com/elbombardi/squrl/src/redirection_service/mocks/util"
	"github.com/elbombardi/squrl/src/redirection_service/util"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type Helper struct {
	app          *fiber.App
	linksManager *core_mocks.MockLinksManager
	logger       *slog.Logger
	config       *util.Config
}

func setup() *Helper {
	s := &Helper{
		config:       util_mock.MockConfig(),
		linksManager: new(core_mocks.MockLinksManager),
		app:          fiber.New(),
	}

	routes := &Routes{
		LinksManager: s.linksManager,
		Config:       s.config,
	}
	s.app = fiber.New()
	s.app.Get("/*", routes.RedirectRoute)
	go func() {
		s.app.Listen(":3000")
	}()
	return s
}

func TestRedirectionMalformedUrl(t *testing.T) {
	s := setup()

	s.linksManager.On("Resolve", mock.Anything).Return(&core.Link{}, core.CoreError{Code: core.ErrAccountNotFound})

	req := httptest.NewRequest("GET", "/acc", nil)
	res, err := s.app.Test(req)

	require.NoError(t, err, "Error while testing the redirection route")
	defer res.Body.Close()
	require.Equal(t, 404, res.StatusCode, "Status code should be 404")
}

func TestRedirectionAccountNotFound(t *testing.T) {
	s := setup()

	s.linksManager.On("Resolve", mock.Anything).Return(&core.Link{}, core.CoreError{Code: core.ErrAccountNotFound})

	req := httptest.NewRequest("GET", "/acc/ur", nil)
	res, err := s.app.Test(req)

	require.NoError(t, err, "Error while testing the redirection route")
	defer res.Body.Close()
	require.Equal(t, 404, res.StatusCode, "Status code should be 404")
}

func TestRedirectionAccountDisabled(t *testing.T) {
	s := setup()

	s.linksManager.On("Resolve", mock.Anything).Return(&core.Link{}, core.CoreError{Code: core.ErrAccountDisabled})

	req := httptest.NewRequest("GET", "/acc/ur", nil)
	res, err := s.app.Test(req)

	require.NoError(t, err, "Error while testing the redirection route")
	defer res.Body.Close()
	require.Equal(t, 404, res.StatusCode, "Status code should be 404")
}

func TestRedirectionLinkNotFound(t *testing.T) {
	s := setup()

	s.linksManager.On("Resolve", mock.Anything).Return(&core.Link{}, core.CoreError{Code: core.ErrLinkNotFound})

	req := httptest.NewRequest("GET", "/acc/ur", nil)
	res, err := s.app.Test(req)

	require.NoError(t, err, "Error while testing the redirection route")
	defer res.Body.Close()
	require.Equal(t, 404, res.StatusCode, "Status code should be 404")
}

func TestRedirectionLinkDisabled(t *testing.T) {
	s := setup()

	s.linksManager.On("Resolve", mock.Anything).Return(&core.Link{}, core.CoreError{Code: core.ErrLinkDisabled})

	req := httptest.NewRequest("GET", "/acc/ur", nil)
	res, err := s.app.Test(req)

	require.NoError(t, err, "Error while testing the redirection route")
	defer res.Body.Close()
	require.Equal(t, 404, res.StatusCode, "Status code should be 404")
}

func TestRedirectionBadParams(t *testing.T) {
	s := setup()

	s.linksManager.On("Resolve", mock.Anything).Return(&core.Link{}, core.CoreError{Code: core.ErrBadParams})

	req := httptest.NewRequest("GET", "/acc/ur", nil)
	res, err := s.app.Test(req)

	require.NoError(t, err, "Error while testing the redirection route")
	defer res.Body.Close()
	require.Equal(t, 500, res.StatusCode, "Status code should be 500")
}

func TestRedirectionInternalError(t *testing.T) {
	s := setup()

	s.linksManager.On("Resolve", mock.Anything).Return(&core.Link{}, errors.New("internal error"))

	req := httptest.NewRequest("GET", "/acc/ur", nil)
	res, err := s.app.Test(req)

	require.NoError(t, err, "Error while testing the redirection route")
	defer res.Body.Close()
	require.Equal(t, 500, res.StatusCode, "Status code should be 500")
}

func TestRedirectionOK(t *testing.T) {
	s := setup()

	parsedUrl, _ := url.Parse("http://www.google.com")
	s.linksManager.On("Resolve", mock.Anything).Return(&core.Link{
		LongUrl: *parsedUrl,
	}, nil)

	req := httptest.NewRequest("GET", "/acc/ur", nil)
	res, err := s.app.Test(req)

	require.NoError(t, err, "Error while testing the redirection route")
	defer res.Body.Close()
	require.Equal(t, 302, res.StatusCode, "Status code should be 302")
	require.Equal(t, "http://www.google.com", res.Header.Get("Location"))
}
