# Emoji Sorter Service

Service for managing emoji.

## Setup 

1. Checkout this repo
2. Make sure `task` and Docker Desktop 3+
    * https://taskfile.dev/#/installation
    * https://www.docker.com/products/docker-desktop
3. Set new relic environment var

## Running 

### Docker

> Docker Desktop version 3.0+ is required.

Setup:
```
docker-compose build emojisorter
```

Running the EmojiSorter service:
```
docker-compose up
```

Unit Tests
```
docker-compose run emojisorter task unit-test
```

All Tests
```
docker-compose up -d
docker-compose exec emojisorter task test
```

Linter
```
docker-compose run emojisorter task lint
```

Go format:
```
docker-compose run emojisorter task fmt
```

Update/generate fakes:
```
docker-compose run emojisorter task fakes
```

Generate Swagger and API docs
```
docker-compose run emojisorter task docs
```

### Local

> Requires installing Go 1.16.5 on your host OS.

Setup:
```
task setup
```

Running the emojisorter service:
```
task start
```

Running unit tests:
```
task test
```

Running the linter:
```
task lint
```

Go format:
```
task fmt
```

Update/generate fakes:
```
task fakes
```

Generate Swagger and API docs
```
task docs
```

# Environment Variables

* ENVIRONMENT (development) - Controls environment specific options
* PORT (8080) - Port to run API server on
* NEW_RELIC_LICENSE_KEY - new relic license key

