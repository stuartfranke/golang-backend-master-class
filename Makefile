postgres:
	docker container run --name=postgres --env=POSTGRES_USER="${POSTGRES_USER}" --env=POSTGRES_PASS="${POSTGRES_PASS}" --env=PGDATA=/var/lib/postgresql/data --volume="${HOME}"/Projects/go/golang-backend-master-class/postgresql/data:/var/lib/postgresql/data:rw --volume=/var/lib/postgresql/data -p 5432:5432 --restart=always --runtime=runc -d postgres:15-bullseye

createdb:
	docker container exec -it postgres createdb --username="${POSTGRES_USER}" --owner="${POSTGRES_USER}" simple_bank

dropdb:
	docker container exec -it postgres dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASS}@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASS}@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test