package controllers

import (
	"github.com/ExcitingFrog/xuanwu/swagger/gen/models"
	"github.com/ExcitingFrog/xuanwu/swagger/gen/server/operations"
	"github.com/go-openapi/runtime/middleware"
)

func (c *Controllers) Hello(params operations.HelloParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	if err := c.service.Hello(ctx); err != nil {
		return operations.NewHelloBadRequest().WithPayload(&models.ErrorResponse{
			Code:    400,
			Message: err.Error(),
		})
	}

	return operations.NewHelloOK().WithPayload("Hello, world!")
}
