run:
	go run ./cmd --port 8085 --host localhost 

test: 
	@echo 
	@echo 
	@echo "______________Started Unit Testing API Service_____________________"
	go test -v --coverprofile orders_service_cover.out -coverpkg=./... ./...
	go tool cover --func=orders_service_cover.out
	rm orders_service_cover.out
	@echo "______________Finished Unit Testing API Service_____________________"
	@echo 
	@echo 

.PHONY: run

