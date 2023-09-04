include ./app.env

network:
	docker network create golang-backend-master-class-network

postgres:
	docker container run --name=postgres --env=POSTGRES_USER="${POSTGRES_USER}" --env=POSTGRES_PASS="${POSTGRES_PASS}" --env=PGDATA=/var/lib/postgresql/data --volume="${HOME}"/Projects/go/golang-backend-master-class/postgresql/data:/var/lib/postgresql/data:rw --volume=/var/lib/postgresql/data -p 54320:5432 --restart=always --runtime=runc -d postgres:15-bullseye

mysql:
	docker run --name mysql8 -p 3306:3306  -e MYSQL_ROOT_PASSWORD=secret -d mysql:8

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
	go test -v -cover -short ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/stuartfranke/golang-backend-master-class/db/sqlc Store
	mockgen -package mockwk -destination worker/mock/distributor.go github.com/stuartfranke/golang-backend-master-class/worker TaskDistributor

proto:
	rm -f pb/*.go
	#rm -f doc/swagger/*.swagger.json
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative --grpc-gateway_opt=logtostderr=true \
	proto/*.proto
#	--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=simple_bank \
#	proto/*.proto
	#statik -src=./doc/swagger -dest=./doc

evans:
	evans --host localhost --port 9090 -r repl

redis:
	docker run --name redis -p 6379:6379 -d redis:7-alpine

.PHONY: network postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 new_migration db_docs db_schema sqlc test server mock proto evans redis
