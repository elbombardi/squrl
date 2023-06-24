swagger_install:
	/bin/bash swagger/install.sh

swagger_validate: 
	swagger validate swagger/orders.yml

swagger_generate: 
	make swagger_generate_orders_server

swagger_generate_orders_server:
	swagger generate server --exclude-main -s api --name shorturl --target . -f swagger/api.yml
	
swagger_generate_documentation: 
	swagger generate markdown --output ./docs/02_usage/api.md -f swagger/api.yml

sqlc:
	sqlc generate

compose:
	make compose_down
	docker-compose up --build --remove-orphans

compose_down:
	docker-compose down

docker_rm: 
	docker stop squrl_db
	docker stop squrl_api_server
	docker stop squrl_redirection_server
	docker rm squrl_db
	docker rm squrl_api_server
	docker rm squrl_redirection_server

run_api_server:
	go run ./cmd/api-server/ --port 5050

run_redirection_server:
	go run ./cmd/redirection-server/ --port 5060

test:
	make test_api_service
	make test_redirection_service

test_api_service: 
	@echo 
	@echo 
	@echo "______________Started Unit Testing URL Shortner API Service_____________________"
	go test -v --coverprofile api_service_cover.out ./orders_service/...
	go tool cover --func=api_service_cover.out
	rm api_service_cover.out
	@echo "______________Finished Unit Testing API Service_____________________"
	@echo 
	@echo 

test_redirection_service: 
	@echo 
	@echo 
	@echo "______________Started Unit Testing Redirection Service_____________________"
	go test -v --coverprofile redirection_service_cover.out ./validation_service/...
	go tool cover --func=redirection_service_cover.out
	rm redirection_service_cover.out
	@echo "______________Finished Unit Testing Redirection Service_____________________"
	@echo 
	@echo 


.PHONY: swagger_install swagger_validate swagger_generate_orders_server swagger_generate_validation_server swagger_generate_documentation
