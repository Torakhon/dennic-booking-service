-include .env
export

CURRENT_DIR=$(shell pwd)
APP=content_service
CMD_DIR=./cmd

.DEFAULT_GOAL = build

# build for current os
.PHONY: build
build:
	go build -ldflags="-s -w" -o ./bin/${APP} ${CMD_DIR}/app/main.go

# build for linux amd64
.PHONY: build-linux
build-linux:
	CGO_ENABLED=0 GOARCH="amd64" GOOS=linux go build -ldflags="-s -w" -o ./bin/${APP} ${CMD_DIR}/app/main.go

# run service
.PHONY: run
run:
	go run ${CMD_DIR}/app/main.go

# migrate
.PHONY: migrate
migrate:
	migrate -source file://migrations -database postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DATABASE}?sslmode=disable up

# proto-gen
.PHONY: proto-gen
proto-gen:
	./scripts/gen-proto.sh  ${CURRENT_DIR}
	ls genproto/*.pb.go | xargs -n1 -IX bash -c "sed -e '/bool/ s/,omitempty//' X > X.tmp && mv X{.tmp,}"

.PHONY: create-migration
create-migration:
	migrate create -dir migrations -ext sql -seq patients_table

# git submodule init 	
.PHONY: pull-proto
pull-proto:
	git submodule update --init --recursive

# go generate	
.PHONY: go-gen
go-gen:
	go generate ./...

# run test
test:
	go test -v -cover -race ./internal/...

# -------------- for deploy --------------
build-image:
	docker build --rm -t ${REGISTRY}/${PROJECT_NAME}/${APP}:${TAG} .
	docker tag ${REGISTRY}/${PROJECT_NAME}/${APP}:${TAG} ${REGISTRY}/${PROJECT_NAME}/${APP}:${ENV_TAG}

push-image:
	docker push ${REGISTRY}/${PROJECT_NAME}/${APP}:${TAG}
	docker push ${REGISTRY}/${PROJECT_NAME}/${APP}:${ENV_TAG}