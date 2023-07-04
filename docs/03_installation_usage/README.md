
# I. Build from code
## Requirment to build this code 
- Golang 1.20

## Building instructions
- **Step 1.1.** Change directory to the root of this project
```bash
cd <root_folder>
```
- **Step 1.2.** Build the API Server
```bash
go build -o build/short_url_api_server ./cmd/api-server
```

- **Step 1.3.** Build the redirection server
```bash
go build -o build/short_url_redirection_server ./cmd/redirection-server
```

# II. Install
The binaries can be installed seperately, in two different machines, the only requirement is that they should be able to access the same database.

- **Step 2.1.** Copy the binary of the API Server (`build/short_url_api_server`) to `/usr/local/bin` (or any other folder in your `PATH`).
```bash
sudo cp build/short_url_api_server /usr/local/bin
```

- **Step 2.2.** Copy the binary of the redirection Server (`build/short_url_redirection_server`) to `/usr/local/bin` (or any other folder in your `PATH`).
```bash
sudo cp build/short_url_redirection_server /usr/local/bin
```

- **Step 2.3.** Copy the `redirection_404.html` and `redirection_500.html` files to a dedicated folder (for example `/opt/short_url/`).
```bash
sudo mkdir /opt/short_url
sudo cp redirection_404.html /opt/short_url/
sudo cp redirection_500.html /opt/short_url/
```

# III. Database preparation
- **Step 3.1.** Create a dedicated user for the application.
```bash
sudo -u postgres createuser <username>
```

- **Step 3.2.** Create a new database in Postgres.
```bash
sudo -u postgres createdb <dbname>
```
- **Step 3.3.** Give the user a password.
```bash
sudo -u postgres psql
psql=# alter user <username> with encrypted password '<password>';
```

- **Step 3.4.** Grant the user access to the database.
```bash
psql=# grant all privileges on database <dbname> to <username>;
```

- **Step 3.5.** Run the following script to create the tables: `./db/migration/000001_init_schema.up.sql`
- **Step 3.6.** Take a note of the following information: 
    - Hostname : The IP adress or the hostname of the Postgres server.  
    - Port : The network port on which the Postgres server is listening (usually *5432*)
    - Database : The name of the database created in step 3.2.
    - Username : The name of the user created in step 3.1.
    - Password : The password of the user created in step 3.3.

# IV. Configuration 
Create the following environment variables:

Please ensure that you set the required environment variables accordingly, while the optional ones can be adjusted as per your specific needs. 
The default values, if applicable, will be used when the optional variables are not explicitly provided.

## Common configuration 
This is a list of common environment variables that are used by both servers (if the servers are installed on different machines, those environment variables should be set on both machines)

| Name                        | Description                                                  | Required/Optional | Default Value | Example                                                                  |
|-----------------------------|--------------------------------------------------------------|-------------------|---------------|--------------------------------------------------------------------------|
| `DB_DRIVER`                 | Database driver name.                                        | Required          |               | `postgres`                                                               |
| `DB_SOURCE`                 | URL of the PostgreSQL database. More detail about the format of this parameter can be found here: https://pkg.go.dev/github.com/lib/pq#hdr-Connection_String_Parameters                              | Required          |               | `postgresql://postgres:password@localhost:5433/postgres?sslmode=disable`1 |
| `DB_MAX_IDLE_CONNS`         | Maximum number of idle connections in the connection pool.    | Optional          | 5             | `5`                                                                      |
| `DB_MAX_OPEN_CONNS`         | Maximum number of open connections in the connection pool.    | Optional          | 10            | `10`                                                                     |
| `DB_CONN_MAX_IDLE_TIME`     | Maximum time (in minutes) a connection can be idle.           | Optional          | 1             | `1`                                                                      |
| `DB_CONN_MAX_LIFE_TIME`     | Maximum time (in minutes) a connection can be kept open.      | Optional          | 30            | `30`                                                                     |

