package controllers

import (
	"github.com/ExcitingFrog/go-core-common/utrace"
	"github.com/ExcitingFrog/xuanwu/swagger/gen/models"
	"github.com/ExcitingFrog/xuanwu/swagger/gen/server/operations"
	"github.com/go-openapi/runtime/middleware"
)

func (c *Controllers) Hello(params operations.HelloParams) middleware.Responder {
	// ctx, span, logger := jaeger.StartSpanAndLogFromContext(params.HTTPRequest.Context(), "Controller:Hello")
	// defer span.End()
	ctx, span := utrace.StartTrace(params.HTTPRequest.Context(), "Controller:Hello")

	defer span.End()

	if err := c.service.Hello(ctx); err != nil {
		// logger.Error(err.Error())
		return operations.NewHelloBadRequest().WithPayload(&models.ErrorResponse{
			Code:    400,
			Message: err.Error(),
		})
	}

	return operations.NewHelloOK().WithPayload("Hello, world!")
}

func (c *Controllers) HelloTrace(params operations.HelloTraceParams) middleware.Responder {
	// ctx, span, logger := jaeger.StartSpanAndLogFromContext(params.HTTPRequest.Context(), "Controller:HelloTrace")
	// defer span.End()
	ctx, span := utrace.StartTrace(params.HTTPRequest.Context(), "Controller:HelloTrace")
	defer span.End()

	if err := c.service.HelloTrace(ctx); err != nil {
		// logger.Error(err.Error())
		return operations.NewHelloTraceBadRequest().WithPayload(&models.ErrorResponse{
			Code:    400,
			Message: err.Error(),
		})
	}

	return operations.NewHelloTraceOK().WithPayload("Hello, world!")
}
