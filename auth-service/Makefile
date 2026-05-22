include .env
export

migrate-up:
	migrate -path migrations -database "$(DATABASE_URL)" up

migrate-down:
	migrate -path migrations -database "$(DATABASE_URL)" down 1

migrate-down-all:
	migrate -path migrations -database "$(DATABASE_URL)" down -all

migrate-create:
	migrate create -ext sql -dir migrations -seq $(MIGRATION_V2)

migrate-status:
	migrate -path migrations -database "$(DATABASE_URL)" version

migrate-force:
	migrate -path migrations -database "$(DATABASE_URL)" force $(version)

gen:
	@protoc \
	--proto_path=protos "protos/protobuf/sso.proto" \
	--go_out=protos/gen/sso --go_opt=paths=source_relative \
	--go-grpc_out=protos/gen/sso --go-grpc_opt=paths=source_relative

	@protoc \
    	--proto_path=protos "protos/protobuf/token.proto" \
    	--go_out=protos/gen/token --go_opt=paths=source_relative \
    	--go-grpc_out=protos/gen/token --go-grpc_opt=paths=source_relative
