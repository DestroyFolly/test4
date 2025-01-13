run-app:
	go run cmd/main.go

docker-up:
	docker-compose up -d

run-tests: docker-up
	go test e2e/login_test.go && go test e2e/reset_test.go

.PHONY: run-app docker-up run-tests
