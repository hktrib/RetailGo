
ifneq (,$(wildcard .env.local))
    include .env.local
    export $(shell sed 's/=.*//' .env.local)
endif

all: run

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

run: $(eval SHELL:=/bin/bash)
	go run cmd/main.go


.PHONY: devpostres devcreatedb devdropdb startredis stopredis test testv