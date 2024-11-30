all: sdk

sdk: pkg/valgo
	openapi-generator-cli generate -i https://api.val.town/openapi.json -g go -o pkg/valgo

pkg/valgo:
	mkdir -p pkg/valgo

.PHONY: all
