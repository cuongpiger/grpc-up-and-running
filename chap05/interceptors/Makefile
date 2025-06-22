protoc:
	mkdir -p server/ecommerce
	mkdir -p client/ecommerce
	protoc \
		--go_out=Mproto/order_management.proto=.:server/ecommerce \
		--go-grpc_out=Mproto/order_management.proto=.:server/ecommerce \
		proto/order_management.proto
	protoc \
		--go_out=Mproto/order_management.proto=.:client/ecommerce \
		--go-grpc_out=Mproto/order_management.proto=.:client/ecommerce \
		proto/order_management.proto

runServer:
	cd server && go run main.go

runClient:
	cd client && go run main.go
