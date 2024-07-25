.PHONY: fmt
fmt:
	gofmt -s -w .

.PHONY: build
build:
	docker build --file ./cmd/server/Dockerfile -t go-service .
