

repo=$1

if [[ -z $repo ]]; then
  echo "Usage: $0 <repo>"
  exit 1
fi

MY_PATH="$(dirname -- "${BASH_SOURCE[0]}")"
MY_PATH="$(cd -- "$MY_PATH" && pwd)"
cd MY_PATH/../../
pwd
for svc in go-fasthttp py-flask; do
  make build-$svc
  docker tag $svc $repo/$svc
  docker push $repo/$svc
done