.PHONY build-go-service:
build-go-service:
	cd services/go_service && make build

.PHONY build-py-service:
build-py-service:
	cd services/py_service && make build