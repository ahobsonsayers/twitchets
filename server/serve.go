package server

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog/v2"

	oapimiddleware "github.com/oapi-codegen/nethttp-middleware"
)

func Start(port int, configPath string) error {
	address := fmt.Sprintf("0.0.0.0:%d", port)

	loggerMiddleware := createLoggerMiddleware()
	openapiValidationMiddleware, err := createOapiValidationMiddleware()
	if err != nil {
		return fmt.Errorf("failed to create openapi validation middleware %w", err)
	}

	// Create router (with middleware)
	router := chi.NewRouter()
	router.Use(
		loggerMiddleware,
		openapiValidationMiddleware,
	)

	// Register server routes on chi router
	server := NewServer(configPath)
	handler := HandlerFromMux(server, router)

	// Start listening
	log.Printf("Server listening on %s\n", address)
	err = http.ListenAndServe(address, handler)
	if err != nil {
		return err
	}

	return nil
}

func createLoggerMiddleware() func(http.Handler) http.Handler {
	// This middleware includes the recover middleware
	return httplog.RequestLogger(
		httplog.NewLogger(
			"abs-tract",
			httplog.Options{
				LogLevel:       slog.LevelDebug,
				RequestHeaders: true,
				Concise:        true,
			},
		),
	)
}

func createOapiValidationMiddleware() (func(http.Handler) http.Handler, error) {
	// Load openapi spec
	spec, err := GetSwagger()
	if err != nil {
		return nil, fmt.Errorf("failed to load openapi spec: %w", err)
	}

	// Set server endpoint to the route. This ensures request validation
	// doesn't fail in the validation middleware.
	// See: https://github.com/oapi-codegen/oapi-codegen/issues/1123
	spec.Servers = openapi3.Servers{&openapi3.Server{URL: "/"}}

	return oapimiddleware.OapiRequestValidatorWithOptions(
		spec,
		&oapimiddleware.Options{
			SilenceServersWarning: true,
			Options: openapi3filter.Options{
				AuthenticationFunc: func(_ context.Context, _ *openapi3filter.AuthenticationInput) error {
					return nil // Do nothing
				},
			},
		},
	), nil
}
