.PHONY: protos protos-helloworld

# Regenerate protobuf and gRPC Go code for every playground example.
protos: protos-helloworld

protos-helloworld:
	protoc \
		--go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		helloworld/helloworld/helloworld.proto
