.PHONY: protos

protos:
# 	protoc -I protos/ protos/currency.proto --go_out=plugins=grpc:protos/currency
	protoc -I protos \
		protos/order/order.proto \
  		--go_out=golang --go_opt=paths=source_relative \
  		--go-grpc_out=golang --go-grpc_opt=paths=source_relative
