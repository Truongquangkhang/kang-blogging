.PHONY: tools

LOCAL_DIR=/usr/local
BIN_DIR=$(LOCAL_DIR)/bin

PB_REL=https://github.com/protocolbuffers/protobuf/releases
PB_VERSION=21.2
#PB_ZIP=protoc-$(PB_VERSION)-osx-universal_binary.zip
PB_ZIP=protoc-$(PB_VERSION)-linux-x86_64.zip

BUF_REL=https://github.com/bufbuild/buf/releases
BUF_VERSION=1.6.0
#BUF_BIN=buf-Darwin-x86_64
BUF_BIN=buf-Linux-x86_64

c4:
	@./scripts/c4.sh tools/c4

prepare: prepare-protoc prepare-go-tools

prepare-protoc:
	apt-get update
	apt-get install zip unzip
	curl -sSLO $(PB_REL)/download/v$(PB_VERSION)/$(PB_ZIP)
	unzip $(PB_ZIP) -d $(LOCAL_DIR) -x "readme.txt"
	chmod +x $(BIN_DIR)/protoc
	rm -f $(PB_ZIP)

prepare-buf:
	curl -sSL $(BUF_REL)/download/v$(BUF_VERSION)/$(BUF_BIN) -o $(BIN_DIR)/buf
	chmod +x $(BIN_DIR)/buf

prepare-go-tools:
	# codegen
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
	# migration, uncomment to install for local development
	#go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

.PHONY: generate
generate: generate-proto generate-openapi

generate-proto:
	@./scripts/proto.sh nio-core internal/nio-core/infra/genproto
	@./scripts/proto.sh qr-search internal/voucher-hub/adapters/genproto

generate-openapi:
	@./scripts/openapi.sh voucher-hub internal/voucher-hub/infra/genoapi voucherhub server
	@./scripts/openapi.sh iam internal/common/adapters/genoapi/iam iam client
	@./scripts/openapi.sh marketing-portal internal/voucher-hub/adapters/genoapi/mktportal mktportal client
	@./scripts/openapi.sh qr-search internal/voucher-hub/adapters/genoapi/qrsearch qrsearch client

.PHONY: lint
lint:
	@./scripts/lint.sh common
	@./scripts/lint.sh voucher-hub
	@./scripts/lint.sh nio-core

.PHONY: dockerize
dockerize:
	@./scripts/dockerize.sh dev asia.gcr.io/teko-registry/nio/nio-core-sql sql nio-core false
	@./scripts/dockerize.sh dev asia.gcr.io/teko-registry/nio/nio-core-app app nio-core false
	@./scripts/dockerize.sh dev asia.gcr.io/teko-registry/nio/voucher-hub-app app voucher-hub false

.PHONY: migrate
migrate-up:
	cd sql && go run . migrate up