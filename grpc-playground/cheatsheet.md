# Phase 1 — Core gRPC foundation

## Study

- [Introduction to gRPC](https://grpc.io/docs/what-is-grpc/introduction/)
- [Core concepts, architecture and lifecycle](https://grpc.io/docs/what-is-grpc/core-concepts/)
- [RPC life cycle](https://grpc.io/docs/what-is-grpc/core-concepts/#rpc-life-cycle)

## What is gRPC?

gRPC is an open-source Remote Procedure Call (RPC) framework. A client can call a method on a server
running on another machine as if it were a local method. The service contract defines which methods
are available and the request and response types they use.

gRPC is useful for communication between services, especially when clients and servers are written
in different languages. It provides generated, idiomatic client and server APIs and supports unary
and streaming calls.

## When would I prefer it over REST?

Choose gRPC when services need a strongly typed contract, generated clients, and efficient
service-to-service communication. It is also a natural fit when streaming is part of the API.

REST/JSON is often a better fit for public HTTP APIs, simple integrations, and direct browser
clients. The two approaches can coexist: the choice depends on the consumers and API requirements,
rather than one replacing the other.

## What are the 4 RPC types?

- **Unary RPC**: the client sends one request and receives one response.
- **Server-streaming RPC**: the client sends one request and receives an ordered stream of
  responses.
- **Client-streaming RPC**: the client sends an ordered stream of requests and receives one
  response.
- **Bidirectional-streaming RPC**: both sides send ordered streams independently.

## Core concepts

- A **service** defines a remote API in a `.proto` file.
- A **method** is an RPC operation in that service, with request and response message types.
- A **message** is a typed data structure made of named fields; it is used for requests and
  responses.
- Protocol Buffers are the default Interface Definition Language (IDL) and message format in gRPC.
  The `protoc` compiler and gRPC plugins generate Go message types plus client and server APIs from
  a `.proto` contract.

## Unary and streaming calls

A unary RPC has one request and one response, so it is the closest to a normal function call. A
streaming RPC allows the client, server, or both to exchange multiple messages. gRPC preserves
message order within each individual stream.

## How a call works

1. The client calls a generated stub (called a client in Go) with a request message and, optionally,
   a deadline.
2. gRPC sends the call over HTTP/2. Each RPC maps to an HTTP/2 stream and the messages are
   serialized Protocol Buffers carried in HTTP/2 data frames.
3. The server receives the method name and request metadata, decodes the request, and runs the
   matching service-method implementation.
4. The server sends a response message, a final status code and optional status message, plus
   optional trailing metadata. It may send initial metadata before the response.
5. With status `OK`, the client receives the response and the call completes. Either side can cancel
   a call; the client deadline can end it with `DEADLINE_EXCEEDED`. Changes made before cancellation
   are not rolled back.

# Phase 2 — gRPC in Go

## Study

- [Quick start | Go](https://grpc.io/docs/languages/go/quickstart/) - [Basics tutorial |
Go](https://grpc.io/docs/languages/go/basics/) - [Generated-code reference |
Go](https://grpc.io/docs/languages/go/generated-code/)

## Toolchain and code generation

A Go gRPC project needs Go, the Protocol Buffers compiler (`protoc`), and two Go plugins:

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
export PATH="$PATH:$(go env GOPATH)/bin"
```

The `.proto` file is the service contract. Its `package` identifies Protocol Buffers types, while
`option go_package` sets the Go import path and package name for generated code.

From the `grpc-playground` directory, regenerate all playground protobuf code with:

```bash
make protos
```

The current `protos` target runs `protos-helloworld`. When another playground example is added, give
it a dedicated target and add that target as a dependency of `protos` in the top-level `Makefile`.

The current target expands to:

```bash
protoc --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
  helloworld/helloworld/helloworld.proto
```

`paths=source_relative` keeps generated files beside their source `.proto` file. The command creates
or updates:

- `helloworld/helloworld/helloworld.pb.go`: Go types for protobuf messages and their serialization
  support.
- `helloworld/helloworld/helloworld_grpc.pb.go`: the generated client and server APIs for the gRPC
  service.

Do not edit generated `*.pb.go` files by hand. Change the `.proto` contract, then run `protoc`
again.

## From contract to running application

The current example follows this flow:

1. `helloworld.proto` declares the `Greeter` service, its unary methods `SayHello` and
   `SayHelloAgain`, and the `HelloRequest` and `HelloReply` messages.
2. `protoc` generates Go message types, `GreeterClient`, `GreeterServer`, `NewGreeterClient`, and
   `RegisterGreeterServer`.
3. The human-written server implements the generated `GreeterServer` interface.
4. The human-written client creates a generated `GreeterClient` and calls its methods.

For a unary method, the generated server interface has this shape:

```go
Method(context.Context, *Request) (*Response, error)
```

The generated client method has this shape:

```go
Method(ctx context.Context, request *Request, opts ...grpc.CallOption) (*Response, error)
```

## Server structure

The server implementation in `helloworld/greeter_server/main.go`:

1. embeds `pb.UnimplementedGreeterServer` for forward compatibility;
2. implements the service methods;
3. opens a TCP listener with `net.Listen`;
4. creates a server with `grpc.NewServer()`;
5. registers the implementation with `pb.RegisterGreeterServer`;
6. blocks in `Serve` while it accepts and dispatches RPCs.

## Client structure

The client in `helloworld/greeter_client/main.go`:

1. creates a connection with `grpc.NewClient` and closes it with `defer`;
2. obtains the generated stub with `pb.NewGreeterClient(conn)`;
3. creates a context with a one-second deadline;
4. calls `SayHello` and `SayHelloAgain` with a `HelloRequest`;
5. reads the returned `HelloReply` or handles the error.

`insecure.NewCredentials()` is suitable for this local learning example only. Use transport security
and appropriate credentials for a real service.

## Generated streaming APIs: preview

Newly generated Go streaming APIs use generics. Client RPC calls and server RPC handlers are safe to
run in concurrent goroutines. Within one stream, however, do not perform concurrent reads or
concurrent writes; one read and one write can proceed independently. Streaming implementation
belongs to Phase 4.

# Phase 3A — Compatibility rules

Treat the `.proto` schema as a long-lived API contract.

- Field numbers are part of the wire format. Never reuse a field number for a different meaning.
- Adding a new field with a new number is usually backward-compatible; old clients ignore it.
- Adding a new RPC, message, or enum value is usually safe when existing meanings remain unchanged.
- Changing a field type, number, meaning, package, or existing RPC shape can break compatibility.
- When removing a field, reserve its number and name:

```proto
message Product {
  reserved 2;
  reserved "old_name";
  string id = 1;
}
```

Do not confuse a source-level rename with a wire-compatible change: the field number and serialized
meaning are what matter. Detailed `v1`/`v2` practice belongs to Phase 3A; Buf validation starts in
Phase 3B.

# Phase 3 map

- **Phase 3A — Protobuf contract design:** learn `.proto`, compatibility, `v1`/`v2`, and generated
  Go API using only `protoc` and the Makefile.
- **Phase 3B — Buf contract workflow:** add format, lint, breaking checks, reproducible generation,
  and dependency/version basics.
- **Phase 3C — Tools around protobuf/gRPC:** compare ConnectRPC, expose the same contract through
  grpc-gateway/OpenAPI, diagnose with grpcurl, and add OpenTelemetry.

The detailed sequence belongs in the [gRPC + Microservices
roadmap](../docs/learning/grpc-microservices-roadmap.md). The primary path remains `grpc-go +
protobuf + Buf`; Phase 3C extends it without replacing typed Go tests or the underlying protobuf
workflow.
