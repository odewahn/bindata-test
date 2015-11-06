run: build
	./main

build:
	go-bindata-assetfs static/...
	go build -o main *.go
