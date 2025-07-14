# Build stage
FROM docker.io/library/golang:latest AS builder

RUN apt-get update && apt-get install -y \
    git \
    ca-certificates \
    tzdata 

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY services/mailing/ ./services/mailing/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags='-w -s -extldflags "-static"' \
    -a -installsuffix cgo \
    -o mailing-service \
    ./services/mailing/cmd/

# Production stage
FROM scratch

COPY --from=builder /app/mailing-service /mailing-service

EXPOSE 8080 8090

USER 65534

ENTRYPOINT ["/mailing-service"]
