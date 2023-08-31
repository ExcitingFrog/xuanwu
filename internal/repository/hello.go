package repository

import (
	"context"

	"github.com/ExcitingFrog/go-core-common/utrace"
	"github.com/ExcitingFrog/xuanwu/internal/schema"
	"go.opentelemetry.io/otel/codes"
)

type IHello interface {
	SaveHello(ctx context.Context, h *schema.Hello) error
}

const helloCollection = "hello"
const testDB = "test"

func (r *repository) SaveHello(ctx context.Context, h *schema.Hello) error {
	ctx, span, logger := utrace.StartSpanAndLogFromContext(ctx, "Repository:Hello")
	defer span.End()

	_, err := r.mongo.Client.Database(testDB).Collection(helloCollection).InsertOne(ctx, h)
	if err != nil {
		logger.Error(err.Error())
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return err
	}
	return nil
}
