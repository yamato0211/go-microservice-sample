ARG GO_VERSION=1.21
ARG ALPINE_VERSION=3.18
# ===== build go binary =====
FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} as go-builder

WORKDIR /go/src/go-micro-sample

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o api todo/cmd/main.go

# ==== build docker image ====
FROM alpine:${ALPINE_VERSION}

WORKDIR /usr/src/go-micro-sample

COPY --from=go-builder /go/src/go-micro-sample/api api
COPY --from=go-builder /go/src/go-micro-sample/api .env
RUN chmod +x "/usr/src/go-micro-sample/api"

ENTRYPOINT ["/usr/src/go-micro-sample/api"]