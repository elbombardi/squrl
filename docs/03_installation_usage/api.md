


# Short URL API
  

## Informations

### Version

1.0.0

## Content negotiation

### URI Schemes
  * https

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



**<span id="post-customer-o-k-body"></span> PostCustomerOKBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| api_key | string| `string` |  | |  |  |



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
