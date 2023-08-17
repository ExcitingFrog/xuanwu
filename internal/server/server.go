package server

import (
	"log"

	"github.com/ExcitingFrog/xuanwu/configs"
	"github.com/ExcitingFrog/xuanwu/internal/repository"
	"github.com/ExcitingFrog/xuanwu/internal/services"
	"github.com/ExcitingFrog/xuanwu/pkg/mongodb"
	"github.com/ExcitingFrog/xuanwu/pkg/provider"
	"github.com/ExcitingFrog/xuanwu/swagger/gen/server"
	"github.com/ExcitingFrog/xuanwu/swagger/gen/server/operations"
	"github.com/go-openapi/loads"
)

type Server struct {
	provider.IProvider
	server  *server.Server
	mongodb *mongodb.MongoDB
}

func NewServer(mongodb *mongodb.MongoDB) *Server {
	return &Server{
		mongodb: mongodb,
	}
}

func (s *Server) Close() error {
	if err := s.server.Shutdown(); err != nil {
		return err
	}
	return nil
}

func (s *Server) Run() error {
	swaggerSpec, err := loads.Embedded(server.SwaggerJSON, server.FlatSwaggerJSON)
	if err != nil {
		return err
	}

	api := operations.NewXuanWuServiceAPI(swaggerSpec)

	repository := repository.NewRepository(s.mongodb)

	xuanwuServices := services.NewService(repository)
	router := NewRouter(api, xuanwuServices)
	router.RegisterRoutes()

	server := server.NewServer(api)
	server.Port = configs.GetConfig().Port

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

	s.server = server
	return nil
}
