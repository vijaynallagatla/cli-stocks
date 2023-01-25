.PHONY: tools
tools:
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest


.PHONY: gen-client
gen-client: tools
	oapi-codegen -package='client' ./client/openapi.yaml > client/yfinance.go

