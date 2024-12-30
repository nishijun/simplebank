postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres createdb --username=postgres --owner=postgres simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://postgres:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up;

dropdb:
	docker exec -it postgres dropdb --username=postgres simple_bank

migratedown:
	migrate -path db/migration -database "postgresql://postgres:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down;

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -destination db/mock/store.go -source db/sqlc/store.go -package mockdb

.PHONY: createdb dropdb postgres migrateup migratedown sqlc test server mock