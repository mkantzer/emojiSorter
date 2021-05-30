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

# Build binary
COPY . .
RUN go build -o app

FROM builder AS linter
ENTRYPOINT [ "go", "fmt" ]

FROM builder AS debugger
ENTRYPOINT [ "dlv", "-l", ":40000", "--headless=true", "--api-version=2", "exec", "./app", "--" ] 

FROM builder AS tester
CMD go test -v ./...

FROM test as test-debugger
CMD dlv -l :40000 --headless=true --api-version=2 test -test.v ./...

FROM gcr.io/distroless/base:nonroot AS runtime
USER nonroot
WORKDIR /
COPY --from=builder --chown=nonroot /build/app .

ENTRYPOINT [ "./app" ]