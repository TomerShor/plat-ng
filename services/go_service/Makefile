.PHONY: fmt
fmt:
	gofmt -s -w .

.PHONY: build
build:
	docker build \
		--file ../../deployment/dockerfiles/go_service/Dockerfile \
		-t go-service \
		.
