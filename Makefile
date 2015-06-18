all: deps proto build install

proto:
	 @protoc --go_out=plugins=grpc:. geoip2/*.proto

build:
	go build -o geoip2-server main.go
	go build -o geoip2-client examples/client.go

install:
	mv geoip2-server $$GOPATH/bin
	mv geoip2-client $$GOPATH/bin

deps:
	go get -v ./...

