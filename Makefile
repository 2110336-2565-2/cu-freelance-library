proto:
	protoc --proto_path=pkg/proto --go_out=. --go-grpc_out=require_unimplemented_servers=false:. common.proto
