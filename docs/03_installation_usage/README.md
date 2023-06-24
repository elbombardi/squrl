# How to test the platform locally ? 
1. Start the components by typing : 
```
make compose
```

2. Orders API is accessible on : [http://[::]:8080/docs](http://[::]:8080/docs)

3. Validation API is accessible on : [http://[::]:8081/docs](http://[::]:8080/docs)

4. Mock Services & Data : 
Mock data is stored in this folder [./mock_data](./mock_data).
    - User Profiles mock data :  [./mock_data/users_mock_data.json](./mock_data/users_mock_data.json),
    - Bank Profiles mock data :  [./mock_data/bank_profile_mock_data.json](./mock_data/bank_profile_mock_data.json),
    - Bank Accounts mock data :  [./mock_data/bank_account_mock_data.json](./mock_data/bank_account_mock_data.json),

    **N.B : if you modify data in the json files, restart the docker-compose, so that the modifications take effect**

6. To try the mock services separately : 
    - User Profile API : [http://[::]:8083/docs](http://[::]:8083/docs)
    - Bank Profile API : [http://[::]:8090/docs](http://[::]:8090/docs)
    - Bank Account API : [http://[::]:8091/docs](http://[::]:8091/docs)

5. To check the database, use a database client like (https://dbeaver.io/) and connect to the postgres database : 
    - Host : *localhost*  
    - Port : *5432*
    - Database : *postgres*
    - Username : *postgres*
    - Password : *postgres*
    
6. To check Kafka : 
[http://[::]:8082/](http://[::]:8082).


# Code Structure
TODO : to be updated
```
├── cmd/
│   ├── orders-server/                          <--- Entry point for the Orders Microservice
│   ├── demo-data/                              <--- A tool to inject demo data in the development environment.
│   └── validation-server/                      <--- Entry point for the Validation Microservice
├── documentation/
│   ├── diagrams/
│   ├── images/
│   ├── api/                                    <--- (Generated with Swagger) Documentation for Service contracts
│   ├── Specification.md                        <--- Specifications
│   └── Configuration.md
├── models/                                     <-- (Generated with Swagger) Models used by the Orders API
├── api/                                        <-- (Generated with Swagger) Orders API: server code
├── validation/                                 <-- (Generated with Swagger) Validation API: client + mocked server code
├── db/                                         <-- Database code
├── handlers/                                   <-- Code of the endpoints
├── core/                                       <-- Business logic, and initialization/finalization code.
├── swagger/                                    <-- Swagger Specification for the Order API and Validation API
├── go.mod
├── go.sum
├── LICENSE
├── Makefile
├── docker-compose.yml
├── platform2.sql
├── platform.sql
└── README.md
```
