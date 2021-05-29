install:
		go mod vendor

run:
		go run main.go

clean:
		rm -rf vendor bin

build: install
		go build -ldflags="-s -w" -o bin/main ./main.go

deploy: clean build

.PHONY: run install build clean deploy