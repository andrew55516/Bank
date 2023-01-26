postgres:
	docker run --name ps_bank -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=pwd123 -d postgres:12-alpine

createdb:
	winpty docker exec -it ps_bank createdb --username=root --owner=root bank

dropdb:
	winpty docker exec -it ps_bank dropdb bank

migrateup:
	migrate -path db/migration -database "postgresql://root:pwd123@localhost:5432/bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:pwd123@localhost:5432/bank?sslmode=disable" -verbose down

sqlc:
	docker run --rm -v "%cd%:/src" -w /src kjconroy/sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb migrateup sqlc