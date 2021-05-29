FROM golang:1.15-alpine AS builder

# Set up build/debug env
WORKDIR /build
ENV CGO_ENABLED=1
ENV GO111MODULE=on
ENV GOOS=linux
RUN go get github.com/go-delve/delve/cmd/dlv@v1

#  Get dependancies before build, and cache the mod cache folder
COPY go.mod go.sum ./
RUN --mount=type=cache,target=$GOPATH/pkg/mod go mod download

# To prime the import cache:
# `go generate ./...` outside of container
# COPY ./internal/imports ./internal/imports
# RUN go build ./internal/imports

# Build binary
COPY . .
RUN go build -o app

FROM builder AS debug
ENTRYPOINT [ "dlv", "-l", ":40000", "--headless=true", "--api-version=2", "exec", "./app", "--" ] 

# To execute tests: `docker run --rm $(docker build -q --target test .)`
FROM builder AS test
CMD go test -v ./...

FROM test as debugTests
CMD dlv -l :40000 --headless=true --api-version=2 test -test.v ./...


FROM gcr.io/distroless/base:nonroot AS runtime
# set user to nonroot
USER nonroot
WORKDIR /
COPY --from=builder /build/app .

ENTRYPOINT [ "./app" ]