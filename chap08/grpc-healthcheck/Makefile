protoc:
	mkdir -p server/healthcheck
	protoc \
		--go_out=Mproto/healthcheck.proto=.:server/healthcheck \
		--go-grpc_out=Mproto/healthcheck.proto=.:server/healthcheck \
		proto/healthcheck.proto

runServer:
	cd server && go run main.go

.PHONY: protoc runServer
