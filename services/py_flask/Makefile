.PHONY: build
build: ## Build the service
	@echo "Building service"
	docker build \
		--file ../../deployment/dockerfiles/py_flask/Dockerfile \
		-t py-flask \
		.