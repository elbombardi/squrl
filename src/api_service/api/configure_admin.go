// This file is safe to edit. Once it exists it will not be overwritten

package api

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/elbombardi/squrl/src/api_service/api/operations"
	"github.com/elbombardi/squrl/src/api_service/api/operations/accounts"
	"github.com/elbombardi/squrl/src/api_service/api/operations/general"
	"github.com/elbombardi/squrl/src/api_service/api/operations/urls"
)

//go:generate swagger generate server --target ../../api_service --name Admin --spec ../swagger.yml --model-package api/models --server-package api --principal interface{} --exclude-main

func configureFlags(api *operations.AdminAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.AdminAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.AccountsCreateAccountHandler == nil {
		api.AccountsCreateAccountHandler = accounts.CreateAccountHandlerFunc(func(params accounts.CreateAccountParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation accounts.CreateAccount has not yet been implemented")
		})
	}
	if api.UrlsCreateURLHandler == nil {
		api.UrlsCreateURLHandler = urls.CreateURLHandlerFunc(func(params urls.CreateURLParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation urls.CreateURL has not yet been implemented")
		})
	}
	if api.GeneralHealthcheckHandler == nil {
		api.GeneralHealthcheckHandler = general.HealthcheckHandlerFunc(func(params general.HealthcheckParams) middleware.Responder {
			return middleware.NotImplemented("operation general.Healthcheck has not yet been implemented")
		})
	}
	if api.GeneralLoginHandler == nil {
		api.GeneralLoginHandler = general.LoginHandlerFunc(func(params general.LoginParams) middleware.Responder {
			return middleware.NotImplemented("operation general.Login has not yet been implemented")
		})
	}
	if api.AccountsUpdateAccountHandler == nil {
		api.AccountsUpdateAccountHandler = accounts.UpdateAccountHandlerFunc(func(params accounts.UpdateAccountParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation accounts.UpdateAccount has not yet been implemented")
		})
	}
	if api.UrlsUpdateURLHandler == nil {
		api.UrlsUpdateURLHandler = urls.UpdateURLHandlerFunc(func(params urls.UpdateURLParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation urls.UpdateURL has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
