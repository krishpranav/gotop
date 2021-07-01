cd ..
mkdir -p bin

archs=("amd64" "386" "arm" "arm64")

for arch in ${archs[@]}; do
  env GOOS=linux GOARCH=${arch} go build ./gotop.go
  cp gotop bin/gotop_${arch}
  echo "Compiled gotop_${arch}"
done