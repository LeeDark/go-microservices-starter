## Study
- [Introduction to gRPC](https://grpc.io/docs/what-is-grpc/introduction/)
- [Core concepts, architecture and lifecycle](https://grpc.io/docs/what-is-grpc/core-concepts/)
- [RPC life cycle](https://grpc.io/docs/what-is-grpc/core-concepts/#rpc-life-cycle)

## Write a short personal note answering:

- What is gRPC?
    - gRPC is a modern open source high performance Remote Procedure Call (RPC) framework that can run in any environment. It can efficiently connect services in and across data centers with pluggable support for load balancing, tracing, health checking and authentication. It is also applicable in last mile of distributed computing to connect devices, mobile applications and browsers to backend services.
    - The main usage scenarios
        - Efficiently connecting polyglot services in microservices style architecture
        - Connecting mobile devices, browser clients to backend services
        - Generating efficient client libraries
    - Core features that make it awesome
        - Idiomatic client libraries in 11 languages
        - Highly efficient on wire and with a simple service definition framework
        - Bi-directional streaming with http/2 based transport
        - Pluggable auth, tracing, load balancing and health checking

- When would I prefer it over REST?
    - gRPC largely follows HTTP semantics over HTTP/2 but we explicitly allow for full-duplex streaming. We diverge from typical REST conventions as we use static paths for performance reasons during call dispatch as parsing call parameters from paths, query parameters and payload body adds latency and complexity. We have also formalized a set of errors that we believe are more directly applicable to API use cases than the HTTP status codes.

- What are the 4 RPC types?
    - Unary RPC
    - Server streaming RPC
    - Client streaming RPC
    - Bidirectional streaming RPC

- What a **service**, **method**, and **message** are
    - A **service** defines a remote API in a .proto file
    - A **method** is an RPC operation inside that service with request and response types
    - A **message** is a typed data structure made of fields, used as the input or output of RPC methods.

- The difference between **unary** and **streaming** RPC
    - A unary RPC sends one request and receives one response, like a normal function call with network misery attached
    - A streaming RPC lets the client, the server, or both sides send multiple messages over a stream while preserving message order within each stream.

- How a client and server communicate over HTTP/2
    - In gRPC, the client calls a local stub/client method, and gRPC sends the request to the server over HTTP/2.
    - Under the hood, RPCs map to HTTP/2 streams, and gRPC messages are carried using HTTP/2 data frames.

- Why Protocol Buffers matter
    - Protocol Buffers provide a language-neutral, platform-neutral way to define and serialize structured data.
    - They are smaller and faster than JSON in many cases, generate native language bindings, and help keep service contracts explicit.

- What happens during one RPC call from request to response
    - The client calls a generated stub method, which wraps the request into a Protocol Buffer message and sends it to the server.
    - The server decodes the request, runs the service method implementation, encodes the response, and sends it back to the client.

