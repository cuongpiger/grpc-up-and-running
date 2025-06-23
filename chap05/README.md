# Interceptors

In Go, **interceptors** in gRPC provide a mechanism to execute common logic **before or after the execution of a remote function** for both client and server applications. They are considered a key extension mechanism for purposes like logging, authentication, metrics, and tracing.

Go interceptors are categorized into two main types based on the RPC call they intercept:

- **Unary Interceptors**: These are used for simple request-response RPCs.

  - **Server-Side**: You implement a function of type `UnaryServerInterceptor`. This function receives the `context.Context`, the request, server information, and a `UnaryHandler` that is used to invoke the actual RPC method. You can add preprocessing logic before calling `handler(ctx, req)` and postprocessing logic after. They are registered with the gRPC server using `grpc.UnaryInterceptor()` when creating the server instance.
  - **Client-Side**: You implement a function of type `UnaryClientInterceptor`. This function gives you access to the RPC context, method name, request, reply, client connection, and an `UnaryInvoker` to call the actual remote method. You can modify the RPC call before it's sent and process the response or error afterward. They are registered using `grpc.WithUnaryInterceptor()` when setting up the client connection via `grpc.Dial`.

- **Streaming Interceptors**: These intercept streaming RPCs (server-streaming, client-streaming, or bidirectional-streaming).
  - **Server-Side**: You implement a function of type `StreamServerInterceptor`. Within this interceptor, a `wrappedStream` that implements `grpc.ServerStream` is often used to intercept the `RecvMsg` (for incoming messages) and `SendMsg` (for outgoing messages) methods. This allows you to process messages as they are received or sent over the stream. They are registered using `grpc.StreamInterceptor()` when creating the gRPC server.
  - **Client-Side**: You implement a function of type `StreamClientInterceptor`. Similar to the server side, a `wrappedStream` that implements `grpc.ClientStream` can be used to intercept `RecvMsg` and `SendMsg` for client-side stream operations. They are registered using `grpc.WithStreamInterceptor()` when dialing the gRPC connection.

The Go gRPC Middleware project further extends this concept by providing **interceptor chaining**, allowing you to apply multiple interceptors sequentially for both unary and streaming RPCs on both the client and server sides.

My demonstration of interceptors in this chapter includes:
![](./assets/01.png)

## Deadlines

In the "Deadlines" section of Chapter 5, the following important knowledge is presented regarding Go (Golang):

- **Deadlines** are a crucial concept in distributed computing, specifically in gRPC applications, allowing a client to specify an **absolute time by which an RPC must complete**. This is distinct from timeouts, which are durations applied locally.
- It is considered **good practice to use deadlines** in gRPC applications to prevent clients from infinitely waiting for responses, which can lead to resource exhaustion and increased latency.
- In **Go**, setting a deadline for an RPC is achieved using the `context` package, specifically with the **`context.WithDeadline`** operation. This operation creates a new context with a specified absolute time.
- When a deadline is set by the client and the RPC does not complete within that time, the RPC call is **terminated with a `DEADLINE_EXCEEDED` error**.
- The **gRPC library on the client side translates the deadline** set in the context into a required gRPC header at the wire level (on HTTP/2).
- On the **server side (in Go)**, you can detect if the client has reached its deadline by checking **`ctx.Err() == context.DeadlineExceeded`** within the remote method implementation. This allows the server to abandon the RPC if the client no longer needs the response, returning an error.

My demonstration of deadlines in this chapter includes:
![](./assets/02.png)

## Cancellation

In Go, **cancellation** provides a mechanism for either the client or the server application to **terminate ongoing gRPC communication**. This is crucial because, in a gRPC connection, both the client and server independently determine the success of an RPC, meaning their conclusions about its outcome might differ. When one party cancels an RPC, this fact is **propagated to the other side**, and no further RPC-related messaging can occur for that call.

Here's how cancellation is handled in Go:

- **Context Package**: The capability for cancellation in Go is provided through the **`context` package**. Specifically, the `context.WithCancel` function (or `context.WithTimeout`, which implicitly includes cancellation functionality) is used to create a new context with a cancel function.
- **Client-Side Initiation**: On the client side, after obtaining the `cancel` function from a `context.WithTimeout` or `context.WithCancel` call, the client can invoke this `cancel()` function at any point when it wishes to terminate the RPC. This action leads the gRPC library on the client side to **create a required gRPC header** to signal the termination to the server.
- **Server-Side Detection**: On the server side, the remote method implementation can **detect if the client has canceled the RPC** by checking the context using `stream.Context().Err() == context.Canceled`. This allows the server to abandon its processing for that RPC if the client no longer requires a response, preventing unnecessary resource usage.

In essence, cancellation is a fundamental feature for robust distributed systems, enabling graceful termination of RPCs and preventing resource leaks when either party no longer needs to continue the communication.

My demonstration of cancellation in this chapter includes:
![](./assets/03.png)
