APP_NAME = "graphqlapp"
PROXY = "https://goproxy.cn"

default:
	go build -o ${APP_NAME}
	# env GOOS=linux GOARCH=amd64 go build -o ${APP_NAME}

install:
	env GOPROXY=${PROXY} go mod download

version:
	go run main.go version

dev:
  # go install -u github.com/cosmtrek/air
	 env GRAPHQLAPP_APP_MODE=development air -c air.toml

deploy:
	make && ./bin/deploy restart

api-doc:
  # go install -u github.com/swaggo/swag/cmd/swag
	swag init -g routes/api.go

mock:
	env GRAPHQLAPP_APP_MODE=development go run main.go mock

gqlgen:
	gqlgen generate

test-controller:
	env GRAPHQLAPP_APP_MODE=test go test -v ./tests/controllers/...

test-graphql:
	env GRAPHQLAPP_APP_MODE=test go test -v ./tests/graphql/...

test-model:
	env GRAPHQLAPP_APP_MODE=test go test -v ./tests/models/...

test-service:
	env GRAPHQLAPP_APP_MODE=test go test -v ./tests/services/...

test:
	make test-model;
	make test-service;
	make test-controller
	make test-graphql

clean:
	if [ -f ${APP_NAME} ]; then rm ${APP_NAME}; fi

help:
	@echo "make - compile the source code"
	@echo "make install - install dep"
	@echo "make version - print version"
	@echo "make dev - run go fresh"
	@echo "make deploy - deploy app"
	@echo "make mock - mock data"
	@echo "make gqlgen - generate graphql files"
	@echo "make test - test all"
	@echo "make test-controller - test controller"
	@echo "make test-graphql - test graphql"
	@echo "make test-model - test model"
	@echo "make test-service - test service"
	@echo "make doc - generate swagger api docs"
	@echo "make clean - remove binary file"
