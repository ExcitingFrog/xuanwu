.PHONY: help
help:
	@echo 'Management commands for this service:'
	@echo
	@echo 'Usage:'
	@echo 'make swagger             Init API swagger(install swagger fisrt [go swagger](https://goswagger.io/install.html))'
	@echo 'make mocks             	Generate test code(install swagger fisrt [go mockery](https://vektra.github.io/mockery/latest/installation/#github-release))'
	@echo

.PHONY: swagger
swagger: swagger-clean swagger-generate

.PHONY: swagger-clean
swagger-clean:
	rm -rf swagger/gen
	mkdir -p swagger/gen

.PHONY: swagger-generate
swagger-generate:
	swagger generate server -t swagger/gen -f swagger/application.yaml -s server --exclude-main

.PHONY: mocks
mocks:
	mockery --dir internal --all --output internal/mocks

.PHONY: build
build:
	CGO_ENABLED=0 go build -o main -a -ldflags '-s -w' ./main.go
