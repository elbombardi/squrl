install_swagger:
	sh install_swagger.sh

install_sqlc: 
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

run_postgres:
	docker run -e POSTGRES_PASSWORD=postgres -e POSTGRES_USER=postgres -p 5432:5432  postgres

run_frontend_server:
	go run ./src/frontend --port 8081 --host localhost

run_redirection_service :
	go run ./src/redirection_service/cmd --port 8085 --host localhost

sqlc:
	sqlc generate

docker_build:
	docker build . -f src/api_service/Dockerfile -t squrl-api-service
	docker build . -f src/redirection_service/Dockerfile -t squrl-redirection-service

docker_clean: 
	docker stop squrl-db
	docker stop squrl-redirection-service
	docker stop squrl-api-service
	docker rm squrl-db
	docker rm squrl-redirection-service
	docker rm squrl-api-service

compose:
	make compose_down
	docker-compose up --build 
	cd ../..

compose_down:
	docker-compose down

.PHONY: swagger_install swagger_validate swagger_generate_orders_server swagger_generate_validation_server swagger_generate_documentation

