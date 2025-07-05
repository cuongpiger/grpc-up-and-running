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

# OpenTelemetry for metrics

- Working directory [`grpc-opentelemetry`](./grpc-opentelemetry/)

```shell
make runServer

# another terminal
make runClient
```

![](./assets/03.png)

- Access the zPages trace visualization at `http://localhost:7777/debug/zpages/tracez`

  ![](./assets/04.png)

# OpenTelemetry for tracing

- Working directory [`grpc-otel-tracing`](./grpc-otel-tracing/)

```shell
make dockerUp
make runServer
# another terminal
make runClient
```

![](./assets/07.png)

- Access the Jaeger UI at `http://localhost:16686`

  ![](./assets/05.png)
  ![](./assets/06.png)

# Prometheus for metrics
- Working directory [`grpc-prometheus`](./grpc-prometheus/)

```shell
make runServer
# another terminal
make runClient
```

![](./assets/09.png)
- Access the Prometheus UI at `http://localhost:9092`

  ![](./assets/08.png)