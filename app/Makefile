

build-golang:
	go build ./main.go

run: build-golang
	./main

cp-openapi-skeleton-code: generate-openapi-skeleton-code
	cp -rfv ./../docs/openapi/skeleton/go/* ./openapi/

generate-openapi-skeleton-code:
	make -C ./../docs/openapi generate-api-skeleton

# go-test:
# 	go test -v \
# 		./mongodb \
# 		./openapi \
# 		./validate
