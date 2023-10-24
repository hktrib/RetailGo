postgres:
	docker run --name rg-postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15.4

createdb:
	docker exec -it rg-postgres15 createdb --username=root --owner=root retail_go

dropdb:
	docker exec -it rg-postgres15 dropdb retail_go


.PHONY: postgres createdb dropdb