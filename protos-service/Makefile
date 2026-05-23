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