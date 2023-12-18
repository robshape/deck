.PHONY: audit benchmark build clean install run test

audit:
	go mod verify
	go vet ./{cmd,internal,pkg}/...
	go fmt ./{cmd,internal,pkg}/...
	go test -coverprofile=cover.out -race -v ./{cmd,internal,pkg}/...
	#go tool cover -html cover.out -o cover.html

benchmark:
	go test -bench=Bench -benchmem -benchtime=100x ./{cmd,internal,pkg}/...

build:
	go build -ldflags="-s -w" -o ./bin/deck ./cmd/deck

clean:
	go clean -cache -i -modcache -r
	rm -fr ./bin/
	rm -fr cover.out

install:
	go get -u ./{cmd,internal,pkg}/...
	go mod vendor
	go mod tidy

run: build
	./bin/deck

test:
	go test -race ./{cmd,internal,pkg}/...
