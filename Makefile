install-http-server:
	npm install --global http-server
run:
	http-server .
build:
	GOOS=js GOARCH=wasm go build -o main.wasm ./wasm
test:
	go test ./pkg/...
