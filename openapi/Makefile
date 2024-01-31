# NOTE: 誤操作防止のためtarget指定なしの場合はエラー扱いにする
all:
	@echo Please specify the target. >&2
	@exit 1

clean-api-client:
	@find ./artifacts/client -depth 1 -not -name ".gitignore" | xargs rm -rfv

clean-api-skeleton:
	@find ./artifacts/skeleton -depth 1 -not -name ".gitignore" | xargs rm -rfv


inspect-openapi:
	docker run \
		--rm \
		-v ${PWD}:/local/src \
		openapitools/openapi-generator-cli \
		validate \
			-i /local/src/reference/root.yaml

gen-openapi-client: clean-api-client
	docker run \
		--rm \
		-v ${PWD}:/local/src \
		openapitools/openapi-generator-cli \
		generate \
			-i /local/src/reference/root.yaml \
			-g typescript-axios\
			 -o /local/src/artifacts/client

gen-openapi-skeleton: clean-api-skeleton
	docker run \
		--rm \
		-v ${PWD}:/local/src \
		openapitools/openapi-generator-cli \
		generate \
			-i /local/src/reference/root.yaml \
			-g go-server \
			-o /local/src/artifacts/skeleton
