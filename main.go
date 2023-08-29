package main

import (
	"github.com/ExcitingFrog/go-core-common/log"
	"github.com/ExcitingFrog/go-core-common/mongodb"
	"github.com/ExcitingFrog/go-core-common/pprof"
	"github.com/ExcitingFrog/go-core-common/provider"
	"github.com/ExcitingFrog/go-core-common/utrace"
	"github.com/ExcitingFrog/xuanwu/internal/server"
)

func main() {
	stack := provider.NewProviders()

	// init log
	logProvider := log.NewLog(nil)
	stack.AddProvider(logProvider)

	// init mongodb
	mongodbProvider := mongodb.NewMongoDB(nil)
	stack.AddProvider(mongodbProvider)

	// init pprof
	pprofProvider := pprof.NewPprof(nil)
	stack.AddProvider(pprofProvider)

	// init jaeger
	// jaegerProvider := jaeger.NewJaeger(nil)
	// stack.AddProvider(jaegerProvider)

	// init utrace
	traceProvider := utrace.NewUTrace(nil)
	stack.AddProvider(traceProvider)

	// init service
	serverProvider := server.NewServer(mongodbProvider)
	stack.AddProvider(serverProvider)

	stack.Run()
}
