include .env
export

.PHONY: run_test
run_test:
	@echo "Running test..."
	go test -v --race ./...
	@echo "Done!"

.PHONY: generate_mocks
generate_mocks:
	@echo "Generating mocks"
	mockgen -source=pkg/domain/account_repository.go -destination=pkg/adapter/mocks/account_repository_mock.go -package=mocks
	@echo "Done!"
