pullPostgresImg:
	docker pull postgres:14-alpine
postgres:
	docker run --name postgres14 -p 5432:5432 -e POSTGRES_USER=root -e  POSTGRES_PASSWORD=secret -d postgres:14-alpine
createdb:
	docker exec -it postgres14 createdb --username=root --owner=root mini_aspire_dev
dropdb:
	docker exec -it postgres14 dropdb  mini_aspire_dev
migrateUp:
	migrate -path ./SQL/migrations  -database "postgresql://root:secret@localhost:5432/mini_aspire_dev?sslmode=disable" -verbose up
migrateDown:
	migrate -path ./SQL/migrations -database "postgresql://root:secret@localhost:5432/mini_aspire_dev?sslmode=disable" -verbose down
sqlcGenerate:
	sqlc generate -f ./sqlc.yaml
runGoTest:
	go test -v -cover ./...
runServer:
	go run main.go
startDockerDb:
	docker start postgres14
stopDockerDb:
	docker stop postgres14
proto:
	rm -f ma/*.go
	protoc --proto_path=proto --go_out=ma --go_opt=paths=source_relative \
	--go-grpc_out=ma --go-grpc_opt=paths=source_relative \
	proto/*.proto
evans:
	evans --host localhost --port 15060 -r repl
server:
	go run main.go

mock:
	 mockgen -package mockdb -destination db/mock/store.go  github.com/KhanSufiyanMirza/mini-aspire-API/db Store

.PHONY: pullPostgresImg postgres createdb dropdb migrateUp migrateDown sqlcGenerate runGoTest proto server mock
