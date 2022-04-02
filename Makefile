build-linux:
	mkdir -p build
	GOARCH=amd64 GOOS=linux go build -o ./build cmd/tunvs/tunvs.go
	docker build -t tunvs-busybox --rm .