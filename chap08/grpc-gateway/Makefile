protoc:
	mkdir -p server/ecommerce
	mkdir -p client/ecommerce
	buf mod update
	buf generate

runServer:
	cd server && go run main.go

runClient:
	cd client && go run main.go

.PHONY: protoc runServer runClient dockerUp dockerDown
