include ./app.env

postgres:
	docker container run --name=postgres --env=POSTGRES_USER="${POSTGRES_USER}" --env=POSTGRES_PASS="${POSTGRES_PASS}" --env=PGDATA=/var/lib/postgresql/data --volume="${HOME}"/Projects/go/golang-backend-master-class/postgresql/data:/var/lib/postgresql/data:rw --volume=/var/lib/postgresql/data -p 54320:5432 --restart=always --runtime=runc -d postgres:15-bullseye

createdb:
	docker container exec -it postgres createdb --username="${POSTGRES_USER}" --owner="${POSTGRES_USER}" simple_bank

dropdb:
	docker container exec -it postgres dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASS}@localhost:54320/simple_bank?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASS}@localhost:54320/simple_bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASS}@localhost:54320/simple_bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASS}@localhost:54320/simple_bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/stuartfranke/golang-backend-master-class/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc test server mock