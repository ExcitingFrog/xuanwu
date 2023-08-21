package main

import (
	"github.com/ExcitingFrog/go-core-common/jaeger"
	"github.com/ExcitingFrog/go-core-common/mongodb"
	"github.com/ExcitingFrog/go-core-common/pprof"
	"github.com/ExcitingFrog/go-core-common/provider"
	"github.com/ExcitingFrog/xuanwu/internal/server"
)

func main() {
	stack := provider.NewProviders()

	// init mongodb
	mongodbProvider := mongodb.NewMongoDB(nil)
	stack.AddProvider(mongodbProvider)

	// init opentelemetry

	// init pprof
	pprofProvider := pprof.NewPprof(nil)
	stack.AddProvider(pprofProvider)

	// init jaeger
	jaegerProvider := jaeger.NewJaeger(nil)
	stack.AddProvider(jaegerProvider)

	// init service
	serverProvider := server.NewServer(mongodbProvider, jaegerProvider)
	stack.AddProvider(serverProvider)

	stack.Run()
}
