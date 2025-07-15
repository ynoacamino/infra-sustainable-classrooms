# Build stage
FROM docker.io/library/golang:latest AS builder

RUN apt-get update && apt-get install -y \
  git \
  ca-certificates \
  tzdata

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY services/ ./services/

RUN CGO_ENABLED=0 GOOS=linux go build \
  -a -installsuffix cgo \
  -o text-service \
  ./services/text/cmd/

# Production stage
FROM scratch

COPY --from=builder /app/text-service /text-service

EXPOSE 8080 9090

USER 65534

ENTRYPOINT ["/text-service"]
