package server

import (
	"github.com/ExcitingFrog/xuanwu/internal/controllers"
	"github.com/ExcitingFrog/xuanwu/internal/services"
	"github.com/ExcitingFrog/xuanwu/swagger/gen/server/operations"
)

type Router struct {
	swaggerAPI  *operations.XuanWuServiceAPI
	Controllers *controllers.Controllers
}

func NewRouter(swaggerAPI *operations.XuanWuServiceAPI, xuanWuService services.IService) *Router {
	router := new(Router)
	router.swaggerAPI = swaggerAPI
	router.Controllers = controllers.NewControllers(xuanWuService)

	return router
}

func (router *Router) RegisterRoutes() {
	router.swaggerAPI.HelloHandler = operations.HelloHandlerFunc(router.Controllers.Hello)
}
