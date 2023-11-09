postgres:
	docker run --name postgres --network store-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16-alpine

createdb:
	docker exec -it postgres createdb --username=root --owner=root store_db

dropdb:
	docker exec -it postgres dropdb store_db

migratecreate:
	migrate create -ext sql -dir db/migration -seq init_schema

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/store_db?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/store_db?sslmode=disable" -verbose down

reset:
	docker exec -it postgres dropdb store_db
	docker exec -it postgres createdb --username=root --owner=root store_db
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/store_db?sslmode=disable" -verbose up

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY:
	postgres
	createdb
	dropdb
	migratecreate
	migrateup
	migratedown
	reset
	sqlc
	test