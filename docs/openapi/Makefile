# docker-compose関連のコマンドを一括実行するために用意したMakefile
# runとupでは実行時に返却される終了ステータスが変わるので注意
# upの場合：起動に成功したら0を返す。コンテナ内のタスクが失敗しても0を返すので注意が必要
# runの場合：コンテナ内のタスクが成功したら0を返す。

# フロー処理
generate-api-document:
	make validate-openapi-document && \
	make generate-openapi-document

generate-api-skeleton: generate-api-document
	make generate-openapi-skeleton

generate:
	make validate-openapi-document && \
	make generate-openapi-document && \
	make generate-openapi-skeleton

build: generate-api-document
	make build-openapi-mock && \
	make build-openapi-document

# タスク処理
validate-openapi-document:
	docker-compose run --rm validate-openapi-document

generate-openapi-document:
	make clean-api-document
	docker-compose run --rm generate-openapi-document

generate-openapi-skeleton:
	make clean-api-skeleton
	docker-compose run --rm generate-openapi-skeleton

clean-api-document:
	find ./dist -depth 1 -not -name ".gitignore" | xargs rm -rfv

clean-api-skeleton:
	find ./skeleton -depth 1 -not -name ".gitignore" | xargs rm -rfv

build-openapi-mock:
	docker-compose build --no-cache --force-rm openapi-mock

build-openapi-document:
	docker-compose build --no-cache --force-rm openapi-document

