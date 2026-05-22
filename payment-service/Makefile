include .env
export

migrate-up:
	migrate -path migrations -database "$(DATABASE_URL)" up

migrate-down:
	migrate -path migrations -database "$(DATABASE_URL)" down 1

migrate-down-all:
	migrate -path migrations -database "$(DATABASE_URL)" down -all

migrate-create:
	migrate create -ext sql -dir migrations -seq $(MIGRATE_V2)

migrate-status:
	migrate -path migrations -database "$(DATABASE_URL)" version

migrate-force:
	migrate -path migrations -database "$(DATABASE_URL)" force $(version)

gen:
	@protoc \
	--proto_path=protos "protos/protobuf/payment.proto" \
	--go_out=protos/gen/payment --go_opt=paths=source_relative \
	--go-grpc_out=protos/gen/payment --go-grpc_opt=paths=source_relative

	@protoc \
    	--proto_path=protos "protos/protobuf/wallet.proto" \
    	--go_out=protos/gen/wallet --go_opt=paths=source_relative \
    	--go-grpc_out=protos/gen/wallet --go-grpc_opt=paths=source_relative