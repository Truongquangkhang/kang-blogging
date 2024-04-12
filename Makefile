
.PHONY: generate
generate: generate-proto generate-openapi

generate-proto:
	@./scripts/proto.sh iam internal/blogging/infra/genproto

generate-openapi:
	@./scripts/openapi.sh iam internal/voucher-hub/infra/genoapi voucherhub server
	@./scripts/openapi.sh iam internal/common/adapters/genoapi/iam iam client

.PHONY: migrate
migrate-up:
	cd sql && go run . migrate up
