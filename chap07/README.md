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
