package main

import (
	"github.com/ExcitingFrog/xuanwu/internal/server"
	"github.com/ExcitingFrog/xuanwu/pkg/mongodb"
	"github.com/ExcitingFrog/xuanwu/pkg/provider"
)

func main() {
	stack := provider.NewProviders()

	// init mongodb
	mongodbProvider := mongodb.NewMongoDB(nil)
	stack.AddProvider(mongodbProvider)

	// init opentelemetry

	// init pprof

	// init probe

	// init service
	serverProvider := server.NewServer(mongodbProvider)
	stack.AddProvider(serverProvider)

	stack.Run()
}
