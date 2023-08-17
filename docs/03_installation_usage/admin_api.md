


# SQURL - ADMIN API
  

## Informations

### Version

2.0.0

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
| POST | /v2/account | [post account](#post-account) | Create account |
| POST | /v2/short-url | [post short URL](#post-short-url) | Create ShortURL |
| PUT | /v2/account | [put account](#put-account) | Update Account |
| PUT | /v2/short-url | [put short URL](#put-short-url) | Update ShortURL |
  


## Paths

### <span id="post-account"></span> Create account (*PostAccount*)

```
POST /v2/account
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| Authorization | `header` | string | `string` |  | ✓ |  | JWT Token. |
| account | `body` | [PostAccountBody](#post-account-body) | `PostAccountBody` | | ✓ | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#post-account-200) | OK | Success |  | [schema](#post-account-200-schema) |
| [400](#post-account-400) | Bad Request | Bad Request |  | [schema](#post-account-400-schema) |
| [401](#post-account-401) | Unauthorized | Unauthorized |  | [schema](#post-account-401-schema) |
| [500](#post-account-500) | Internal Server Error | Internal Server Error |  | [schema](#post-account-500-schema) |

#### Responses


##### <span id="post-account-200"></span> 200 - Success
Status: OK

###### <span id="post-account-200-schema"></span> Schema
   
  

[PostAccountOKBody](#post-account-o-k-body)

##### <span id="post-account-400"></span> 400 - Bad Request
Status: Bad Request

###### <span id="post-account-400-schema"></span> Schema
   
  

[PostAccountBadRequestBody](#post-account-bad-request-body)

##### <span id="post-account-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="post-account-401-schema"></span> Schema
   
  

[PostAccountUnauthorizedBody](#post-account-unauthorized-body)

##### <span id="post-account-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="post-account-500-schema"></span> Schema
   
  

[PostAccountInternalServerErrorBody](#post-account-internal-server-error-body)

###### Inlined models

**<span id="post-account-bad-request-body"></span> PostAccountBadRequestBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| error | string| `string` |  | |  |  |



**<span id="post-account-body"></span> PostAccountBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| email | string| `string` | ✓ | |  |  |
| username | string| `string` | ✓ | |  |  |



**<span id="post-account-internal-server-error-body"></span> PostAccountInternalServerErrorBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| error | string| `string` |  | |  |  |



**<span id="post-account-o-k-body"></span> PostAccountOKBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| api_key | string| `string` |  | |  |  |
| prefix | string| `string` |  | |  |  |



**<span id="post-account-unauthorized-body"></span> PostAccountUnauthorizedBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| error | string| `string` |  | |  |  |



### <span id="post-short-url"></span> Create ShortURL (*PostShortURL*)

```
POST /v2/short-url
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| X-API-KEY | `header` | string | `string` |  | ✓ |  | The account API key. |
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



### <span id="put-account"></span> Update Account (*PutAccount*)

```
PUT /v2/account
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| X-API-KEY | `header` | string | `string` |  | ✓ |  | The admin API key. |
| body | `body` | [PutAccountBody](#put-account-body) | `PutAccountBody` | | ✓ | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#put-account-200) | OK | Success |  | [schema](#put-account-200-schema) |
| [400](#put-account-400) | Bad Request | Bad Request |  | [schema](#put-account-400-schema) |
| [401](#put-account-401) | Unauthorized | Unauthorized |  | [schema](#put-account-401-schema) |
| [404](#put-account-404) | Not Found | Not Found |  | [schema](#put-account-404-schema) |
| [500](#put-account-500) | Internal Server Error | Internal Server Error |  | [schema](#put-account-500-schema) |

#### Responses


##### <span id="put-account-200"></span> 200 - Success
Status: OK

###### <span id="put-account-200-schema"></span> Schema
   
  

[PutAccountOKBody](#put-account-o-k-body)

##### <span id="put-account-400"></span> 400 - Bad Request
Status: Bad Request

###### <span id="put-account-400-schema"></span> Schema
   
  

[PutAccountBadRequestBody](#put-account-bad-request-body)

##### <span id="put-account-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="put-account-401-schema"></span> Schema
   
  

[PutAccountUnauthorizedBody](#put-account-unauthorized-body)

##### <span id="put-account-404"></span> 404 - Not Found
Status: Not Found

###### <span id="put-account-404-schema"></span> Schema
   
  

[PutAccountNotFoundBody](#put-account-not-found-body)

##### <span id="put-account-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="put-account-500-schema"></span> Schema
   
  

[PutAccountInternalServerErrorBody](#put-account-internal-server-error-body)

###### Inlined models

**<span id="put-account-bad-request-body"></span> PutAccountBadRequestBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| error | string| `string` |  | |  |  |



**<span id="put-account-body"></span> PutAccountBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| status | string| `string` | ✓ | |  |  |
| username | string| `string` | ✓ | |  |  |



**<span id="put-account-internal-server-error-body"></span> PutAccountInternalServerErrorBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| error | string| `string` |  | |  |  |



**<span id="put-account-not-found-body"></span> PutAccountNotFoundBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| error | string| `string` |  | |  |  |



**<span id="put-account-o-k-body"></span> PutAccountOKBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| status | string| `string` |  | |  |  |



**<span id="put-account-unauthorized-body"></span> PutAccountUnauthorizedBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| error | string| `string` |  | |  |  |



### <span id="put-short-url"></span> Update ShortURL (*PutShortURL*)

```
PUT /v2/short-url
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| X-API-KEY | `header` | string | `string` |  | ✓ |  | The account API key. |
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



## Models

### <span id="cart-item"></span> CartItem


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| currency | string| `string` |  | |  |  |
| imageUrl | string| `string` |  | |  |  |
| productId | integer| `int64` |  | |  |  |
| productName | string| `string` |  | |  |  |
| quantity | integer| `int64` |  | |  |  |
| unitPrice | number| `float64` |  | |  |  |



### <span id="cart-preview"></span> CartPreview


  

[][CartItem](#cart-item)

### <span id="login-info"></span> LoginInfo


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| email | string| `string` | ✓ | |  |  |
| password | string| `string` | ✓ | |  |  |



### <span id="login-success"></span> LoginSuccess


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| success | boolean| `bool` |  | |  |  |
| token | string| `string` |  | |  |  |


