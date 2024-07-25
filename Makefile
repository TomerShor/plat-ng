.PHONY build-go-service:
build-go-service:
	cd services/go-service && make build

.PHONY build-py-service:
build-py-service:
	cd services/py-service && make build