# Building Microservices in Go

Exercises based on the
[Building Microservices in Go](https://www.youtube.com/playlist?list=PLmD8u-IFdreyh6EUfevBcbiuCKzFk0EW_)
YouTube series by Nicholas Jackson. This directory preserves the tutorial's
incremental examples; it is not a single production-ready application.

## Components

| Path | Role | Default address |
| --- | --- | --- |
| [`currency/`](currency/) | gRPC exchange-rate service. It provides rates to the Product API. | `:9092` |
| [`product-api/`](product-api/) | HTTP JSON API for products; it calls the currency service over gRPC. | `:9090` |
| [`product-images/`](product-images/) | HTTP service for product-image upload and serving. | `:9091` |
| [`frontend/`](frontend/) | React interface for reading products from the Product API. | `:3001` for the current CORS configuration |

Each Go directory is its own module. Run Go commands from the component you are
working on; do not run `go test ./...` from the repository root.

## Local run order

Use separate terminals. The currency service is first because the Product API
connects to it at `localhost:9092`.

```bash
(cd currency && go run .)
(cd product-api && go run .)
```

The image service is independent of the Product API and can be run when needed:

```bash
(cd product-images && go run .)
```

The frontend reads its API and image endpoints from
[`frontend/public/global.js`](frontend/public/global.js). The Go services allow
CORS from `http://localhost:3001`, so start the Create React App dev server on
that port:

```bash
(cd frontend && PORT=3001 yarn start)
```

## Protobuf code generation

The currency service owns its protobuf contract and generated Go bindings:

```bash
make -C currency protos
```

Do not edit generated `*.pb.go` or `*_grpc.pb.go` files manually.

## Tests and known limitations

Run tests from the relevant module:

```bash
(cd currency && go test ./...)
(cd product-api && go test ./...)
(cd product-images && go test ./...)
```

- Currency tests fetch data from the European Central Bank and therefore need
  network access.
- The Product API suite includes an integration test that expects a running API
  on `localhost:9090`; its validation tests also need separate maintenance.
- The frontend Admin form currently has a hard-coded upload target on port
  `8000`, which does not match the Product Images service. Treat that upload
  flow as an unfinished tutorial example.

## Component documentation

- [`currency/README.md`](currency/README.md)
- [`product-api/README.md`](product-api/README.md)
- [`product-images/README.md`](product-images/README.md)
- [`frontend/README.md`](frontend/README.md)

Those READMEs are maintained separately because they document individual
tutorial stages and may contain historical commands.
