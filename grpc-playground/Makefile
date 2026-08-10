.PHONY: protos protos-helloworld protos-catalog-v1 protos-catalog-v2

# Regenerate protobuf and gRPC Go code for every playground example.
protos: protos-helloworld protos-catalog-v1 protos-catalog-v2

protos-helloworld:
	protoc \
		--go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		helloworld/helloworld/helloworld.proto

protos-catalog-v1:
	protoc \
		--go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		catalog/v1/catalog.proto

protos-catalog-v2:
	protoc \
		--go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		catalog/v2/catalog.proto
