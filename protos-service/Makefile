#include .env
export

gen:
	@protoc \
	--proto_path=protos "protos/protobuf/sso.proto" \
	--go_out=protos/gen/sso --go_opt=paths=source_relative \
	--go-grpc_out=protos/gen/sso --go-grpc_opt=paths=source_relative

	@protoc \
    	--proto_path=protos "protos/protobuf/token.proto" \
    	--go_out=protos/gen/token --go_opt=paths=source_relative \
    	--go-grpc_out=protos/gen/token --go-grpc_opt=paths=source_relative

	@protoc \
    	--proto_path=protos "protos/protobuf/payment.proto" \
    	--go_out=protos/gen/payment --go_opt=paths=source_relative \
    	--go-grpc_out=protos/gen/payment --go-grpc_opt=paths=source_relative

	@protoc \
		--proto_path=protos "protos/protobuf/wallet.proto" \
		--go_out=protos/gen/wallet --go_opt=paths=source_relative \
		--go-grpc_out=protos/gen/wallet --go-grpc_opt=paths=source_relative

	@protoc \
		--proto_path=protos "protos/protobuf/file.proto" \
		--go_out=protos/gen/file --go_opt=paths=source_relative \
		--go-grpc_out=protos/gen/file --go-grpc_opt=paths=source_relative

	@protoc \
    		--proto_path=protos "protos/protobuf/chat.proto" \
    		--go_out=protos/gen/chat --go_opt=paths=source_relative \
    		--go-grpc_out=protos/gen/chat --go-grpc_opt=paths=source_relative

	@protoc \
    		--proto_path=protos "protos/protobuf/match.proto" \
    		--go_out=protos/gen/match --go_opt=paths=source_relative \
    		--go-grpc_out=protos/gen/match --go-grpc_opt=paths=source_relative

	@protoc \
    		--proto_path=protos "protos/protobuf/notification.proto" \
    		--go_out=protos/gen/notification --go_opt=paths=source_relative \
    		--go-grpc_out=protos/gen/notification --go-grpc_opt=paths=source_relative
