


# SQURL - ADMIN API
  

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

## Access control

### Security Schemes

#### Bearer (header: Authorization)



> **Type**: apikey

## All endpoints

###  accounts

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /v1/accounts | [create account](#create-account) | Create an account |
| PUT | /v1/accounts | [update account](#update-account) | Update an account |
  


###  general

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /v1/health | [healthcheck](#healthcheck) | Healthcheck |
| POST | /v1/login | [login](#login) | Login |
  


###  urls

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /v1/urls | [create URL](#create-url) | Create a new URL |
| PUT | /v1/urls | [update URL](#update-url) | Update a URL |
  


## Paths

### <span id="create-account"></span> Create an account (*CreateAccount*)

```
POST /v1/accounts
```

Create a new account

#### Security Requirements
  * Bearer

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| Authorization | `header` | string | `string` |  | ✓ |  | Bearer <JWT Token> |
| account | `body` | [Account](#account) | `models.Account` | | ✓ | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#create-account-200) | OK | Success |  | [schema](#create-account-200-schema) |
| [400](#create-account-400) | Bad Request | Bad Request |  | [schema](#create-account-400-schema) |
| [401](#create-account-401) | Unauthorized | Unauthorized |  | [schema](#create-account-401-schema) |
| [500](#create-account-500) | Internal Server Error | Internal Server Error |  | [schema](#create-account-500-schema) |

#### Responses


##### <span id="create-account-200"></span> 200 - Success
Status: OK

###### <span id="create-account-200-schema"></span> Schema
   
  

[AccountCreated](#account-created)

##### <span id="create-account-400"></span> 400 - Bad Request
Status: Bad Request

###### <span id="create-account-400-schema"></span> Schema
   
  

[Error](#error)

##### <span id="create-account-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="create-account-401-schema"></span> Schema
   
  

[Error](#error)

##### <span id="create-account-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="create-account-500-schema"></span> Schema
   
  

[Error](#error)

### <span id="create-url"></span> Create a new URL (*CreateURL*)

```
POST /v1/urls
```

Create a new URL

#### Security Requirements
  * Bearer

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| Authorization | `header` | string | `string` |  | ✓ |  | Bearer <JWT Token> |
| body | `body` | [URL](#url) | `models.URL` | | ✓ | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#create-url-200) | OK | Success |  | [schema](#create-url-200-schema) |
| [400](#create-url-400) | Bad Request | Bad Request |  | [schema](#create-url-400-schema) |
| [401](#create-url-401) | Unauthorized | Unauthorized |  | [schema](#create-url-401-schema) |
| [500](#create-url-500) | Internal Server Error | Internal Server Error |  | [schema](#create-url-500-schema) |

#### Responses


##### <span id="create-url-200"></span> 200 - Success
Status: OK

###### <span id="create-url-200-schema"></span> Schema
   
  

[URLCreated](#url-created)

##### <span id="create-url-400"></span> 400 - Bad Request
Status: Bad Request

###### <span id="create-url-400-schema"></span> Schema
   
  

[Error](#error)

##### <span id="create-url-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="create-url-401-schema"></span> Schema
   
  

[Error](#error)

##### <span id="create-url-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="create-url-500-schema"></span> Schema
   
  

[Error](#error)

### <span id="healthcheck"></span> Healthcheck (*Healthcheck*)

```
GET /v1/health
```

Health check endpoint

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#healthcheck-200) | OK | Ok |  | [schema](#healthcheck-200-schema) |

#### Responses


##### <span id="healthcheck-200"></span> 200 - Ok
Status: OK

###### <span id="healthcheck-200-schema"></span> Schema
   
  



### <span id="login"></span> Login (*Login*)

```
POST /v1/login
```

Returns JWT token for authorized user

#### Consumes
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| login | `body` | [LoginInfo](#login-info) | `models.LoginInfo` | | ✓ | | Login Payload |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#login-200) | OK | Successful login |  | [schema](#login-200-schema) |
| [400](#login-400) | Bad Request | Bad Request |  | [schema](#login-400-schema) |
| [401](#login-401) | Unauthorized | Unauthorized |  | [schema](#login-401-schema) |
| [500](#login-500) | Internal Server Error | Internal Server Error |  | [schema](#login-500-schema) |

#### Responses


##### <span id="login-200"></span> 200 - Successful login
Status: OK

###### <span id="login-200-schema"></span> Schema
   
  

[LoginSuccess](#login-success)

##### <span id="login-400"></span> 400 - Bad Request
Status: Bad Request

###### <span id="login-400-schema"></span> Schema
   
  

[Error](#error)

##### <span id="login-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="login-401-schema"></span> Schema
   
  

[Error](#error)

##### <span id="login-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="login-500-schema"></span> Schema
   
  

[Error](#error)

### <span id="update-account"></span> Update an account (*UpdateAccount*)

```
PUT /v1/accounts
```

Update an account

#### Security Requirements
  * Bearer

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| Authorization | `header` | string | `string` |  | ✓ |  | Bearer <JWT Token> |
| body | `body` | [AccountUpdate](#account-update) | `models.AccountUpdate` | | ✓ | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#update-account-200) | OK | Success |  | [schema](#update-account-200-schema) |
| [400](#update-account-400) | Bad Request | Bad Request |  | [schema](#update-account-400-schema) |
| [401](#update-account-401) | Unauthorized | Unauthorized |  | [schema](#update-account-401-schema) |
| [404](#update-account-404) | Not Found | Not Found |  | [schema](#update-account-404-schema) |
| [500](#update-account-500) | Internal Server Error | Internal Server Error |  | [schema](#update-account-500-schema) |

#### Responses


##### <span id="update-account-200"></span> 200 - Success
Status: OK

###### <span id="update-account-200-schema"></span> Schema
   
  



##### <span id="update-account-400"></span> 400 - Bad Request
Status: Bad Request

###### <span id="update-account-400-schema"></span> Schema
   
  

[Error](#error)

##### <span id="update-account-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="update-account-401-schema"></span> Schema
   
  

[Error](#error)

##### <span id="update-account-404"></span> 404 - Not Found
Status: Not Found

###### <span id="update-account-404-schema"></span> Schema
   
  

[Error](#error)

##### <span id="update-account-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="update-account-500-schema"></span> Schema
   
  

[Error](#error)

### <span id="update-url"></span> Update a URL (*UpdateURL*)

```
PUT /v1/urls
```

Update URL

#### Security Requirements
  * Bearer

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| Authorization | `header` | string | `string` |  | ✓ |  | Bearer <JWT Token> |
| body | `body` | [URLUpdate](#url-update) | `models.URLUpdate` | | ✓ | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#update-url-200) | OK | Success |  | [schema](#update-url-200-schema) |
| [400](#update-url-400) | Bad Request | Bad Request |  | [schema](#update-url-400-schema) |
| [401](#update-url-401) | Unauthorized | Unauthorized |  | [schema](#update-url-401-schema) |
| [404](#update-url-404) | Not Found | Not Found |  | [schema](#update-url-404-schema) |
| [500](#update-url-500) | Internal Server Error | Internal Server Error |  | [schema](#update-url-500-schema) |

#### Responses


##### <span id="update-url-200"></span> 200 - Success
Status: OK

###### <span id="update-url-200-schema"></span> Schema
   
  

[UpdateURLOKBody](#update-url-o-k-body)

##### <span id="update-url-400"></span> 400 - Bad Request
Status: Bad Request

###### <span id="update-url-400-schema"></span> Schema
   
  

[Error](#error)

##### <span id="update-url-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="update-url-401-schema"></span> Schema
   
  

[Error](#error)

##### <span id="update-url-404"></span> 404 - Not Found
Status: Not Found

###### <span id="update-url-404-schema"></span> Schema
   
  

[Error](#error)

##### <span id="update-url-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="update-url-500-schema"></span> Schema
   
  

[Error](#error)

###### Inlined models

**<span id="update-url-o-k-body"></span> UpdateURLOKBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| long_url | string| `string` |  | |  |  |
| status | string| `string` |  | |  |  |
| tracking_status | string| `string` |  | |  |  |



## Models

### <span id="account"></span> Account


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| email | string| `string` | ✓ | |  |  |
| username | string| `string` | ✓ | |  |  |



### <span id="account-created"></span> AccountCreated


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| password | string| `string` |  | |  |  |
| prefix | string| `string` |  | |  |  |



### <span id="account-update"></span> AccountUpdate


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| status | string| `string` | ✓ | |  |  |
| username | string| `string` | ✓ | |  |  |



### <span id="error"></span> Error


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| error | string| `string` |  | |  |  |



### <span id="login-info"></span> LoginInfo


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| password | string| `string` | ✓ | |  |  |
| username | string| `string` | ✓ | |  |  |



### <span id="login-success"></span> LoginSuccess


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| success | boolean| `bool` |  | |  |  |
| token | string| `string` |  | |  |  |



### <span id="url"></span> URL


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| long_url | string| `string` | ✓ | |  |  |



### <span id="url-created"></span> URLCreated


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| short_url | string| `string` |  | |  |  |
| short_url_key | string| `string` |  | |  |  |



### <span id="url-update"></span> URLUpdate


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| new_long_url | string| `string` |  | |  |  |
| short_url_key | string| `string` | ✓ | |  |  |
| status | string| `string` |  | |  |  |
| tracking_status | string| `string` |  | |  |  |


