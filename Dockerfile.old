ARG ECR_HOST
ARG APP_NAME

# Golang fromline not available in dev account. 
# FROM ${ECR_HOST}/golang:1.16-alpine AS preloader
FROM 652798529812.dkr.ecr.us-east-1.amazonaws.com/golang:1.16-alpine AS preloader

WORKDIR /build
ENV CGO_ENABLED=0
ENV GO111MODULE=on
ENV GOOS=linux
RUN go get github.com/go-delve/delve/cmd/dlv@v1

#  Get dependencies before build
COPY go.mod go.sum ./
RUN --mount=type=cache,target=$GOPATH/pkg/mod go mod download

FROM ${ECR_HOST}/drizlyinc/${APP_NAME}/preloader:${TAG:-local} AS builder

# Set up build/debug env
WORKDIR /build
ENV CGO_ENABLED=0
ENV GO111MODULE=on
ENV GOOS=linux
RUN go get github.com/go-delve/delve/cmd/dlv@v1

#  Get dependencies before build
COPY go.mod go.sum ./
RUN --mount=type=cache,target=$GOPATH/pkg/mod go mod download

# Build binary
COPY . .
RUN go build ./cmd/...

FROM builder AS formatter
ENTRYPOINT [ "go", "fmt", "./..."]

FROM builder AS debugger
ENTRYPOINT [ "dlv", "-l", ":40000", "--headless=true", "--api-version=2", "exec", "./app", "--" ]

# To execute tests: `docker run --rm $(docker build -q --target test .)`
FROM builder AS unit-tester
CMD go test -v ./...

FROM builder as unit-test-debugger
CMD dlv -l :40000 --headless=true --api-version=2 test -test.v ./...


FROM gcr.io/distroless/base:nonroot AS server
# set user to nonroot
USER nonroot
WORKDIR /
COPY --from=builder /build/app .

EXPOSE 8080
ENTRYPOINT [ "./app" ]
