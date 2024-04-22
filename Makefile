
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
	@./scripts/proto.sh blogging internal/blogging/infra/genproto
	@./scripts/proto.sh iam internal/blogging/infra/genproto

generate-openapi:
	@./scripts/openapi.sh blogging internal/blogging/infra/genoapi blogging server
	@./scripts/openapi.sh iam internal/common/adapters/genoapi/iam iam client

lint:
	@./scripts/lint.sh blogging
	@./scripts/lint.sh common

.PHONY: migrate
migrate-up:
	cd sql && go run . migrate up
