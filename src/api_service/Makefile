swagger_generate: 
	make swagger_validate
	make swagger_generate_api_server
	make swagger_generate_documentation

swagger_validate: 
	swagger validate swagger.yml

swagger_generate_api_server:
	rm -fr api
	swagger generate server --exclude-main -s api -m api/models --name Admin --target . -f swagger.yml
	
swagger_generate_documentation: 
	swagger generate markdown --output ../../docs/03_installation_usage/admin_api.md -f swagger.yml

run:
	go run ./cmd --port 8080 --host localhost 

test: 
	@echo 
	@echo 
	@echo "______________Started Unit Testing API Service_____________________"
	go test -v --coverprofile api_service_cover.out -coverpkg=./... ./...
	go tool cover --func=api_service_cover.out
	rm api_service_cover.out
	@echo "______________Finished Unit Testing API Service_____________________"
	@echo 
	@echo 


.PHONY: swagger_generate swagger_validate swagger_generate_api_server swagger_generate_documentation run
