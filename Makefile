all: sdk

sdk: pkg/valgo
	openapi-generator-cli generate -i https://api.val.town/openapi.json \
		-g go \
		--git-repo-id valgo --git-user-id 404wolf \
		--additional-properties=packageName=valgo,packageVersion=v1.0.0,moduleName=valgo

pkg/valgo:
	mkdir -p pkg/valgo

.PHONY: all
