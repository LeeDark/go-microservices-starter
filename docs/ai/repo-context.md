# Repository context

## Purpose and layout

This is an educational monorepo for learning Go microservices and gRPC. Its
directories are intentionally independent learning tracks, not parts of one
production service.

| Path | Role |
| --- | --- |
| `grpc-playground/` | Primary hands-on gRPC learning project. The current example is `helloworld`. |
| `docs/learning/grpc-microservices-roadmap.md` | Phased gRPC learning roadmap. |
| `tutorials/building-microservices-youtube/` | Code from the Building Microservices YouTube series. |
| `books/grpc-microservices-in-go/` | Book notes and exercises. |

The repository has a `go.work` workspace, but no root `go.mod`. The workspace
currently includes `grpc-playground` and the three tutorial modules:
`currency`, `product-api`, and `product-images`. The book's Chapter 3 module is
standalone and is not listed in `go.work`.

## Active gRPC playground

`grpc-playground/` is a separate Go module and the current focus for the gRPC
roadmap.

- `cheatsheet.md` contains notes for completed roadmap phases.
- `helloworld/helloworld/helloworld.proto` is the source of truth for the
  `Greeter` contract.
- `helloworld/helloworld/helloworld.pb.go` and
  `helloworld/helloworld/helloworld_grpc.pb.go` are generated files.
- `helloworld/greeter_server/` and `helloworld/greeter_client/` are the
  hand-written server and client applications.
- `helloworld/greeter_server/main_test.go` is an in-memory gRPC integration
  test using `bufconn`; it avoids network ports.

### Protobuf workflow

Run this command from the repository root:

```bash
make -C grpc-playground protos
```

The top-level playground `Makefile` owns code generation. `protos` is an
aggregating target; each playground example has its own target such as
`protos-helloworld`. Add a target and include it in `protos` when adding another
example. Generated code uses source-relative paths.

After changing a `.proto` contract, regenerate its Go files and run:

```bash
(cd grpc-playground && go test ./...)
```

## Tutorial and book code

The tutorial directories preserve code from a step-by-step external series.
Avoid broad modernizations, dependency alignment, or structural rewrites unless
explicitly requested. Each module has its own dependency versions and may use
older APIs intentionally.

- `currency/` is a gRPC currency service. Regenerate its protobuf code with
  `make -C tutorials/building-microservices-youtube/currency protos`.
- `product-api/` and `product-images/` are HTTP services from the tutorial.
- `frontend/` is a separate React application; do not run Node package commands
  unless the task targets it.
- `books/grpc-microservices-in-go/chapter03/golang/order/` is a separate module.
  Run Go commands from that directory.

## Change and verification conventions

1. Start with the smallest directory that contains the requested work.
2. Treat generated files, SDK files, and tutorial snapshots as intentional.
3. Use `gofmt` for changed Go files.
4. Run tests from the changed module, not from the repository root.
5. Use `git diff --check` before handoff when files were edited.
6. Append a concise entry to `docs/ai/task-history.md` after an implementation
   task is complete.

## Documentation conventions

- `AGENTS.md` is the short mandatory entrypoint for agents.
- `docs/ai/repo-context.md` explains the repository and durable workflows.
- `docs/ai/task-history.md` is append-only and records completed changes,
  affected files, and verification. Do not rewrite earlier entries.
