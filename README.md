# squrl
URL Shortner Service

## Requirements

## Design

## Usage

## Tools 

### Architecture diagrams editor
Open this file, [ArchitectureDiagrams.xml](docs/design/diagrams/ArchitectureDiagrams.xml), using https://www.diagrameditor.com/

### Database diagrams editor
Open this file, [DB_Schema.txt](docs/design/diagrams/DB_Schema.txt), using https://dbdiagram.io/ 

### Sequence diagrams editor
Open this file, [SequenceDiagrams.txt](docs/design/diagrams/SequenceDiagrams.txt), using https://sequencediagram.org/

## Swagger
Swagger (GoSwagger) is used to generate the API from a specification file `swagger.yml` : Swagger specification for the URL Shortner API (this represents the core of our microservice),

### Installation

To install swagger:
```
make swagger_install
```

### Usage
To generate Order API service, Validation API client, and the documentation, run the following command:
```
make swagger_generate
```
## Database migration
This tool is used to upgrade and downgrade database schema in development environment.
[``migrate`` installation instructures](https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md)

## Sqlc
This tool is used to automatically generate data access layer code from SQL queries and database schema : 

```
go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
```
