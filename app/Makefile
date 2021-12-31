
COVERAGE_DIR := ../docs/coverage/go
COVERAGE_FILE_PATH := $(COVERAGE_DIR)/cover.out

run: build
	./main

build:
	go build ./main.go

test:
	go test -v \
		-cover \
		-coverprofile=$(COVERAGE_FILE_PATH) \
		./datastructure/queue \
		./openapi \
		./util \
		./verify \
		./validate \
		./util

generate-coverage: test
	go tool cover -html=$(COVERAGE_FILE_PATH) -o $(COVERAGE_DIR)/index.html && rm $(COVERAGE_FILE_PATH)

generate-openapi-skeleton-code:
	make -C ./../docs/openapi generate-api-skeleton

cp-openapi-skeleton-code: generate-openapi-skeleton-code
	cp -rfv ./../docs/openapi/skeleton/go/* ./openapi/

