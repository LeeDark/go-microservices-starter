# AI task history

This is an append-only record of completed implementation work. Add the newest entry at the end.
Keep entries factual and concise.

## Entry format

```md
## YYYY-MM-DD — Short task title

- Scope: What changed and why.
- Files: Key files added or changed.
- Verification: Commands run and their outcome.
- Notes: Follow-up context, if needed.
```

## 2026-07-15 — Complete Phase 2 gRPC in Go learning scaffold

- Scope: Refined Phase 1 notes, added Phase 2 study notes, documented reproducible protobuf
  generation, and added an in-memory end-to-end test for the Hello World gRPC service.
- Files: `grpc-playground/cheatsheet.md`, `grpc-playground/Makefile`, generated Hello World protobuf
  files, and `helloworld/greeter_server/main_test.go`.
- Verification: Ran `make -C grpc-playground protos`, `(cd grpc-playground && go test ./...)`, and
  `(cd grpc-playground && go test -race ./...)` successfully. Re-running `make protos` produced
  identical generated files.
- Notes: `protos` is an aggregating Make target; add a dedicated target for each future playground
  example.

## 2026-07-15 — Add AI repository documentation

- Scope: Added durable guidance for AI agents, repository context, and this append-only task
  history.
- Files: `AGENTS.md`, `docs/ai/repo-context.md`, and `docs/ai/task-history.md`.
- Verification: Documentation links and module commands were checked against the current repository
  layout.
- Notes: Future implementation tasks should append an entry using the format above.

## 2026-07-15 — Upgrade Go directives to 1.26.4

- Scope: Updated the Go version in every `go.mod` and in `go.work`, then ran `go mod tidy` for each
  module and `go work sync` for the workspace.
- Files: `go.work`, every module's `go.mod`, and the affected `go.sum` and `go.work.sum` files.
- Verification: `go test ./...` passed for `grpc-playground` and `product-images`; the standalone
  book module passed with `GOWORK=off`.
- Notes: The currency test needs network access to the ECB. Product API tests require a local
  service on port 9090 and contain existing validation expectations that fail under the resolved
  dependency graph.

## 2026-07-15 — Refresh root project documentation

- Scope: Replaced the stale root README with an accurate repository guide and moved the gRPC
  learning roadmap into `docs/learning` with current Phase 1–2 progress.
- Files: `README.md`, `docs/learning/grpc-microservices-roadmap.md`, and `docs/ai/repo-context.md`.
- Verification: Checked internal documentation references and ran `git diff --check`.
- Notes: Module-specific README files were reviewed but intentionally left for separate, scoped
  updates.

## 2026-07-15 — Add YouTube tutorial overview

- Scope: Added a tutorial-level README with component ownership, local run order, port mapping,
  protobuf generation, and known limitations.
- Files: `tutorials/building-microservices-youtube/README.md` and `docs/ai/task-history.md`.
- Verification: Checked service ports, frontend endpoint configuration, and internal links against
  the current source tree.
- Notes: Component-specific README files remain pending individual review.

## 2026-08-10 — Add Russian gRPC cheatsheet translation

- Scope: Added a Russian translation of the Phase 1–2 gRPC cheatsheet while preserving the English
  source document and technical identifiers.
- Files: `grpc-playground/cheatsheet.ru.md` and `docs/ai/task-history.md`.
- Verification: Checked the translated document structure against `grpc-playground/cheatsheet.md`
  and ran `git diff --check` successfully.
- Notes: Commands, code identifiers, paths, and links were retained.

## 2026-08-10 — Add Russian gRPC roadmap translation

- Scope: Added a Russian translation of the complete gRPC and microservices learning roadmap without
  changing the English source roadmap.
- Files: `docs/learning/grpc-microservices-roadmap.ru.md` and `docs/ai/task-history.md`.
- Verification: Checked that all roadmap phases, checklists, project sequences, links, and the
  Mermaid learning map were retained; `git diff --check` passed.
- Notes: Technical identifiers, commands, project names, and URLs were kept in their original form
  where appropriate.

## 2026-08-10 — Integrate Buf learning track

