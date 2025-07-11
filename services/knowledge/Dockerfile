# Build stage
FROM docker.io/library/golang:latest AS builder

RUN apt-get update && apt-get install -y \
    git \
    ca-certificates \
    tzdata 

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY services/knowledge/ ./services/knowledge/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags='-w -s -extldflags "-static"' \
    -a -installsuffix cgo \
    -o knowledge-service \
    ./services/knowledge/cmd/

# Production stage
FROM scratch

COPY --from=builder /app/knowledge-service /knowledge-service

EXPOSE 8080 8090

USER 65534

ENTRYPOINT ["/knowledge-service"]
