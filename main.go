package main

import (
	"github.com/ExcitingFrog/xuanwu/internal/server"
	"github.com/ExcitingFrog/xuanwu/pkg/jaeger"
	"github.com/ExcitingFrog/xuanwu/pkg/mongodb"
	"github.com/ExcitingFrog/xuanwu/pkg/pprof"
	"github.com/ExcitingFrog/xuanwu/pkg/provider"
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
