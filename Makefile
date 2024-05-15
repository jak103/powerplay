production_image:
	docker build -f build/Dockerfile.production -t powerplay .

production_debug:
	docker build --progress=plain --no-cache -f build/Dockerfile.production -t powerplay .

run_local_image:
	docker run -p 127.0.0.1:9001:9001/tcp powerplay:latest

test:
	@echo "🚀 Testing code: Running go test inside the backend container"
	@docker-compose -f docker-compose.yml exec -T backend bash -c "cd /app/backend && go test -v ./..."