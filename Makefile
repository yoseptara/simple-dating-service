postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_dating_service

dropdb:
	docker exec -it postgres12 dropdb simple_dating_service

migrateup:
	source "./.env" && migrate -path db/migration -database "postgresql://$${POSTGRES_USERNAME}:$${POSTGRES_PASS}@$${POSTGRES_HOST}:$${POSTGRES_PORT}/$${POSTGRES_DB_NAME}?sslmode=disable" -verbose up

migrateup1:
	source "./.env" && migrate -path db/migration -database "postgresql://$${POSTGRES_USERNAME}:$${POSTGRES_PASS}@$${POSTGRES_HOST}:$${POSTGRES_PORT}/$${POSTGRES_DB_NAME}?sslmode=disable" -verbose up 1

migratedown:
	source "./.env" && migrate -path db/migration -database "postgresql://$${POSTGRES_USERNAME}:$${POSTGRES_PASS}@$${POSTGRES_HOST}:$${POSTGRES_PORT}/$${POSTGRES_DB_NAME}?sslmode=disable" -verbose down

migratedown1:
	source "./.env" && migrate -path db/migration -database "postgresql://$${POSTGRES_USERNAME}:$${POSTGRES_PASS}@$${POSTGRES_HOST}:$${POSTGRES_PORT}/$${POSTGRES_DB_NAME}?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go tara/simplebank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc test server mock