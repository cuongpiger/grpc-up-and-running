# Enabling a One-Way Secured Connection

The "Enabling a One-Way Secured Connection" section in Chapter 6 focuses on how to establish **Transport Level Security (TLS)** for gRPC communication where **only the client validates the server's identity**. This ensures that the client is indeed communicating with the intended server, providing privacy and data integrity through symmetric cryptography and message integrity checks.

Here's a summary of the key information:

- **Purpose**: In a one-way secured connection, the primary goal is for the client to **validate the server's identity**. This is crucial for ensuring that data is received from a trusted source.
- **Mechanism**: When a connection is established, the server shares its **public certificate** with the client. The client then validates this certificate, often through a Certificate Authority (CA) if it's a CA-signed certificate. Once validated, the client can send encrypted data using a secret key negotiated for the session.
- **Required Keys and Certificates**: To enable TLS, you need to create:
  - `server.key`: A private RSA key used to sign and authenticate the public key.
  - `server.pem`/`server.crt`: Self-signed X.509 public keys for distribution.
  - Tools like OpenSSL, mkcert, or certstrap can be used for generating these keys and certificates.
- **Enabling on a gRPC Server (Go example)**:
  - The server must be initialized with a **public/private key pair**.
  - In Go, this involves loading the X.509 key pair (`server.crt` and `server.key`).
  - Then, you enable TLS for incoming connections by adding these certificates as **TLS server credentials** using `grpc.Creds(credentials.NewServerTLSFromCert(&cert))` when creating a new gRPC server instance.
- **Enabling on a gRPC Client (Go example)**:
  - The client needs the **server's self-certified public key** to connect securely.
  - In Go, you create client TLS credentials from the server's public certificate file (`server.crt`) using `credentials.NewClientTLSFromFile(crtFile, hostname)`.
  - These transport credentials are then passed as a **DialOption** (`grpc.WithTransportCredentials(creds)`) when setting up a secure connection with the server using `grpc.Dial`. This process initiates the TLS handshake.

It's important to note that this one-way TLS authentication **only authenticates the server's identity** to the client; it does not authenticate the client's identity to the server.

The implementation is located in the [secure-channel](./secure-channel) directory.

My demonstration of load balancing in this chapter includes:
![](./assets/01.png)