OPENAPI_VERSION := 5.2.0

.PHONY: clean
clean: | clean-go

.PHONY: clean-go
clean-go:
	@rm -rf ./client/*.go \
					./client/*.md \
					./client/.openapi-generator \
					./client/api \
					./client/docs \
					./client/go.*

.PHONY: generate-go-sdk
generate-go-sdk: | clean
	@docker run --rm \
		-v `pwd`:/workspace \
		openapitools/openapi-generator-cli:v${OPENAPI_VERSION} \
		generate -i /workspace/openapi/specs.yaml -g go -o /workspace/client -c /workspace/openapi/config.yaml --git-repo-id cli_stocks/client --git-user-id vijaynallagatla
