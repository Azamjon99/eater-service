# PWD=$(shell pwd)
# SERVICE=eater-svc
# MIGRATION_PATH=${PWD}/src/infrastructure/migrations
# migration-file:
# docker run -v ${MIGRATION_PATH}:/migrations migrate/migrate create -ext sql -dir /migrations -seq $(NAME)
# clear:
# rm -rf ${PWD}/bin/${SERVICE}
# bin: clear
# CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -installsuffix cgo -o ${PWD}/bin/${SERVICE} ${PWD}/main.go
# chmod +x ${PWD}/bin/${SERVICE}
# bin-linux: clear
# CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o ${PWD}/bin/${SERVICE} ${PWD}/main.go
# chmod +x ${PWD}/bin/${SERVICE}
# bin-windows: clear
# CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -a -installsuffix cgo -o ${PWD}/bin/${SERVICE} ${PWD}/main.go
# chmod +x ${PWD}/bin/${SERVICE}

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