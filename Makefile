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
check_swagger_install:
	which swagger || GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger
swagger: check_swagger_install
    GO111MODULE=off swagger generate spec -o ./swagger.yaml --scan-models
runServer:
	go run main.go
startDockerDb:
	docker start postgres14
stopDockerDb:
	docker stop postgres14
.PHONY: pullPostgresImg postgres createdb dropdb migrateUp migrateDown sqlcGenerate runGoTest
