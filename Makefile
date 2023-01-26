OPENAPI_VERSION := 5.2.0

.PHONY: clean
clean: | clean-go

.PHONY: clean-go
clean-go:
	@rm -rf ./yapi/*.go \
					./yapi/*.md \
					./yapi/.openapi-generator \
					./yapi/api \
					./yapi/docs \
					./yapi/go.*

.PHONY: generate-go-sdk
generate-go-sdk: | clean
	@docker run --rm \
		-v `pwd`:/workspace \
		openapitools/openapi-generator-cli:v${OPENAPI_VERSION} \
		generate -i /workspace/openapi/specs.yaml -g go -o /workspace/yapi -c /workspace/openapi/config.yaml --git-repo-id cli-stocks/yapi --git-user-id vijaynallagatla

.PHONY: build
build:
	CGO_ENABLED=0 go build -v -ldflags='-s -w' -trimpath ./cmd/cstock

.PHONY: install
install: build
	sudo mv cstock /usr/local/bin

