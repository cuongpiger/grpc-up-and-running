protoc:
	mkdir -p server/ecommerce
	protoc \
		--go_out=Mproto/product_info.proto=.:server/ecommerce \
		--go-grpc_out=Mproto/product_info.proto=.:server/ecommerce \
		proto/product_info.proto

runServer:
	cd server && go run main.go

.PHONY: protoc runServer runClient dockerUp dockerDown
