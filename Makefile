TEST_ARGS :=

.PHONY: build
build:
	@mkdir -p build
	CGO_ENABLED=0 go build -o build ./...

.PHONY: fmt
fmt:
	gofmt -w acme

.PHONY: clean
clean:
	rm -rf build/*

.PHONY: test
test:
	@CGO_ENABLED=0 go test $(TEST_ARGS) ./acme

.PHONY: testacc
testacc: build
	@CGO_ENABLED=0 go test $(TEST_ARGS) ./test
