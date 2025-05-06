# api makefile

MOCK_TARGETS=$(shell find . -type f -name "*.go" | grep -v "testing/" | grep -v "_test.go" | grep -v "dst/")

build_api:
	go build -o ./dst/api apps/api/main.go
build_batch:
	go build -o ./dst/batch apps/batch/main.go

build: build_api build_batch

test:
	go test --short ./...

test_integral:
	go test ./...

gen:
	sqlboiler mysql

gen_swag:
	(cd openapi/src && npm run compile && npm run lint)
	oapi-codegen -config ./openapi/src/oapi-codegen.yaml openapi/openapi.yaml

genmock: $(MOCK_TARGETS)

$(MOCK_TARGETS):
	mockgen -source $@ -destination testing/mocks/$@

clean:
	rm -rf ./testing/mocks/
	rm -rf ./dst

fmt:
	go fmt ./...
