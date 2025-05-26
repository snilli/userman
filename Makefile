total-coverage:
	go test $$(go list ./... | grep -v mock | grep -v cmd | grep -v proto | grep -v http/handler | grep -v http/middleware) -coverprofile=coverage.out ./... && go tool cover -func=coverage.out | tail -1

migrate-db:
	migrate -database mongodb://root:L1ttleK1tten@localhost:27017/main_db\?authSource=admin -path db/migrations up

gen-mock:
	mockery

run:
	go run cmd/http/main.go