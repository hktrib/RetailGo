postgres:
	docker run --name rg-postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15.4

createdb:
	docker exec -it rg-postgres15 createdb --username=root --owner=root retail_go

dropdb:
	docker exec -it rg-postgres15 dropdb retail_go

dbschema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml
	mv dbml-error.log doc

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/retail_go?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/retail_go?sslmode=disable" -verbose down

sqlc:
	sqlc generate 

sqlc_delete:
	rm -rf ./db/sqlc

.PHONY: postgres createdb dropdb dbschema migrateup migratedown sqlc sqlc_delete