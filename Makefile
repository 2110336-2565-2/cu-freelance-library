proto:
	protoc --proto_path=pkg/proto --go_out=. --go-grpc_out=require_unimplemented_servers=false:. common.proto
	protoc --proto_path=pkg/proto --go_out=. --go-grpc_out=require_unimplemented_servers=false:. user.proto
	protoc --proto_path=pkg/proto --go_out=. --go-grpc_out=require_unimplemented_servers=false:. portfolio.proto
	protoc --proto_path=pkg/proto --go_out=. --go-grpc_out=require_unimplemented_servers=false:. auth.proto
	protoc --proto_path=pkg/proto --go_out=. --go-grpc_out=require_unimplemented_servers=false:. search.proto
	protoc --proto_path=pkg/proto --go_out=. --go-grpc_out=require_unimplemented_servers=false:. suggest.proto
	protoc --proto_path=pkg/proto --go_out=. --go-grpc_out=require_unimplemented_servers=false:. order.proto
	protoc --proto_path=pkg/proto --go_out=. --go-grpc_out=require_unimplemented_servers=false:. orderTemplate.proto
	protoc --proto_path=pkg/proto --go_out=. --go-grpc_out=require_unimplemented_servers=false:. request.proto
	protoc --proto_path=pkg/proto --go_out=. --go-grpc_out=require_unimplemented_servers=false:. file.proto
	protoc --proto_path=pkg/proto --go_out=. --go-grpc_out=require_unimplemented_servers=false:. issue.proto
	protoc --proto_path=pkg/proto --go_out=. --go-grpc_out=require_unimplemented_servers=false:. chat.proto

test:
	go test -v -coverpkg . -coverprofile coverage.out -covermode count .