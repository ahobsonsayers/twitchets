package server

import (
	"context"
	"fmt"

	"github.com/ahobsonsayers/twitchets/config"
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
