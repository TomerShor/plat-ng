.PHONY build-go-fasthttp:
build-go-fasthttp:
	cd services/go_fasthttp && make build

.PHONY build-py-flask:
build-py-flask:
	cd services/py_flask && make build