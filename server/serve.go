package server

import (
	"context"
	"fmt"
	"io/fs"
	"log"
	"log/slog"
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/httplog/v2"

	oapimiddleware "github.com/oapi-codegen/nethttp-middleware"
)

func Start(port int, frontendFS fs.FS, configPath string) error {
	address := fmt.Sprintf("0.0.0.0:%d", port)

	// Create middlewares
	loggerMiddleware := createLoggerMiddleware()
	corsMiddleware := cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	})
	openapiValidationMiddleware, err := createOapiValidationMiddleware()
	if err != nil {
		return fmt.Errorf("failed to create openapi validation middleware %w", err)
	}

	// Create base router
	router := chi.NewRouter()
	router.Use(loggerMiddleware, corsMiddleware)

	// Register routes for frontend
	fileServer := http.FileServer(http.FS(frontendFS))
	router.Handle("/*", http.StripPrefix("/", fileServer))

	// Create api router and mount
	apiRouter := chi.NewRouter()
	apiRouter.Use(openapiValidationMiddleware)
	apiHandler := HandlerFromMux(NewServer(configPath), apiRouter)
	router.Mount("/api", apiHandler)

	// Start listening
	log.Printf("Server listening on %s\n", address)
	err = http.ListenAndServe(address, router)
	if err != nil {
		return err
	}

	return nil
}

func createLoggerMiddleware() func(http.Handler) http.Handler {
	// This middleware includes the recover middleware
	return httplog.RequestLogger(
		httplog.NewLogger(
			"twitchets",
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
	spec.Servers = openapi3.Servers{&openapi3.Server{URL: "/api"}}

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
