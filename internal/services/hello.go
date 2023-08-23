package services

import (
	"context"

	"github.com/ExcitingFrog/go-core-common/jaeger"
	"github.com/ExcitingFrog/xuanwu/internal/schema"
	uuid "github.com/satori/go.uuid"
)

type IHello interface {
	Hello(ctx context.Context) error
	HelloTrace(ctx context.Context) error
}

func (s *Service) Hello(ctx context.Context) error {
	ctx, span := jaeger.StartSpanFromContext(ctx, "Service:Hello")
	defer span.End()

	err := s.repository.SaveHello(ctx, &schema.Hello{
		ID: uuid.NewV4().String(),
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) HelloTrace(ctx context.Context) error {
	ctx, span := jaeger.StartSpanFromContext(ctx, "Service:HelloTrace")
	defer span.End()

	err := s.repository.SaveHello(ctx, &schema.Hello{
		ID: uuid.NewV4().String(),
	})
	if err != nil {
		return err
	}

	err = s.xuyu.Hello(ctx)
	if err != nil {
		return err
	}

	return nil
}
