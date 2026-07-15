# Go Microservices Starter

An educational monorepo for learning Go microservices, gRPC, and the practical
concerns around service-based systems. The repository intentionally contains
independent learning tracks rather than one deployable application.

## Current focus

The active track is [`grpc-playground/`](grpc-playground/). Phase 1 and Phase 2
of the gRPC roadmap are complete; the next topic is Protocol Buffers and
contract design.

- [gRPC learning roadmap](docs/learning/grpc-microservices-roadmap.md)
- [gRPC playground notes](grpc-playground/cheatsheet.md)
- [AI repository guidance](AGENTS.md)

## Repository layout

| Path | Description |
| --- | --- |
| [`grpc-playground/`](grpc-playground/) | Small Go gRPC examples. The current `helloworld` example has a unary server, client, generated protobuf code, and an in-memory integration test. |
| [`tutorials/building-microservices-youtube/`](tutorials/building-microservices-youtube/) | Exercises based on the [Building Microservices in Go](https://www.youtube.com/playlist?list=PLmD8u-IFdreyh6EUfevBcbiuCKzFk0EW_) series. Contains currency, product API, product images, and frontend projects. |
| [`books/grpc-microservices-in-go/`](books/grpc-microservices-in-go/) | Notes and exercises from *gRPC Microservices in Go*. |
| [`docs/`](docs/) | Learning roadmap and AI-maintained repository documentation. |

## Requirements

- Go 1.26.4
- For protobuf regeneration: `protoc`, `protoc-gen-go`, and
  `protoc-gen-go-grpc` available on `PATH`.

The repository has a `go.work` workspace for `grpc-playground` and the tutorial
Go modules. The Chapter 3 book module is standalone and must be used outside
that workspace.

## Quick start: gRPC playground

Regenerate the Go bindings after changing the `.proto` contract:

```bash
make -C grpc-playground protos
```

Run the test suite:

```bash
(cd grpc-playground && go test ./...)
```

Run the server, then run the client in another terminal:

```bash
(cd grpc-playground && go run ./helloworld/greeter_server)
(cd grpc-playground && go run ./helloworld/greeter_client --name=Lee)
```

The server listens on `:50051` by default. The example uses insecure transport
credentials only for local learning.

## Working with modules

There is no root `go.mod`. Run Go commands from the module you are changing.
For example:

```bash
(cd tutorials/building-microservices-youtube/currency && go test ./...)
(cd tutorials/building-microservices-youtube/product-api && go test ./...)
(cd tutorials/building-microservices-youtube/product-images && go test ./...)
(cd books/grpc-microservices-in-go/chapter03/golang/order && GOWORK=off go test ./...)
```

The tutorial projects preserve the progression of their source material. Avoid
cross-module dependency upgrades or broad refactors unless that is the explicit
task.

## Generated code

Treat `.proto` files as API source. Do not edit generated `*.pb.go` or
`*_grpc.pb.go` files manually; regenerate them with the appropriate Make target
and review the resulting diff.

## Documentation

- [`docs/learning/grpc-microservices-roadmap.md`](docs/learning/grpc-microservices-roadmap.md): phased gRPC study plan.
- [`docs/ai/repo-context.md`](docs/ai/repo-context.md): durable repository and workflow context.
- [`docs/ai/task-history.md`](docs/ai/task-history.md): append-only history of completed work.
