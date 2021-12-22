
S3_BUCKET_NAME=testing-file-generator

build-api:
	docker-compose build --no-cache --force-rm testing-file-generator-api

# build-web:
#   docker-compose build --no-cache --force-rm testing-file-generator-web

# push-image:
# 	make push-api-image
# 	make push-web-image

# push-api-image:
# TODO: dockerhub にpushを実装
# TODO: github にpushを実装

# push-web-image:
# TODO: dockerhub にpushを実装
# TODO: github にpushを実装

upload-documents:
	make upload-document-coverage-go && \
	make upload-document-coverage-js && \
	make upload-document-api-reference

upload-document-coverage-go:
	make -C app generate-coverage && \
	aws s3 cp docs/coverage/go/index.html s3://${S3_BUCKET_NAME}/coverage/go/

upload-document-coverage-js:
#	make -C hoge hogehoge &&
#	aws s3 sync ./docs/phpdoc s3://$S3_BUCKET_NAME/coverage/js

upload-document-api-reference:
	make -C docs/openapi generate-api-document && \
	aws s3 cp docs/openapi/dist/openapi/openapi.yaml s3://${S3_BUCKET_NAME}/api-reference/ && \
	aws s3 cp docs/openapi/rapidoc/index.html s3://${S3_BUCKET_NAME}/api-reference/
