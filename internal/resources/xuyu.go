package resources

import (
	"context"

	"github.com/ExcitingFrog/go-core-common/log"
	"github.com/ExcitingFrog/go-core-common/provider"
	"github.com/ExcitingFrog/go-core-common/utrace"
	pb "github.com/ExcitingFrog/go-proto-lib/grpc/xuyu/proto/gen/go/proto/api"
	"github.com/ExcitingFrog/xuanwu/configs"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Xuyu struct {
	provider.IProvider

	hello pb.HelloAPIClient
}

func NewXuyu() (*Xuyu, error) {
	xuyu := &Xuyu{}
	unaryInterceptors := []grpc.UnaryClientInterceptor{
		otelgrpc.UnaryClientInterceptor(),
	}
	streamInterceptors := []grpc.StreamClientInterceptor{
		otelgrpc.StreamClientInterceptor(),
	}

	conn, err := grpc.Dial(
		configs.GetConfig().XuyuHost,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(unaryInterceptors...),
		grpc.WithChainStreamInterceptor(streamInterceptors...),
	)
	if err != nil {
		log.Logger().Error(err.Error())
		return nil, err
	}

	xuyu.hello = pb.NewHelloAPIClient(conn)

	return xuyu, nil
}

func (x *Xuyu) Hello(ctx context.Context) error {
	ctx, span, logger := utrace.StartSpanAndLogFromContext(ctx, "Resources:Hello")
	defer span.End()

	_, err := x.hello.Hello(ctx, &pb.HelloRequest{})
	if err != nil {
		logger.Error(err.Error())
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return err
	}
	return nil
}
