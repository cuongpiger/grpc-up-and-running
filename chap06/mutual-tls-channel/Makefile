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

genCAAndSelfSignedCerts:
	openssl genrsa -aes256 -out certs/ca.key -passout pass:Mypassword@123 4096
	openssl req -new -x509 -sha256 -days 730 -key certs/ca.key -out certs/ca.crt \
		-passin pass:Mypassword@123 \
		-subj "/C=VN/ST=Ho Chi Minh/L=Ho Chi Minh City/O=MyOrg/CN=localhost" \
		-addext "subjectAltName=DNS:localhost,IP:127.0.0.1"

verifyRootCACert:
	openssl x509 -noout -text -in certs/ca.crt

genServerExfile:
	echo "[v3_req]" > certs/server_ext.conf
	echo "subjectAltName=DNS:localhost,IP:127.0.0.1" >> certs/server_ext.conf

genClientExfile:
	echo "[v3_req]" > certs/client_ext.conf
	echo "subjectAltName=DNS:localhost,IP:127.0.0.1" >> certs/client_ext.conf

genPublicKeyAndCert: genServerExfile
	openssl req -new -sha256 -key certs/server.key -out certs/server.csr \
		-subj "/C=VN/ST=Ho Chi Minh/L=Ho Chi Minh City/O=MyOrg/CN=localhost" \
		-addext "subjectAltName=DNS:localhost,IP:127.0.0.1"
	openssl x509 -req -days 365 -sha256 -in certs/server.csr -CA certs/ca.crt \
		-CAkey certs/ca.key -set_serial 1 -out certs/server.crt \
		-passin pass:Mypassword@123 \
        -extfile certs/server_ext.conf \
		-extensions v3_req

genClientKeyAndCert: genClientExfile
	openssl genrsa -out certs/client.key 2048
	openssl req -new -key certs/client.key -out certs/client.csr \
		-subj "/C=VN/ST=Ho Chi Minh/L=Ho Chi Minh City/O=MyOrg/CN=localhost" \
		-addext "subjectAltName=DNS:localhost,IP:127.0.0.1"
	openssl x509 -req -days 365 -sha256 -in certs/client.csr -CA certs/ca.crt \
		-CAkey certs/ca.key -set_serial 2 -out certs/client.crt \
		-passin pass:Mypassword@123 \
        -extfile certs/client_ext.conf \
		-extensions v3_req

convertServerPrivateKeyToPEM:
	openssl pkcs8 -topk8 -inform pem -in certs/server.key -outform pem -nocrypt -out certs/server.pem

convertClientPrivateKeyToPEM:
	openssl pkcs8 -topk8 -inform pem -in certs/client.key -outform pem -nocrypt -out certs/client.pem

genCerts: genPrivateRSAServerKey genCAAndSelfSignedCerts verifyRootCACert genPublicKeyAndCert genClientKeyAndCert convertServerPrivateKeyToPEM convertClientPrivateKeyToPEM

runServer:
	cd server && go run main.go

runClient:
	cd client && go run main.go


.PHONY: protoc genPrivateRSAServerKey genCAAndSelfSignedCerts verifyRootCACert genServerExfile genClientExfile genPublicKeyAndCert genClientKeyAndCert convertServerPrivateKeyToPEM convertClientPrivateKeyToPEM convertClientPrivateKeyToPEM genCerts runServer runClient