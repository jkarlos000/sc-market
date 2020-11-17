.PHONY: protoc

protos:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative internal/infrastructure/delivery/grpc/proto/client.proto


protoex:
	protoc internal/infrastructure/delivery/grpc/proto/$service.proto --go_out=plugins=grpc:.
