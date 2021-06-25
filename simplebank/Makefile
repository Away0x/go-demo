# Start postgres container
postgres:
	docker run --name postgres12 --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

# Start mysql container
mysql:
	docker run --name mysql8 -p 3306:3306  -e MYSQL_ROOT_PASSWORD=secret -d mysql:8

# Create simple_bank database:
createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

# Drop simple_bank database:
dropdb:
	docker exec -it postgres12 dropdb simple_bank

# Create a new db migration. eg. "make migration name=init_schema"
migration:
	migrate create -ext sql -dir db/migration -seq $(name)

# Run db migration up all versions
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

# Run db migration up 1 version
migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

# Run db migration down all versions
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

# Run db migration down 1 version
migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

# Generate SQL CRUD with sqlc
sqlc:
	sqlc generate

# Run test
test:
	go test -v -cover ./...

# Run server
server:
	go run main.go

# Generate DB mock with gomock
mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/techschool/simplebank/db/sqlc Store

.PHONY: postgres createdb dropdb makemigration migrateup migratedown migrateup1 migratedown1 sqlc test server mock