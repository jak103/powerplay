.PHONY: migrate production_image, production_debug run_local_image test


migrate:
	@echo "🚀 Running app in detached mode and applying migrations"
	@docker-compose -f docker-compose.yml up -d
	@docker-compose -f docker-compose.yml exec backend bash -c "cd /app/backend && go run . -migrate" || (echo "Migration failed, halting" && exit 1)

nuke-migrations:
	@echo "💥 Truncating migrations table"
	@docker-compose -f docker-compose.yml exec database bash -c "psql -d powerplay -c 'TRUNCATE migrations';"
	@$(MAKE) migrate

production_image:
	docker build -f build/Dockerfile.production -t powerplay .

production_debug:
	docker build --progress=plain --no-cache -f build/Dockerfile.production -t powerplay .

run_local_image:
	docker run -p 127.0.0.1:9001:9001/tcp powerplay:latest

test:
	@echo "🚀 Testing code: Running go test inside the backend container"
	@docker-compose -f docker-compose.yml exec -T backend bash -c "cd /app/backend && go test -v ./..."