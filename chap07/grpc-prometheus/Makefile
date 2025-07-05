protoc:
	mkdir -p server/ecommerce
	mkdir -p client/ecommerce
	protoc \
		--go_out=Mproto/product_info.proto=.:server/ecommerce \
		--go-grpc_out=Mproto/product_info.proto=.:server/ecommerce \
		proto/product_info.proto
	protoc \
		--go_out=Mproto/product_info.proto=.:client/ecommerce \
		--go-grpc_out=Mproto/product_info.proto=.:client/ecommerce \
		proto/product_info.proto

runServer:
	cd server && go run main.go

runClient:
	cd client && go run main.go

.PHONY: protoc runServer runClient dockerUp dockerDown
