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

test:
	@echo "Testing server..."
	@cd server && go test -v ./...
	@echo "Testing client..."
	@cd client && go test -v ./...

mockgen:
	@echo "Mocking resources..."
	@cd client && \
		mkdir -p mocks && \
		mockgen github.com/cuongpiger/golang/ecommerce ProductInfoClient > mocks/product_info_client_mock.go

loadTest:
	ghz --insecure \
		--proto proto/product_info.proto \
		--call ecommerce.ProductInfo/addProduct \
		-d '{"name":"iPhone","description":"Apple iPhone","price":999}' \
		-n 2000 -c 20 localhost:50051

.PHONY: protoc runServer runClient test mockgen loadTest
