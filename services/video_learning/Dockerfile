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
  -o video_learning-service \
  ./services/video_learning/cmd/

# Production stage
FROM scratch

COPY --from=builder /app/video_learning-service /video_learning-service

EXPOSE 8080 9090

USER 65534

ENTRYPOINT ["/video_learning-service"]
