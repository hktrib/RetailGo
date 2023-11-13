devpostres:
	docker run --name rg-postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15.4

devcreatedb:
	docker exec -it rg-postgres15 createdb --username=root --owner=root retail_go

devdropdb:
	docker exec -it rg-postgres15 dropdb retail_go

startredis:
	docker-compose up -d

stopredis:
	docker-compose down

test:
	go test ./tests

testv:
	go test -test.v ./tests

.PHONY: devpostres devcreatedb devdropdb startredis stopredis test testv