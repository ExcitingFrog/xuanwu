package services

import (
	"context"

	"github.com/ExcitingFrog/xuanwu/internal/schema"
	"github.com/ExcitingFrog/xuanwu/pkg/jaeger"
	uuid "github.com/satori/go.uuid"
)

type IHello interface {
	Hello(ctx context.Context) error
}

func (s *Service) Hello(ctx context.Context) error {
	ctx, span := jaeger.StartSpanFromContext(ctx, "Service:Hello")
	defer span.End()

	s.repository.SaveHello(ctx, &schema.Hello{
		ID: uuid.NewV4().String(),
	})
	return nil
}
