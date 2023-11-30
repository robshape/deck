.PHONY: audit build clean install run test

audit:
	go mod verify
	go vet ./...
	go fmt ./...
	go test -coverprofile=cover.out -race -v ./...
	go tool cover -func cover.out

build:
	go build -ldflags="-s -w" -o ./bin/deck ./cmd/deck

clean:
	go clean -cache -i -modcache -r
	rm -fr ./bin/
	rm -fr cover.out

install:
	go get -u ./...
	go mod vendor
	go mod tidy

run: build
	./bin/deck

test:
	go test -race -v ./...
