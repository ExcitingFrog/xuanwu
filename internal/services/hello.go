package services

import (
	"context"

	"github.com/ExcitingFrog/xuanwu/internal/schema"
	uuid "github.com/satori/go.uuid"
)

type IHello interface {
	Hello(ctx context.Context) error
}

func (s *Service) Hello(ctx context.Context) error {
	s.repository.SaveHello(ctx, &schema.Hello{
		ID: uuid.NewV4().String(),
	})
	return nil
}
