build:
	protoc -I api/ --go_out=plugins=grpc:./api api/hello.proto
	docker build -t grpc-server .

run:
	docker run -p 9090:9090 grpc-server
