

.PHONY: all test clean

all: test build

test:
	go test ./...

build: clean
	go build -v -o bin/server  ./main.go

protoc:
	protoc  --go-grpc_out=require_unimplemented_servers=false:./proto/ ./proto/*.proto --go_out=./proto
	protoc  --dart_out=grpc:client/lib/ proto/service.proto

run: protoc build
	./bin/server 
asset-up:
	@docker compose -f ./assets/postgres/docker-compose.yml up -d
	@envoy -c ./assets/envoy/config.yml --log-path /tmp/envoy.log &

asset-down:
	@docker compose -f ./assets/postgres/docker-compose.yml down
	@a=$$(ps -aux | grep envoy | grep -v grep | awk '{print $$2}'); \
	if [ -n "$$a" ]; then \
		kill -9 $$a; \
		echo killed; \
	fi  


clean:
	rm -f ./bin/*