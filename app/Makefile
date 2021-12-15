
run: build
	./main

build:
	go build ./main.go

test:
	go test -v \
		./datastructure/queue \
		./openapi \
		./util \
		./verify \
		./validate \
		./util

cp-openapi-skeleton-code: generate-openapi-skeleton-code
	cp -rfv ./../docs/openapi/skeleton/go/* ./openapi/

generate-openapi-skeleton-code:
	make -C ./../docs/openapi generate-api-skeleton