## API Server configuration
This is a list of environment variables that are used by the API Server only.
| Name                        | Description                                                  | Required/Optional | Default Value | Example                                                                  |
|-----------------------------|--------------------------------------------------------------|-------------------|---------------|--------------------------------------------------------------------------|
| `ADMIN_API_KEY`             | API key used by the admin.                                   | Required          |               | `1234`                                                                   |
| `REDIRECTION_SERVER_BASE_URL` | Base URL of the redirection server, this must be a public accessible URL. If the redirection server is deployed behind a reverse proxy, use the URL of the reverse proxy and not the local address of the application server.                            | Required          |               | `https://domain.name`                                                  |

## Redirection Server configuration
This is a list of environment variables that are used by the Redirection Server only.
| Name                        | Description                                                  | Required/Optional | Default Value | Example                                                                  |
|-----------------------------|--------------------------------------------------------------|-------------------|---------------|--------------------------------------------------------------------------|
| `REDIRECTION_404_PAGE`      | Path to the 404 error page for the redirection server.       | Required          |               | `/opt/short_url/redirection_404.html`                                    |
| `REDIRECTION_500_PAGE`      | Path to the 500 error page for the redirection server.       | Required          |               | `/opt/short_url/redirection_500.html`                                    |

# V. Launch

## API Server
Launch the API Server using the following command, by specifyging port and host:

```bash
short_url_api_server --port 8080 --host localhost 
```

### Parameters : 
- `--port` : The port on which the API Server will listen to.
- `--host` : The hostname or IP address of the host machine.

