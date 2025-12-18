include .env
MIGRATIONS_PATH = ./cmd/migrate/migrations

# ----------------------------
# Create a new migration
# Usage: make migrate-create migration_name
# ----------------------------
.PHONY: migrate-create
migrate-create:
	@if [ -z "$(filter-out $@,$(MAKECMDGOALS))" ]; then \
		echo "Error: migration name required"; exit 1; \
	fi
	@migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(filter-out $@,$(MAKECMDGOALS))

# ----------------------------
# Apply all pending migrations
# ----------------------------
.PHONY: migrate-up
migrate-up:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_ADDR) up

# ----------------------------
# Rollback migrations
# Usage: make migrate-down [steps]
# ----------------------------
.PHONY: migrate-down
migrate-down:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_ADDR) down $(filter-out $@,$(MAKECMDGOALS))
