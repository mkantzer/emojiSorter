ARG ECR_HOST
ARG APP_NAME
ARG TAG=local

# Golang fromline not available in dev account. 
# FROM ${ECR_HOST}/golang:1.16-alpine AS preloader
FROM 652798529812.dkr.ecr.us-east-1.amazonaws.com/golang:1.16-alpine as base

ARG GIT_HASH

# hadolint ignore=DL3018
RUN apk add --no-cache \
    curl>7.77.0-r1 \
    git>2.30.2-r0

WORKDIR /build
ENV CGO_ENABLED=0
ENV GO111MODULE=on
ENV GOOS=linux

# Build binary
COPY . .

SHELL ["/bin/ash", "-eo", "pipefail", "-c"]
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.40.1 \
 && curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s \
 && curl -sSfL https://taskfile.dev/install.sh | sh -s -- -b /usr/local/bin \
 && go build -o bin/ github.com/go-swagger/go-swagger/cmd/swagger \
 && go build -o bin/ -ldflags "-X main.gitHash=${GIT_HASH}" ./cmd/... 

FROM base AS development
EXPOSE 8080
CMD [ "./bin/air" ]

FROM gcr.io/distroless/base:nonroot AS production
# set user to nonroot
USER nonroot
WORKDIR /
COPY --from=development /build/bin/orders .
EXPOSE 8080
CMD [ "./orders" ]
