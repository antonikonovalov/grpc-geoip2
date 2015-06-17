
all: proto

proto:
	 @protoc --go_out=plugins=grpc:. geoip2/*.proto
