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

genPrivateRSAServerKey:
	mkdir -p certs
	openssl genrsa -out certs/server.key 2048

genPublicKeyAndCert: genPrivateRSAServerKey
	openssl req -new -x509 -sha256 -key certs/server.key -out certs/server.crt -days 3650 \
		-subj "/C=VN/ST=Ho Chi Minh/L=Ho Chi Minh City/O=MyOrg/CN=localhost" \
		-addext "subjectAltName=DNS:localhost,IP:127.0.0.1"

convertServerPrivateKeyToPEM:
	openssl pkcs8 -topk8 -inform pem -in certs/server.key -outform pem -nocrypt -out certs/server.pem

runServer:
	cd server && go run main.go

runClient:
	cd client && go run main.go


.PHONY: protoc genPrivateRSAServerKey genPublicKeyAndCert convertServerPrivateKeyToPEM convertClientPrivateKeyToPEM runServer runClient