# Phase 1 — Core gRPC foundation

## Study

- [Introduction to gRPC](https://grpc.io/docs/what-is-grpc/introduction/)
- [Core concepts, architecture and lifecycle](https://grpc.io/docs/what-is-grpc/core-concepts/)
- [RPC life cycle](https://grpc.io/docs/what-is-grpc/core-concepts/#rpc-life-cycle)

## What is gRPC?

gRPC is an open-source Remote Procedure Call (RPC) framework. A client can call
a method on a server running on another machine as if it were a local method.
The service contract defines which methods are available and the request and
response types they use.

gRPC is useful for communication between services, especially when clients and
servers are written in different languages. It provides generated, idiomatic
client and server APIs and supports unary and streaming calls.

## When would I prefer it over REST?

Choose gRPC when services need a strongly typed contract, generated clients,
and efficient service-to-service communication. It is also a natural fit when
streaming is part of the API.

REST/JSON is often a better fit for public HTTP APIs, simple integrations, and
direct browser clients. The two approaches can coexist: the choice depends on
the consumers and API requirements, rather than one replacing the other.

## What are the 4 RPC types?

- **Unary RPC**: the client sends one request and receives one response.
- **Server-streaming RPC**: the client sends one request and receives an ordered
  stream of responses.
- **Client-streaming RPC**: the client sends an ordered stream of requests and
  receives one response.
- **Bidirectional-streaming RPC**: both sides send ordered streams independently.

## Core concepts

- A **service** defines a remote API in a `.proto` file.
- A **method** is an RPC operation in that service, with request and response
  message types.
- A **message** is a typed data structure made of named fields; it is used for
  requests and responses.
- Protocol Buffers are the default Interface Definition Language (IDL) and
  message format in gRPC. The `protoc` compiler and gRPC plugins generate Go
  message types plus client and server APIs from a `.proto` contract.

## Unary and streaming calls

A unary RPC has one request and one response, so it is the closest to a normal
function call. A streaming RPC allows the client, server, or both to exchange
multiple messages. gRPC preserves message order within each individual stream.

## How a call works

1. The client calls a generated stub (called a client in Go) with a request
   message and, optionally, a deadline.
2. gRPC sends the call over HTTP/2. Each RPC maps to an HTTP/2 stream and the
   messages are serialized Protocol Buffers carried in HTTP/2 data frames.
3. The server receives the method name and request metadata, decodes the
   request, and runs the matching service-method implementation.
4. The server sends a response message, a final status code and optional status
   message, plus optional trailing metadata. It may send initial metadata before
   the response.
5. With status `OK`, the client receives the response and the call completes.
   Either side can cancel a call; the client deadline can end it with
   `DEADLINE_EXCEEDED`. Changes made before cancellation are not rolled back.
