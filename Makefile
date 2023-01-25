OPENAPI_VERSION := 5.2.0

.PHONY: clean
clean: | clean-go

.PHONY: clean-go
clean-go:
	@rm -rf ./gen/client/*.go \
					./gen/client/*.md \
					./gen/client/.openapi-generator \
					./gen/client/api \
					./gen/client/docs \
					./gen/client/go.*

.PHONY: generate-go-sdk
generate-go-sdk: | clean
	@docker run --rm \
		-v `pwd`:/workspace \
		openapitools/openapi-generator-cli:v${OPENAPI_VERSION} \
		generate -i /workspace/openapi/specs.yaml -g go -o /workspace/gen/client -c /workspace/openapi/config.yaml --git-repo-id cli_stocks/gen/client --git-user-id vijaynallagatla
