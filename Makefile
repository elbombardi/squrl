swagger_install:
	/bin/bash swagger/install.sh

swagger_validate: 
	swagger validate swagger/api.yml

swagger_generate: 
	make swagger_generate_api_server
	make swagger_generate_documentation

swagger_generate_api_server:
	swagger generate server --exclude-main -s api --name ShortURL --target api_service -f swagger/api.yml
	
swagger_generate_documentation: 
	swagger generate markdown --output ./docs/03_installation_usage/api.md -f swagger/api.yml

sqlc:
	sqlc generate

compose:
	make compose_down
	docker-compose up --build --remove-orphans

run_postgres:
	docker run -e POSTGRES_PASSWORD=password -e POSTGRES_USER=postgres -p 5433:5432  postgres

run_api_server:
	go run ./cmd/api-server/ --port 8080 --host localhost 

run_redirection_server:
	go run ./cmd/redirection-server/ --port 8085 --host localhost

.PHONY: swagger_install swagger_validate swagger_generate_orders_server swagger_generate_validation_server swagger_generate_documentation
