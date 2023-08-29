package repository

import (
	"context"

	"github.com/ExcitingFrog/go-core-common/utrace"
	"github.com/ExcitingFrog/xuanwu/internal/schema"
)

type IHello interface {
	SaveHello(ctx context.Context, h *schema.Hello) error
}

const helloCollection = "hello"
const testDB = "test"

func (r *repository) SaveHello(ctx context.Context, h *schema.Hello) error {
	// ctx, span, logger := jaeger.StartSpanAndLogFromContext(ctx, "Repository:Hello")
	// defer span.End()
	ctx, span := utrace.StartTrace(ctx, "Repository:Hello")
	defer span.End()

	_, err := r.mongo.Client.Database(testDB).Collection(helloCollection).InsertOne(ctx, h)
	if err != nil {
		// logger.Error(err.Error())
		return err
	}
	return nil
}
