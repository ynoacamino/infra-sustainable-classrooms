# Kong Gateway Configuration

Kong Gateway serves as the API Gateway for the Sustainable Classrooms microservices architecture, providing centralized routing, load balancing, authentication, and security features.

## 🎯 Overview

Kong Gateway acts as the single entry point for all client requests, routing them to the appropriate microservices. It provides:

- **API Routing** - Route requests to appropriate microservices
- **Load Balancing** - Distribute traffic across service instances
- **Rate Limiting** - Protect services from overload
- **Logging & Monitoring** - Request logging and metrics collection
- **Security** - SSL termination and security headers

## 📋 Configuration

### Services and Routes

The `kong.yml` file defines all services and their routing configuration:

| Service | Path | Methods | Description |
|---------|------|---------|-------------|
| auth-service | `/api/auth` | GET, POST | Authentication and authorization |
| profiles-service | `/api/profiles` | GET, POST, PUT | User profile management |
| text-service | `/api/text` | GET, POST, PUT, DELETE, PATCH | Text content management |
| video_learning-service | `/api/video_learning` | GET, POST, PUT, DELETE | Video learning platform |
| knowledge-service | `/api/knowledge` | GET, POST, PUT, DELETE | Knowledge base management |
| codelab-service | `/api/codelab` | GET, POST, PUT, DELETE | Interactive code labs |
| mailing-service | `/api/mailing` | GET, POST | Email communication |
| frontend | `/` | GET | Next.js frontend application |

### Service Discovery

Kong is configured to discover services through Docker Compose networking:
- Services are accessible by their container names
- Internal communication happens on port 8080
- External access through Kong on port 8000

## 🔧 Configuration Management

### Declarative Configuration

Kong uses declarative configuration through `kong.yml`:

```yaml
_format_version: "3.0"
_transform: true

services:
  - name: service-name
    host: service-host
    port: 8080
    protocol: http
    routes:
      - name: route-name
        paths:
          - /api/path
        strip_path: false
        methods:
          - GET
          - POST
```

### Adding New Services

To add a new service to Kong:

1. Add service definition to `kong.yml`:
```yaml
services:
  - name: new-service
    host: new-service
    port: 8080
    protocol: http
    routes:
      - name: new-service-route
        paths:
          - /api/new-service
        strip_path: false
```

2. Restart Kong or reload configuration:
```bash
docker compose restart kong
```
