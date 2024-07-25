repo=$1
cd ../../
make build-go-service build-py-service
for svc in go-service py-service; do
  docker tag $svc $repo/$svc
  docker push $repo/$svc
done