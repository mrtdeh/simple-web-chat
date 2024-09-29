

.PHONY: all test clean

all: test build

test:
	go test ./...

build: clean
	go build -v -o bin/server  ./main.go

protoc:
	protoc --go-grpc_out=require_unimplemented_servers=false:./proto/ ./proto/*.proto --go_out=./proto

run: protoc build
	./bin/server


clean:
	rm -f ./bin/*