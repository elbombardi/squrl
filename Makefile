install_swagger:
	sh install_swagger.sh

install_sqlc: 
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

run_postgres:
	docker run -e POSTGRES_PASSWORD=password -e POSTGRES_USER=postgres -p 5433:5432  postgres

run_frontend_server:
	go run ./src/frontend --port 8081 --host localhost

run_redirection_service :
	go run ./src/redirection_service/cmd --port 8085 --host localhost

sqlc:
	sqlc generate

.PHONY: swagger_install swagger_validate swagger_generate_orders_server swagger_generate_validation_server swagger_generate_documentation

