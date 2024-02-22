prev:
	#nothing

postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root bank

dropdb:
	docker exec -it postgres12 dropdb --username=root --owner=root bank

table:
	migrate create -ext sql -dir db/migration -seq [NAME]

docker-psql:
	docker exec -it postgres12 psql -U root # to open postgres12 in docker

migrate_up:
	migrate -path db/migration -database postgresql://root:secret@localhost:5432/bank?sslmode=disable -verbose up

migrate_down:
	migrate -path db/migration -database postgresql://root:secret@localhost:5432/bank?sslmode=disable -verbose down

sqlc: 
	sqlc generate
 
test:
	go test -v -cover ./...
	
.PHONY: postgres createdb dropdb table docker-psql migrate_up migrate_down sqlc test