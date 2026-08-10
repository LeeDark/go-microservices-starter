# AI task history

This is an append-only record of completed implementation work. Add the newest
entry at the end. Keep entries factual and concise.

## Entry format

```md
## YYYY-MM-DD — Short task title

- Scope: What changed and why.
- Files: Key files added or changed.
- Verification: Commands run and their outcome.
- Notes: Follow-up context, if needed.
```

## 2026-07-15 — Complete Phase 2 gRPC in Go learning scaffold

- Scope: Refined Phase 1 notes, added Phase 2 study notes, documented
  reproducible protobuf generation, and added an in-memory end-to-end test for
  the Hello World gRPC service.
- Files: `grpc-playground/cheatsheet.md`, `grpc-playground/Makefile`, generated
  Hello World protobuf files, and `helloworld/greeter_server/main_test.go`.
- Verification: Ran `make -C grpc-playground protos`,
  `(cd grpc-playground && go test ./...)`, and
  `(cd grpc-playground && go test -race ./...)` successfully. Re-running
  `make protos` produced identical generated files.
- Notes: `protos` is an aggregating Make target; add a dedicated target for
  each future playground example.

## 2026-07-15 — Add AI repository documentation

- Scope: Added durable guidance for AI agents, repository context, and this
  append-only task history.
- Files: `AGENTS.md`, `docs/ai/repo-context.md`, and `docs/ai/task-history.md`.
- Verification: Documentation links and module commands were checked against
  the current repository layout.
- Notes: Future implementation tasks should append an entry using the format
  above.

## 2026-07-15 — Upgrade Go directives to 1.26.4

- Scope: Updated the Go version in every `go.mod` and in `go.work`, then ran
  `go mod tidy` for each module and `go work sync` for the workspace.
- Files: `go.work`, every module's `go.mod`, and the affected `go.sum` and
  `go.work.sum` files.
- Verification: `go test ./...` passed for `grpc-playground` and
  `product-images`; the standalone book module passed with `GOWORK=off`.
- Notes: The currency test needs network access to the ECB. Product API tests
  require a local service on port 9090 and contain existing validation
  expectations that fail under the resolved dependency graph.

## 2026-07-15 — Refresh root project documentation

- Scope: Replaced the stale root README with an accurate repository guide and
  moved the gRPC learning roadmap into `docs/learning` with current Phase 1–2
  progress.
- Files: `README.md`, `docs/learning/grpc-microservices-roadmap.md`, and
  `docs/ai/repo-context.md`.
- Verification: Checked internal documentation references and ran
  `git diff --check`.
- Notes: Module-specific README files were reviewed but intentionally left for
  separate, scoped updates.

## 2026-07-15 — Add YouTube tutorial overview

- Scope: Added a tutorial-level README with component ownership, local run
  order, port mapping, protobuf generation, and known limitations.
- Files: `tutorials/building-microservices-youtube/README.md` and
  `docs/ai/task-history.md`.
- Verification: Checked service ports, frontend endpoint configuration, and
  internal links against the current source tree.
- Notes: Component-specific README files remain pending individual review.

## 2026-08-10 — Add Russian gRPC cheatsheet translation

- Scope: Added a Russian translation of the Phase 1–2 gRPC cheatsheet while
  preserving the English source document and technical identifiers.
- Files: `grpc-playground/cheatsheet.ru.md` and `docs/ai/task-history.md`.
- Verification: Checked the translated document structure against
  `grpc-playground/cheatsheet.md` and ran `git diff --check` successfully.
- Notes: Commands, code identifiers, paths, and links were retained.
