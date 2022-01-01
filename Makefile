################################################################################
# Required environment variables
################################################################################
IMAGE_OWNER :=
IMAGE_NAME :=
RELEASE_VERSION :=

PAT_GITHUB :=
PAT_DOCKERHUB :=


################################################################################
# Other environment variables
################################################################################
PROJECT_NAME := testing-file-generator
GITHUB_USER := ${IMAGE_OWNER}
DOCKERHUB_USER := ${IMAGE_OWNER}

IMAGE_NAME_FOR_DH := ${DOCKERHUB_USER}/${IMAGE_NAME}
IMAGE_NAME_LATEST_FOR_DH := ${IMAGE_NAME_FOR_DH}:latest
IMAGE_NAME_VERSION_FOR_DH := ${IMAGE_NAME_FOR_DH}:${RELEASE_VERSION}

IMAGE_NAME_FOR_GHCR := ghcr.io/${GITHUB_USER}/${IMAGE_NAME}
IMAGE_NAME_LATEST_FOR_GHCR := ${IMAGE_NAME_FOR_GHCR}:latest
IMAGE_NAME_VERSION_FOR_GHCR := ${IMAGE_NAME_FOR_GHCR}:${RELEASE_VERSION}

S3_BUCKET_NAME := ${PROJECT_NAME}


################################################################################
# イメージビルド・デプロイタスク
################################################################################

# NOTE: 手動実行用のタスク（CI/CDではこのtarget内の処理を実行する）
# make deploy-dh -e IMAGE_OWNER={} IMAGE_NAME={} RELEASE_VERSION={} PAT_GITHUB={} PAT_DOCKERHUB={}
deploy-dh:
	make build && \
	make set-tag-dh && \
	make login-dh && \
	make push-dh

# NOTE: 手動実行用のタスク（CI/CDではこのtarget内の処理を実行する）
# make deploy-ghcr -e IMAGE_OWNER={} IMAGE_NAME={} RELEASE_VERSION={} PAT_GITHUB={} PAT_DOCKERHUB={}
deploy-ghcr:
	make build && \
	make set-tag-ghcr && \
	make login-ghcr && \
	make push-ghcr

# make build -e IMAGE_NAME={}
build:
	docker-compose build --no-cache --force-rm ${IMAGE_NAME}


# make set-tag-dh -e IMAGE_OWNER={} IMAGE_NAME={} RELEASE_VERSION={}
set-tag-dh:
	docker tag ${IMAGE_NAME} ${IMAGE_NAME_VERSION_FOR_DH} && \
	docker tag ${IMAGE_NAME} ${IMAGE_NAME_LATEST_FOR_DH}

# make login-dh -e PAT_DOCKERHUB={} DOCKERHUB_USER={}
login-dh:
	@echo ${PAT_DOCKERHUB} | docker login -u ${DOCKERHUB_USER} --password-stdin

# make push-dh -e IMAGE_OWNER={} IMAGE_NAME={} RELEASE_VERSION={}
push-dh:
	docker push ${IMAGE_NAME_VERSION_FOR_DH} && \
	docker push ${IMAGE_NAME_LATEST_FOR_DH}


# make set-tag-ghcr -e IMAGE_OWNER={} IMAGE_NAME={} RELEASE_VERSION={}
set-tag-ghcr:
	docker tag ${IMAGE_NAME} ${IMAGE_NAME_LATEST_FOR_GHCR} && \
	docker tag ${IMAGE_NAME} ${IMAGE_NAME_VERSION_FOR_GHCR}

# make set-tag-ghcr -e PAT_DOCKERHUB={} DOCKERHUB_USER={}
login-ghcr:
	@echo ${PAT_GITHUB} | docker login ghcr.io -u ${GITHUB_USER} --password-stdin

# make push-ghcr -e IMAGE_OWNER={} IMAGE_NAME={} RELEASE_VERSION={}
push-ghcr:
	docker push ${IMAGE_NAME_VERSION_FOR_GHCR} && \
	docker push ${IMAGE_NAME_LATEST_FOR_GHCR}


################################################################################
# ドキュメント・リファレンス生成タスク
################################################################################
upload-documents:
	make upload-document-coverage-go && \
	make upload-document-coverage-js && \
	make upload-document-api-reference

upload-document-coverage-go:
	make -C app generate-coverage && \
	aws s3 cp docs/coverage/go/index.html s3://${S3_BUCKET_NAME}/coverage/go/

upload-document-coverage-js:
#	TODO: フロント実行時に作成

upload-document-api-reference:
	make -C docs/openapi generate-api-document && \
	aws s3 cp docs/openapi/dist/openapi/openapi.yaml s3://${S3_BUCKET_NAME}/api-reference/ && \
	aws s3 cp docs/openapi/rapidoc/index.html s3://${S3_BUCKET_NAME}/api-reference/