### API Documentation: 
A detailed documentation of the API can be found in [the annexe](#vi-annexe-short-url-api-documentation)

## Redirection Server
Launch the Redirection Server using the following command, by specifyging port and host:

```bash
short_url_redirection_server --port 8085 --host localhost 
```
### Parameters : 
- `--port` : The port on which the Redirection Server will listen to.
- `--host` : The hostname or IP address of the host machine.

## VI. Annexe: Short URL API Documentation


## Informations

### Version

1.0.0

## Content negotiation

### URI Schemes
  * http

### Consumes
  * application/json

### Produces
  * application/json

## All endpoints

###  operations

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /api/customer | [post customer](#post-customer) | Create Customer |
| POST | /api/short-url | [post short URL](#post-short-url) | Create ShortURL |
| PUT | /api/customer | [put customer](#put-customer) | Update Customer |
| PUT | /api/short-url | [put short URL](#put-short-url) | Update ShortURL |
  


## Paths

### <span id="post-customer"></span> Create Customer (*PostCustomer*)

```
POST /api/customer
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| X-API-KEY | `header` | string | `string` |  | ✓ |  | The admin API key. |
| customer | `body` | [PostCustomerBody](#post-customer-body) | `PostCustomerBody` | | ✓ | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#post-customer-200) | OK | Success |  | [schema](#post-customer-200-schema) |
| [400](#post-customer-400) | Bad Request | Bad Request |  | [schema](#post-customer-400-schema) |
| [401](#post-customer-401) | Unauthorized | Unauthorized |  | [schema](#post-customer-401-schema) |
| [500](#post-customer-500) | Internal Server Error | Internal Server Error |  | [schema](#post-customer-500-schema) |

#### Responses


##### <span id="post-customer-200"></span> 200 - Success
Status: OK

###### <span id="post-customer-200-schema"></span> Schema
   
  

[PostCustomerOKBody](#post-customer-o-k-body)

##### <span id="post-customer-400"></span> 400 - Bad Request
Status: Bad Request

###### <span id="post-customer-400-schema"></span> Schema
   
  

[PostCustomerBadRequestBody](#post-customer-bad-request-body)

##### <span id="post-customer-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="post-customer-401-schema"></span> Schema
   
  

[PostCustomerUnauthorizedBody](#post-customer-unauthorized-body)

##### <span id="post-customer-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="post-customer-500-schema"></span> Schema
   
  

[PostCustomerInternalServerErrorBody](#post-customer-internal-server-error-body)

###### Inlined models

**<span id="post-customer-bad-request-body"></span> PostCustomerBadRequestBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| error | string| `string` |  | |  |  |



**<span id="post-customer-body"></span> PostCustomerBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| email | string| `string` | ✓ | |  |  |
| username | string| `string` | ✓ | |  |  |



**<span id="post-customer-internal-server-error-body"></span> PostCustomerInternalServerErrorBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| error | string| `string` |  | |  |  |



**<span id="post-customer-o-k-body"></span> PostCustomerOKBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| api_key | string| `string` |  | |  |  |
| prefix | string| `string` |  | |  |  |



**<span id="post-customer-unauthorized-body"></span> PostCustomerUnauthorizedBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| error | string| `string` |  | |  |  |



### <span id="post-short-url"></span> Create ShortURL (*PostShortURL*)

```
POST /api/short-url
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| X-API-KEY | `header` | string | `string` |  | ✓ |  | The customer API key. |
| body | `body` | [PostShortURLBody](#post-short-url-body) | `PostShortURLBody` | | ✓ | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#post-short-url-200) | OK | Success |  | [schema](#post-short-url-200-schema) |
| [400](#post-short-url-400) | Bad Request | Bad Request |  | [schema](#post-short-url-400-schema) |
| [401](#post-short-url-401) | Unauthorized | Unauthorized |  | [schema](#post-short-url-401-schema) |
| [500](#post-short-url-500) | Internal Server Error | Internal Server Error |  | [schema](#post-short-url-500-schema) |

#### Responses


##### <span id="post-short-url-200"></span> 200 - Success
Status: OK

###### <span id="post-short-url-200-schema"></span> Schema
   
  

[PostShortURLOKBody](#post-short-url-o-k-body)

##### <span id="post-short-url-400"></span> 400 - Bad Request
Status: Bad Request

###### <span id="post-short-url-400-schema"></span> Schema
   
  

[PostShortURLBadRequestBody](#post-short-url-bad-request-body)

##### <span id="post-short-url-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="post-short-url-401-schema"></span> Schema
   
  

[PostShortURLUnauthorizedBody](#post-short-url-unauthorized-body)

##### <span id="post-short-url-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="post-short-url-500-schema"></span> Schema
   
  

[PostShortURLInternalServerErrorBody](#post-short-url-internal-server-error-body)

###### Inlined models

**<span id="post-short-url-bad-request-body"></span> PostShortURLBadRequestBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| error | string| `string` |  | |  |  |



**<span id="post-short-url-body"></span> PostShortURLBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| long_url | string| `string` | ✓ | |  |  |



**<span id="post-short-url-internal-server-error-body"></span> PostShortURLInternalServerErrorBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| error | string| `string` |  | |  |  |



**<span id="post-short-url-o-k-body"></span> PostShortURLOKBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| short_url | string| `string` |  | |  |  |
| short_url_key | string| `string` |  | |  |  |



**<span id="post-short-url-unauthorized-body"></span> PostShortURLUnauthorizedBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| error | string| `string` |  | |  |  |



### <span id="put-customer"></span> Update Customer (*PutCustomer*)

```
PUT /api/customer
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| X-API-KEY | `header` | string | `string` |  | ✓ |  | The admin API key. |
| body | `body` | [PutCustomerBody](#put-customer-body) | `PutCustomerBody` | | ✓ | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#put-customer-200) | OK | Success |  | [schema](#put-customer-200-schema) |
| [400](#put-customer-400) | Bad Request | Bad Request |  | [schema](#put-customer-400-schema) |
| [401](#put-customer-401) | Unauthorized | Unauthorized |  | [schema](#put-customer-401-schema) |
| [404](#put-customer-404) | Not Found | Not Found |  | [schema](#put-customer-404-schema) |
| [500](#put-customer-500) | Internal Server Error | Internal Server Error |  | [schema](#put-customer-500-schema) |

#### Responses


##### <span id="put-customer-200"></span> 200 - Success
Status: OK

###### <span id="put-customer-200-schema"></span> Schema
   
  

[PutCustomerOKBody](#put-customer-o-k-body)

##### <span id="put-customer-400"></span> 400 - Bad Request
Status: Bad Request

###### <span id="put-customer-400-schema"></span> Schema
   
  

[PutCustomerBadRequestBody](#put-customer-bad-request-body)

##### <span id="put-customer-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="put-customer-401-schema"></span> Schema
   
  

[PutCustomerUnauthorizedBody](#put-customer-unauthorized-body)

##### <span id="put-customer-404"></span> 404 - Not Found
Status: Not Found

###### <span id="put-customer-404-schema"></span> Schema
   
  

[PutCustomerNotFoundBody](#put-customer-not-found-body)

##### <span id="put-customer-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="put-customer-500-schema"></span> Schema
   
  

[PutCustomerInternalServerErrorBody](#put-customer-internal-server-error-body)

###### Inlined models

**<span id="put-customer-bad-request-body"></span> PutCustomerBadRequestBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| error | string| `string` |  | |  |  |



**<span id="put-customer-body"></span> PutCustomerBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| status | string| `string` | ✓ | |  |  |
| username | string| `string` | ✓ | |  |  |



**<span id="put-customer-internal-server-error-body"></span> PutCustomerInternalServerErrorBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| error | string| `string` |  | |  |  |



**<span id="put-customer-not-found-body"></span> PutCustomerNotFoundBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| error | string| `string` |  | |  |  |



**<span id="put-customer-o-k-body"></span> PutCustomerOKBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| status | string| `string` |  | |  |  |



**<span id="put-customer-unauthorized-body"></span> PutCustomerUnauthorizedBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| error | string| `string` |  | |  |  |



### <span id="put-short-url"></span> Update ShortURL (*PutShortURL*)

```
PUT /api/short-url
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| X-API-KEY | `header` | string | `string` |  | ✓ |  | The customer API key. |
| body | `body` | [PutShortURLBody](#put-short-url-body) | `PutShortURLBody` | | ✓ | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#put-short-url-200) | OK | Success |  | [schema](#put-short-url-200-schema) |
| [400](#put-short-url-400) | Bad Request | Bad Request |  | [schema](#put-short-url-400-schema) |
| [401](#put-short-url-401) | Unauthorized | Unauthorized |  | [schema](#put-short-url-401-schema) |
| [404](#put-short-url-404) | Not Found | Not Found |  | [schema](#put-short-url-404-schema) |
| [500](#put-short-url-500) | Internal Server Error | Internal Server Error |  | [schema](#put-short-url-500-schema) |

#### Responses


##### <span id="put-short-url-200"></span> 200 - Success
Status: OK

###### <span id="put-short-url-200-schema"></span> Schema
   
  

[PutShortURLOKBody](#put-short-url-o-k-body)

##### <span id="put-short-url-400"></span> 400 - Bad Request
Status: Bad Request

###### <span id="put-short-url-400-schema"></span> Schema
   
  

[PutShortURLBadRequestBody](#put-short-url-bad-request-body)

##### <span id="put-short-url-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="put-short-url-401-schema"></span> Schema
   
  

[PutShortURLUnauthorizedBody](#put-short-url-unauthorized-body)

##### <span id="put-short-url-404"></span> 404 - Not Found
Status: Not Found

###### <span id="put-short-url-404-schema"></span> Schema
   
  

[PutShortURLNotFoundBody](#put-short-url-not-found-body)

##### <span id="put-short-url-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="put-short-url-500-schema"></span> Schema
   
  

[PutShortURLInternalServerErrorBody](#put-short-url-internal-server-error-body)

###### Inlined models

**<span id="put-short-url-bad-request-body"></span> PutShortURLBadRequestBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| error | string| `string` |  | |  |  |



**<span id="put-short-url-body"></span> PutShortURLBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| new_long_url | string| `string` |  | |  |  |
| short_url_key | string| `string` | ✓ | |  |  |
| status | string| `string` |  | |  |  |
| tracking_status | string| `string` |  | |  |  |



**<span id="put-short-url-internal-server-error-body"></span> PutShortURLInternalServerErrorBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| error | string| `string` |  | |  |  |



**<span id="put-short-url-not-found-body"></span> PutShortURLNotFoundBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| error | string| `string` |  | |  |  |



**<span id="put-short-url-o-k-body"></span> PutShortURLOKBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| long_url | string| `string` |  | |  |  |
| status | string| `string` |  | |  |  |
| tracking_status | string| `string` |  | |  |  |



**<span id="put-short-url-unauthorized-body"></span> PutShortURLUnauthorizedBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| error | string| `string` |  | |  |  |
