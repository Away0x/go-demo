APP_NAME = "graphqlapp"
PROXY = "https://goproxy.cn"

.PHONY :default
default:
	go build -o ${APP_NAME}

.PHONY :install
install:
	env GOPROXY=${PROXY} go mod download

.PHONY :version
version:
	go run main.go version

.PHONY :dev
dev:
  # go install -u github.com/cosmtrek/air
	air -c air.toml

.PHONY :prod
prod:
	make
	./${APP_NAME} server -c config/production.yaml

.PHONY :api-doc
api-doc:
  # go install -u github.com/swaggo/swag/cmd/swag
	swag init -g routes/api.go

.PHONY :mock
mock:
	env SERVER_APP_MODE=test go run main.go mock

.PHONY :gqlgen
gqlgen:
	gqlgen generate

.PHONY :test-controller
test-controller:
	go test -v ./tests/controllers/...

.PHONY :test-graphql
test-graphql:
	go test -v ./tests/graphql/...

.PHONY :test-model
test-model:
	go test -v ./tests/models/...

.PHONY :test-service
test-service:
	go test -v ./tests/services/...

.PHONY :test
test:
	make test-model;
	make test-service;
	make test-controller
	make test-graphql

.PHONY :clean
clean:
	if [ -f ${APP_NAME} ]; then rm ${APP_NAME}; fi

.PHONY :help
help:
	@echo "make - compile the source code"
	@echo "make install - install dep"
	@echo "make version - print version"
	@echo "make dev - Development environment run service"
	@echo "make prod - Production environment run service"
	@echo "make mock - mock data"
	@echo "make gqlgen - generate graphql files"
	@echo "make test - test all"
	@echo "make test-controller - test controller"
	@echo "make test-graphql - test graphql"
	@echo "make test-model - test model"
	@echo "make test-service - test service"
	@echo "make doc - generate swagger api docs"
	@echo "make clean - remove binary file"