- Scope: Integrated the full Buf learning path across the existing gRPC roadmap, from local CLI and
  contract checks through CI, BSR, and advanced governance.
- Files: Both roadmap translations, both gRPC cheatsheets, and this task history.
- Verification: Checked the English and Russian additions, preserved the `protoc`/Makefile baseline,
  and ran `git diff --check` successfully.
- Notes: No Buf configuration or repository workflow was implemented; this change updates the
  learning plan only.

## 2026-08-10 — Add adjacent gRPC tooling tracks

- Scope: Extended both roadmaps and cheatsheets with ConnectRPC, grpc-gateway, OpenAPI,
  OpenTelemetry, and grpcurl as connected learning tracks around the existing `grpc-go` and Buf
  path.
- Files: Both roadmap translations, both gRPC cheatsheets, and this task history.
- Verification: Checked phase placement, bilingual structure, retained the `protoc`/Makefile
  baseline, and ran `git diff --check` successfully.
- Notes: Documentation only; no runtime dependencies, generated code, gateway, telemetry, or CLI
  configuration was added.

## 2026-08-10 — Split Phase 3 into 3A, 3B, and 3C

- Scope: Reorganized Phase 3 into protobuf contract design, Buf workflow, and adjacent protobuf/gRPC
  tooling, with explicit goals and Definition of Done.
- Files: Both roadmap translations, both gRPC cheatsheets, and this task history.
- Verification: Checked bilingual phase structure, learning order, tool boundaries, and ran `git
  diff --check` successfully.
- Notes: The primary `grpc-go + protobuf + Buf` path and `protoc`/Makefile baseline remain
  unchanged.

## 2026-08-10 — Add Phase 3A compatibility rules to cheatsheets

- Scope: Added a concise bilingual reference for protobuf field numbers, compatible additions,
  breaking changes, and `reserved` fields.
- Files: `grpc-playground/cheatsheet.md`, `grpc-playground/cheatsheet.ru.md`, and
  `docs/ai/task-history.md`.
- Verification: Checked both sections and ran `git diff --check` successfully.
- Notes: Detailed `v1`/`v2` implementation remains a separate Phase 3A task.

## 2026-08-10 — Document catalog v1 and v2 design

- Scope: Added a pre-test Phase 3A cheatsheet description of the current `catalog.v1` and
  `catalog.v2` contract design and compatible field additions.
- Files: `grpc-playground/cheatsheet.md` and `docs/ai/task-history.md`.
- Verification: Checked the description against both `.proto` files and the current Makefile
  targets; no generated code or tests were changed.
- Notes: The Russian cheatsheet remains unchanged until the English description is reviewed.

## 2026-08-10 — Add catalog focused compatibility tests

- Scope: Added contract-focused tests for `catalog.v1` and `catalog.v2` message round-trip and
  bidirectional wire compatibility.
- Files: `grpc-playground/catalog/v1/catalog_test.go` and this task history.
- Verification: Focused and full `grpc-playground` tests are run during handoff.
- Notes: Tests do not start a server or require external services.

## 2026-08-10 — Document catalog focused tests in cheatsheets

- Scope: Added synchronized English and Russian cheatsheet sections describing contract round-trip
  and `v1`/`v2` wire-compatibility tests.
- Files: `grpc-playground/cheatsheet.md`, `grpc-playground/cheatsheet.ru.md`, and this task history.
- Verification: Checked the commands and test scope against the implemented
  `catalog/v1/catalog_test.go`; `git diff --check` passed.
- Notes: Documentation only; test code was not changed.

## 2026-08-10 — Close Phase 3A protobuf contract design

- Scope: Updated bilingual roadmap and cheatsheets with Phase 3A completion status, review evidence,
  and transition to Phase 3B.
- Files: Both roadmap translations, both gRPC cheatsheets, and this task history.
- Verification: Ran `make -C grpc-playground protos`, `GOCACHE=/tmp/go-microservices-starter-gocache
  go test ./catalog/...`, `GOCACHE=/tmp/go-microservices-starter-gocache go test ./...`, and `git
  diff --check` successfully.
- Notes: Phase 3A is closed; Buf work begins in Phase 3B.
