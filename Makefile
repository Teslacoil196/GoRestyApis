postgres:
	docker run --name postgres-trixie -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:17 

createdb:
	docker exec -it postgres-trixie createdb --username=root --owner=root TeslaBank

migrate-up:
	migrate -path db/migrate -database "postgres://root:root@localhost:5432/TeslaBank?sslmode=disable" -verbose up

migrate-down:
	migrate -path db/migrate -database "postgres://root:root@localhost:5432/TeslaBank?sslmode=disable" -verbose down

dropdb:
	docker exec -it postgres-trixie dropdb TeslaBank

sqlc-gen :
	sqlc generate

test :
	go test -v -cover ./...

server :
	go run main.go

.PHONY: createdb dropdb postgres migrate-up migrate-down sqlc-gen test server