FROM golang:1.15-alpine AS builder

# Set up build/debug env
WORKDIR /build
ENV CGO_ENABLED=0
ENV GO111MODULE=on
ENV GOOS=linux
RUN go get github.com/go-delve/delve/cmd/dlv@v1

#  Get dependancies before build, and cache the mod cache folder
COPY go.mod go.sum ./
RUN --mount=type=cache,target=$GOPATH/pkg/mod go mod download

# Build server
COPY . .
# RUN go build -o server cmd/server/main.go
RUN go build -o cli cmd/cli/main.go

FROM builder AS linter
ENTRYPOINT [ "go", "fmt" ]

FROM builder AS debugger
ENTRYPOINT [ "dlv", "-l", ":40000", "--headless=true", "--api-version=2", "exec", "./app", "--" ] 

FROM builder AS tester
CMD go test -v ./...

FROM test as test-debugger
CMD dlv -l :40000 --headless=true --api-version=2 test -test.v ./...

FROM gcr.io/distroless/base:nonroot AS server
USER nonroot
WORKDIR /
COPY --from=builder --chown=nonroot /build/server app

ENTRYPOINT [ "./app" ]

FROM gcr.io/distroless/base:nonroot AS cli
USER nonroot
WORKDIR /
COPY --from=builder --chown=nonroot /build/cli app

ENTRYPOINT [ "./app" ]