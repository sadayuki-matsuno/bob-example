psql:
	docker compose exec psql psql -U bob -d testdb

gooseup:
	go run github.com/pressly/goose/v3/cmd/goose@v3.13.1 -dir migration postgres "host=localhost port=5433 user=bob dbname=testdb password=test sslmode=disable" up

bob:
	(cd bobgen-psql && go run main.go)
