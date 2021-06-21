# Orders Service

Service for managing orders.

## Setup 

1. Checkout this repo
2. Make sure `task` and Docker Desktop
    * https://taskfile.dev/#/installation
    * https://www.docker.com/products/docker-desktop

## Running 

### Docker

TBD

### Local

> Requires installing the latest version of Go 1.16.x.

Setup:
```
task local:setup
```

Run lint:
```
task local:lint
```

Update/generate fakes:
```
go generate ./...
```

Run tests:
```
task local:test
```

Run service:
```
task local:orders:start
```

# Environment Variables

* ENVIRONMENT (development) - Controls environment specific options
* PORT (8080) - Port to run API server on

