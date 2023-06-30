
# Build from code
## Requirment to build this code 
- Golang 1.20

## Building instructions
Step 1.1. Change directory to the root of this project:
Step 1.2. Build the API Server
Step 1.3. Build the redirection server
   
# Install
The binaries can be installed seperately, in two different machines, the only requirement is that they should be able to access the same database.

Step 2.1. Copy the binary of the API Server (short_url_api_server) to `/usr/local/bin` (or any other folder in your PATH).
Step 2.2. Copy the binary of the redirection Server (short_url_api_server) to `/usr/local/bin` (or any other folder in your PATH).
Step 2.3. Copy the redirection_404.html and redirection_500.html files to a dedicated folder (for example /opt/short_url/).

# Database preparation
Step 3.1 Create a dedicated user for the application.
Step 3.2. Create a new database in Postgres.
Step 3.3. Run the following script to create the tables: ./db/migration/000001_init_schema.up.sql
Step 3.4. Take a note of the following information: 
    - Host name : The IP adress or the hostname of the Postgres server.  
    - Port : The network port on which the Postgres server is listening (usually *5432*)
    - Database : The name of the database created in step 3.2.
    - Username : The name of the user created in step 3.1.
    - Password : The password of the user created in step 3.1.

# Configuration 
Create the following environment variables:
## Common configuration 
This is a list of common environment variables that are needed by both servers (if the servers are installed on different machines, those environment variables should be set on both machines)

| Name                           | Description                                                                                                        | Example                                                                                                               |
|--------------------------------|--------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------|
| `DB_DRIVER`                    | Should always be `postgres`.                                                                                    |                                                                                                                 |
| `DB_SOURCE`                    | URL of the PostgreSQL database                                                                                     | `postgresql://{*user name*}:{*user password*}@{*host name*}:{*port number*}/{*database name*}?{*more configuration options*}`                                                |
| `ADMIN_API_KEY`                | API key used by the admin.                                                                                          | `1234`                                                                                                                |
| `REDIRECTION_SERVER_BASE_URL`  | Base URL of the redirection server.                                                                                 | `http://localhost:8085`                                                                                              |
| `REDIRECTION_404_PAGE`         | Page shown by the redirection server when the customer or short URL is not found or disabled.                      | `redirection_404.html`                                                                                                |
| `REDIRECTION_500_PAGE`         | Page shown by the redirection server in case of an unexpected internal error.                                       | `redirection_500.html`                                                                                                |

Please ensure that you set these environment variables according to your specific setup and requirements.
- DB_DRIVER=postgres
- DB_SOURCE=<Url of Postgres database> for example `postgresql://postgres:password@localhost:5433/postgres?sslmode=disable`
- ADMIN_API_KEY=<API Key to be used by the admin> for example `1234`
- REDIRECTION_SERVER_BASE_URL=<Base URL of the redirection server> for our example `http://localhost:8085`
- REDIRECTION_404_PAGE=<Page to be shown by the redirection server if customer or short url are not found or if they are disabled>, default `redirection_404.html`
- REDIRECTION_500_PAGE=<Page to be show by the redirection server if there is an unexpected internal error>, default: redirection_500.html

# Launch

# Future development

# How to test the platform locally ? 
## 1. Go to the root of the project.
There should be the following files: 
```
   env => configuration, used in development mode
   redirection_400.html => html page to be shown by the redirection server if customer or short url are not found or if they are disabled
   redirection_500.html => html page to be show by the redirection server if there is an unexpected internal error
```

## 2. Start postgres database using Docker (on a dedicated terminal window): 
```
	docker run -e POSTGRES_PASSWORD=password -e POSTGRES_USER=postgres -p 5433:5432  postgres
```
Keep the terminal open, the database will be stopped and deleted when you close it.

## 3. Connect to the database with the tool of your choice (dbeaver, pgadmin, etc.).
    - Host : *localhost*  
    - Port : *5433*
    - Database : *postgres*
    - Username : *postgres*
    - Password : *postgres*

## 4. Initialize the database by running the script that you can find here: 
```
<project>/db/migration/000001_init_schema.up.sql
```
## 5. Start the Short API Server on port 8081 : 
```
	go run ./cmd/api-server/ --port 8081 --host localhost 
```
Then Short API server is now accessible on : [http://localhost:8081/api/docs](http://localhost:8081/api/docs)

## 6. Start the Redirection server on port 8085 : 
```
    go run ./cmd/redirection-server/ --port 8085 --host localhost 
```

## 7. Using the [Swagger UI](http://localhost:8081/docs) or using cUrl:  
### 6.1 Create a new customer (Admin API key is 1234).
### 6.2 Create a new short url for the customer, by using the API key returned by the previous call.
### 6.3 Try to access the short url generated by the previous call, you should be redirected to the long URL.

# Environment Variables
This is a list of the environment variables that can be used to configure the application.

TODO: 
- DB_DRIVER=postgres
- DB_SOURCE=<Url of Postgres database> for example `postgresql://postgres:password@localhost:5433/postgres?sslmode=disable`
- ADMIN_API_KEY=<API Key to be used by the admin> for example `1234`
- REDIRECTION_SERVER_BASE_URL=<Base URL of the redirection server> for our example `http://localhost:8085`
- REDIRECTION_404_PAGE=<Page to be shown by the redirection server if customer or short url are not found or if they are disabled>, default `redirection_404.html`
- REDIRECTION_500_PAGE=<Page to be show by the redirection server if there is an unexpected internal error>, default: redirection_500.html

# Code Structure
TODO