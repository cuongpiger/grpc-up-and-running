# Installation necessary to run the code in this chapter

Install `mockgen`

```shell
go install go.uber.org/mock/mockgen@latest

brew install ghz
```

# Mocking & Unit testing

- Working directory [`grpc-continous-integration`](./grpc-continous-integration/)

```shell
make test
```

![](./assets/02.png)

# Load testing

- Working directory [`grpc-continous-integration`](./grpc-continous-integration/)
- Run the `product_info` server

```shell
make runServer

# another terminal
make loadTest
```

![](./assets/01.png)

# OpenTelemetry for tracing

- Working directory [`grpc-opentelemetry`](./grpc-opentelemetry/)

```shell
make runServer

# another terminal
make runClient
```

![](./assets/03.png)

- Access the zPages trace visualization at `http://localhost:7777/debug/zpages/tracez`

  ![](./assets/04.png)
