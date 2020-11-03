PACKAGES := $(shell go list ./... | grep -v '/simulation')
VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')
SIMAPP = ./simapp

export GO111MODULE=on

BUILD_TAGS := netgo
BUILD_TAGS := $(strip ${BUILD_TAGS})

LD_FLAGS := -s -w \
    -X github.com/sentinel-official/hub/version.Name=sentinel-hub \
    -X github.com/sentinel-official/hub/version.ServerName=sentinel-hubd \
    -X github.com/sentinel-official/hub/version.ClientName=sentinel-hubcli \
	-X github.com/sentinel-official/hub/version.Version=${VERSION} \
	-X github.com/sentinel-official/hub/version.Commit=${COMMIT} \
	-X github.com/sentinel-official/hub/version.BuildTags=${BUILD_TAGS}
BUILD_FLAGS := -tags "${BUILD_TAGS}" -ldflags "${LD_FLAGS}"

all: install test

build: dep_verify
ifeq (${OS},Windows_NT)
	go build -mod=readonly ${BUILD_FLAGS} -o bin/sentinel-hubd.exe cmd/sentinel-hubd/main.go
	go build -mod=readonly ${BUILD_FLAGS} -o bin/sentinel-hubcli.exe cmd/sentinel-hubcli/main.go
else
	go build -mod=readonly ${BUILD_FLAGS} -o bin/sentinel-hubd cmd/sentinel-hubd/main.go
	go build -mod=readonly ${BUILD_FLAGS} -o bin/sentinel-hubcli cmd/sentinel-hubcli/main.go
endif

install: dep_verify
	go install -mod=readonly ${BUILD_FLAGS} ./cmd/sentinel-hubd
	go install -mod=readonly ${BUILD_FLAGS} ./cmd/sentinel-hubcli

test:
	@go test -mod=readonly -cover ${PACKAGES}

SIM_NUM_BLOCKS ?= 500
SIM_BLOCK_SIZE ?= 100
SIM_COMMIT ?= true

test_sim_hub_fast:
	@echo "Running hub simulation for numBlocks=100, blockSize=200. This may take awhile!"
	@go test -mod=readonly $(SIMAPP) -run TestFullAppSimulation -Genesis=${HOME}/.sentinel-hubd/config/genesis.json \
	    -Enabled=true -NumBlocks=100 -BlockSize=200 -Commit=true -Seed=99  -v -Period=5


test_sim_benchmark:
	@echo "Running hub benchmark for numBlocks=$(SIM_NUM_BLOCKS), blockSize=$(SIM_BLOCK_SIZE). This may take awhile!"
	@go test -mod=readonly -benchmem -run=^$$  $(SIMAPP) -bench=BenchmarkFullAppSimulation  \
		-Enabled=true -NumBlocks=$(SIM_NUM_BLOCKS) -BlockSize=$(SIM_BLOCK_SIZE) -Commit=$(SIM_COMMIT) -timeout 24h -Seed 15

benchmark:
	@go test -mod=readonly -bench=. ${PACKAGES}

dep_verify:
	@echo "--> Ensure dependencies have not been modified"
	@go mod verify

build-docker:
	@docker-compose build --pull

build-docker-no-cache:
	@docker-compose build --pull --no-cache

up:
	@sudo sysctl -w fs.file-max=10000000
	@docker-compose up


up-d:
	@sudo sysctl -w fs.file-max=10000000
	@docker-compose up -d

down:
	@docker-compose down

up-prod: network
	@sudo sysctl -w fs.file-max=10000000
	@docker-compose -f ./docker-compose.yml -f ./docker-compose.prod.yml up

up-prod-d: network
	@sudo sysctl -w fs.file-max=10000000
	@docker-compose -f ./docker-compose.yml -f ./docker-compose.prod.yml up -d

down-prod:
	@docker-compose -f ./docker-compose.yml -f ./docker-compose.prod.yml down --remove-orphans

shell:
	@docker-compose exec sentinel-hub bash

reset:
	@docker-compose exec sentinel-hub sentinel-hubd unsafe-reset-all

check-pub-key:
	@docker-compose exec sentinel-hub sentinel-hubd tendermint show-validator

create-account:
	@docker-compose exec sentinel-hub entrypoint create-account

create-validator:
	@docker-compose exec sentinel-hub entrypoint create-validator

add-genesis-account:
	@docker-compose exec sentinel-hub entrypoint add-genesis-account

create-offline-genesis-transaction:
	@docker-compose exec sentinel-hub entrypoint create-offline-genesis-transaction

.PHONY: all build install test benchmark dep_verify test_sim_hub_fast test_sim_benchmark
