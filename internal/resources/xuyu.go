package resources

import (
	"context"

	"github.com/ExcitingFrog/go-core-common/log"
	"github.com/ExcitingFrog/go-core-common/provider"
	"github.com/ExcitingFrog/go-core-common/utrace"
	"github.com/ExcitingFrog/xuanwu/configs"
	pb "github.com/ExcitingFrog/xuyu/proto/gen/go/proto/api"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
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
	// ctx, span, logger := jaeger.StartSpanAndLogFromContext(ctx, "Resources:Hello")
	// defer span.End()
	ctx, span := utrace.StartTrace(ctx, "Resources:Hello")
	defer span.End()

	_, err := x.hello.Hello(ctx, &pb.HelloRequest{})
	if err != nil {
		// logger.Error(err.Error())
		return err
	}
	return nil
}
