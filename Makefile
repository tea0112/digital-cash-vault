run:
	env DCV_APP_PREFIX=DCV go run ./...
migrate:
	env DCV_APP_PREFIX=DCV go run ./... migrate
