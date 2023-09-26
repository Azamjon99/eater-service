PWD=$(shell pwd)
SERVICE=eater-svc
MIGRATION_PATH=${PWD}/src/infrastructure/migrations
PROTOS_PATH=$(PWD)/src/infrastructure/protos

server:
	go run main.go

migration:
	migrate create -ext sql -dir pkg/database/migrations -seq $(table)

migrateup:
	migrate -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable&search_path=public" -path ./pkg/database/migrations up

migratedown:
	migrate -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable&search_path=public" -path ./pkg/database/migrations down

PWD=$(shell pwd)
PROTOS_PATH=$(PWD)/src/infrastructure/protos

clear:
	rm -rf genprotos/*
gen-eater:
	protoc \
	--go_out=./src/application/protos \
	--go_opt=paths=import \
	--go-grpc_out=./src/application/protos \
	--go-grpc_opt=paths=import \
	-I=$(PWD)/src/infrastructure/protos/eater \
	$(PWD)/src/infrastructure/protos/eater/*.proto


docker: bin-lunix
	docker build --rm -t eater-svc -f ${PWD}/deploy/Dockerfile .





server-docker: server
	docker build --rm -t server-app -f ./docker/server/Dockerfile .

deploy-server:
	docker run --rm -p 3030:3030 --name=app server-app

deploy-client:
	docker run --rm --network=host --name=app client-app

compose-up:
	docker-compose -f docker-compose/docker-compose.yml up

compose-down:
	docker-compose -f docker-compose/docker-compose.yml down	