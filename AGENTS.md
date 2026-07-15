# Repository guidance for AI agents

This repository is a learning monorepo for Go microservices, gRPC, a YouTube
tutorial series, and book exercises. Preserve its educational structure: make
small, scoped changes and do not refactor unrelated examples into a shared
application.

Read [docs/ai/repo-context.md](docs/ai/repo-context.md) before non-trivial
work. Record completed implementation tasks in
[docs/ai/task-history.md](docs/ai/task-history.md).

## Working rules

- Inspect the target module before changing it. There is no root `go.mod`; run
  Go commands from the relevant module directory.
- Preserve unrelated working-tree changes. Do not reset, discard, or reformat
  files outside the requested scope.
- Treat `.proto` files as source contracts. Never edit generated `*.pb.go` or
  `*_grpc.pb.go` files manually; regenerate them with the owning project's
  documented command.
- Keep dependency changes local to the module that needs them. Use `go mod tidy`
  only when the task requires dependency changes.
- Prefer tests that do not require external services or fixed ports. Run the
  smallest relevant test suite, then report the exact command and outcome.
- Update documentation when a change affects project structure, commands,
  generated-code workflow, or an AI-relevant convention.

## Common commands

```bash
# Primary gRPC learning project
make -C grpc-playground protos
(cd grpc-playground && go test ./...)

# Tutorial modules: run from the target module
(cd tutorials/building-microservices-youtube/currency && go test ./...)
(cd tutorials/building-microservices-youtube/product-api && go test ./...)
(cd tutorials/building-microservices-youtube/product-images && go test ./...)
```

For the full layout and module-specific notes, see
[docs/ai/repo-context.md](docs/ai/repo-context.md).
