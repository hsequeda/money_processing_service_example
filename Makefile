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

# database
.PHONY: migration_new
migration_new:
	sql-migrate new -config=sql/dbconfig.yml $(name) 

.PHONY: migration_run
migration_run:
	sql-migrate up -config=sql/dbconfig.yml

.PHONY: migration_revert
migration_revert:
	sql-migrate down -config=sql/dbconfig.yml

.PHONY: migration_status
migration_status:
	sql-migrate status -config=sql/dbconfig.yml
