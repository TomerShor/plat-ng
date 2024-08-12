# palt-ng Helm Chart!

## Installation

```bash
$ helm install plat-ng-test --namespace test --create-namespace .
```

## Uninstall

```bash
$ helm uninstall plat-ng-test --namespace test
```

## Development and Testing

Once both pods are up and running, port-forward both services to your local machine:

```bash
$ kubectl -n test port-forward services/go-fasthttp http 5000:8000
$ kubectl -n test port-forward services/py-flask http 5001:8000
```

Now you can access the services at:

- go-fasthttp: http://localhost:5000
- py-flask: http://localhost:5001

Try to proxy requests between the services, e.g:
```bash
$ curl -X http://localhost:5000/py-proxy?path=runtime
$ curl -X http://localhost:5001/go-proxy?path=hello
```