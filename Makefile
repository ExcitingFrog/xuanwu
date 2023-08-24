.PHONY: help
help:
	@echo 'Management commands for this service:'
	@echo
	@echo 'Usage:'
	@echo 'make swagger             Init API swagger(install swagger fisrt [go swagger](https://goswagger.io/install.html))'
	@echo 'make mocks             	Generate test code(install swagger fisrt [go mockery](https://vektra.github.io/mockery/latest/installation/#github-release))'
	@echo 'make lint             	Lint code(install lint fisrt [go lint](https://golangci-lint.run/usage/install/))'
	@echo

.PHONY: swagger
swagger: swagger-clean swagger-generate-server swagger-generate-client

.PHONY: swagger-clean
swagger-clean:
	rm -rf swagger/gen
	rm -rf swagger/client
	mkdir -p swagger/gen
	mkdir -p swagger/client

.PHONY: swagger-generate-server
swagger-generate-server:
	swagger generate server -t swagger/gen -f swagger/application.yaml -s server --exclude-main

.PHONY: swagger-generate-client
swagger-generate-client:
	swagger generate client -t swagger/client -f swagger/application.yaml

.PHONY: mocks
mocks:
	mockery --dir internal --all --output internal/mocks

.PHONY: build
build:
	CGO_ENABLED=0 go build -o main -a -ldflags '-s -w' ./main.go

.PHONY: lint
lint:
	golangci-lint run
