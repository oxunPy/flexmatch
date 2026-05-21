include .env
export

migrate-up:
	migrate -path migrations -database "$(DATABASE_URL)" up

migrate-down:
	migrate -path migrations -database "$(DATABASE_URL)" down 1

migrate-down-all:
	migrate -path migrations -database "$(DATABASE_URL)" down -all

migrate-create:
	migrate create -ext sql -dir migrations -seq $(MIGRATION_V3)

migrate-status:
	migrate -path migrations -database "$(DATABASE_URL)" version

migrate-force:
	migrate -path migrations -database "$(DATABASE_URL)" force 3