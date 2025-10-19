package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/ahobsonsayers/twitchets/config"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	configPath string
}

var _ StrictServerInterface = Server{}

func (s Server) GetConfig(ctx context.Context, request GetConfigRequestObject) (GetConfigResponseObject, error) {
	conf, err := config.Load(s.configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load: %w", err)
	}

	return GetConfig200JSONResponse(conf), nil
}

func (s Server) PutConfig(ctx context.Context, request PutConfigRequestObject) (PutConfigResponseObject, error) {
	conf := config.Config(*request.Body)
	err := config.Save(conf, s.configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to save: %w", err)
	}

	return PutConfig200Response{}, nil
}

func NewServer(configPath string) ServerInterface {
	server := Server{
		configPath: configPath,
	}
	return NewStrictHandler(server, nil)
}

func Serve(port int, configPath string) error {
	address := fmt.Sprintf("0.0.0.0:%d", port)

	router := chi.NewMux()
	server := NewServer(configPath)
	handler := HandlerFromMux(server, router)

	// Start the Server
	log.Printf("Server listening on %s\n", address)
	err := http.ListenAndServe(address, handler)
	if err != nil {
		return err
	}

	return nil
}
