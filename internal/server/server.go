package server

import (
	"github.com/ExcitingFrog/go-core-common/log"
	"github.com/ExcitingFrog/go-core-common/mongodb"
	"github.com/ExcitingFrog/go-core-common/provider"
	"github.com/ExcitingFrog/xuanwu/configs"
	"github.com/ExcitingFrog/xuanwu/internal/repository"
	"github.com/ExcitingFrog/xuanwu/internal/resources"
	"github.com/ExcitingFrog/xuanwu/internal/services"
	"github.com/ExcitingFrog/xuanwu/swagger/gen/server"
	"github.com/ExcitingFrog/xuanwu/swagger/gen/server/operations"
	"github.com/go-openapi/loads"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
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

func (s *Server) Init() error {
	return nil
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

	xuyu, err := resources.NewXuyu()
	if err != nil {
		log.Logger().Error(err.Error())
		return err
	}

	xuanwuServices := services.NewService(repository, xuyu)
	router := NewRouter(api, xuanwuServices)

	router.RegisterRoutes()

	server := server.NewServer(api)
	server.Port = configs.GetConfig().Port

	server.SetHandler(otelhttp.NewHandler(api.Serve(nil), "Middleware:Xuanwu"))

	if err := server.Serve(); err != nil {
		log.Logger().Error(err.Error())
		return err
	}

	s.server = server
	return nil
}
