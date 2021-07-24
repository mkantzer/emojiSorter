


# Emoji API
API for managing and voting on emoji

Schemas: http
  

## Informations

### Version

v1

## Content negotiation

### URI Schemes
  * http

### Consumes
  * application/json

### Produces
  * application/json
  * plain/text

## All endpoints

###  operations

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | / | [hello](#hello) |  |
| GET | /healthz | [health status](#health-status) |  |
| GET | /unhealthz | [unhealth status](#unhealth-status) |  |
  


## Paths

### <span id="hello"></span> hello (*Hello*)

```
GET /
```

Returns Hello World

#### Produces
  * plain/text

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#hello-200) | OK | Body contains "Hello World!" |  | [schema](#hello-200-schema) |

#### Responses


##### <span id="hello-200"></span> 200 - Body contains "Hello World!"
Status: OK

###### <span id="hello-200-schema"></span> Schema

### <span id="health-status"></span> health status (*healthStatus*)

```
GET /healthz
```

Returns server health status

#### Produces
  * application/json

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#health-status-200) | OK | healthResponse |  | [schema](#health-status-200-schema) |

#### Responses


##### <span id="health-status-200"></span> 200 - healthResponse
Status: OK

###### <span id="health-status-200-schema"></span> Schema
   
  

[HealthResponse](#health-response)

### <span id="unhealth-status"></span> unhealth status (*unhealthStatus*)

```
GET /unhealthz
```

Returns a 500, useful for testing

#### Produces
  * application/json

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [500](#unhealth-status-500) | Internal Server Error | healthResponse |  | [schema](#unhealth-status-500-schema) |

#### Responses


##### <span id="unhealth-status-500"></span> 500 - healthResponse
Status: Internal Server Error

###### <span id="unhealth-status-500-schema"></span> Schema
   
  

[HealthResponse](#health-response)

## Models

### <span id="health-response"></span> healthResponse


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Message | string| `string` |  | | Optional message for error responses |  |
| Status | string| `string` | ✓ | | The health of the service instance | `\"ok\" or \"error\` |


