DB_DSN := "postgres://postgres:12345678@localhost:5432/postgres?sslmode=disable"
MIGRATE := migrate -path ./migrations -database ${DB_DSN}

migrate-new:
	migrate create -ext sql -dir ./migrations ${NAME}

migrate-up:
	${MIGRATE} up

migrate-down:
	${MIGRATE} down

run:
	go run cmd/server/main.go