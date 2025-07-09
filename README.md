# Infrastructure - Microservices Architecture

Microservices architecture in Go with a focus on authentication, built with modern technologies and robust design patterns.

## 🏗️ General Architecture

This project implements a microservices architecture using:

* **Go** as the main language
* **PostgreSQL** as the database
* **Docker** for containerization
* **gRPC** for inter-service communication
* **HTTP/REST** for client APIs
* **Goa Framework** for code generation

## 📁 Project Structure

```
infrastructure/
├── README.md
├── go.mod                          # Project dependencies
├── go.sum
├── docker-compose.yml              # Service orchestration
├── docker-compose.dev.yml          # Development configuration
├── generate.sh                     # Code generation script
├── sqlc.json                       # SQLC configuration
│
├── db/                             # Database
│   ├── schemas/                    # Schemas and migrations
│   │   └── 0001_auth.sql          # Initial authentication migration
│   ├── .env                       # DB environment variables
│   ├── .env.dev                   # Development environment variables
│   └── .env.example               # Example configuration
│
└── services/                      # Microservices
    ├── auth/                      # Authentication service
    │   ├── Dockerfile             # Production image
    │   ├── Dockerfile.dev         # Development image
    │   ├── generate.sh            # Service code generation
    │   ├── README.md              # Service documentation
    │   │
    │   ├── cmd/                   # Entry point
    │   │   ├── main.go            # Main application
    │   │   ├── http.go            # HTTP server
    │   │   └── grpc.go            # gRPC server
    │   │
    │   ├── config/                # Configuration
    │   │   ├── config.go          # Main configuration
    │   │   ├── helpers.go         # Config utilities
    │   │   └── README.md          # Config documentation
    │   │
    │   ├── design/                # API design with Goa
    │   │   ├── api/               # API definitions
    │   │   │   ├── design.go      # Global config
    │   │   │   ├── auth.go        # Auth endpoints
    │   │   │   ├── types.go       # Data types
    │   │   │   └── README.md      # Design documentation
    │   │   └── queries/           # SQL queries
    │   │       ├── auth_attempts.sql
    │   │       ├── backup_codes.sql
    │   │       ├── sessions.sql
    │   │       └── users.sql
    │   │
    │   ├── gen/                   # Generated code (do not edit)
    │   │   ├── auth/              # Goa generated code
    │   │   │   ├── client.go
    │   │   │   ├── endpoints.go
    │   │   │   └── service.go
    │   │   ├── database/          # SQLC generated code
    │   │   │   ├── db.go
    │   │   │   ├── models.go
    │   │   │   └── *.sql.go
    │   │   ├── grpc/              # Generated gRPC code
    │   │   │   └── auth/
    │   │   └── http/              # Generated HTTP code
    │   │       ├── openapi.json
    │   │       ├── openapi.yaml
    │   │       └── auth/
    │   │
    │   └── internal/              # Business logic
    │       ├── controllers/       # Controllers
    │       │   ├── endpoints.go   # Endpoint implementation
    │       │   ├── helpers.go     # Utilities
    │       │   ├── service.go     # Main service
    │       │   └── *_test.go      # Tests
    │       ├── database/          # DB connection
    │       │   ├── connect.go     # Connection config
    │       │   └── migration.go   # Migrations
    │       ├── mappers/           # Data mappers
    │       │   └── mappers.go     # Layer conversions
    │       ├── ports/             # Interfaces (Hexagonal Architecture)
    │       │   ├── user_repository.go
    │       │   ├── session_repository.go
    │       │   ├── backup_code_repository.go
    │       │   └── transaction_manager.go
    │       └── repositories/      # Repository implementations
    │           ├── repository_manager.go
    │           ├── user_repository.go
    │           ├── session_repository.go
    │           ├── backup_code_repository.go
    │           ├── transaction_manager.go
    │           └── mocks/         # Mocks for testing
    │
    └── executor/                  # Task executor service (future)
```

## 🚀 Services

### Auth Service

Authentication microservice with the following features:

* **TOTP (Time-based One-Time Password)** for MFA
* **Backup codes** for account recovery
* **Session management** with secure cookies
* **Rate limiting** to prevent abuse
* **Dual transport**: HTTP for clients, gRPC for inter-service communication

**Main Endpoints:**

* `POST /auth/totp/generate` - Generate TOTP secret
* `POST /auth/totp/verify` - Verify TOTP code
* `POST /auth/backup/verify` - Verify backup code
* `POST /auth/session/refresh` - Refresh session
* `GET /auth/profile` - Get user profile

## 🛠️ Technologies

### Backend

* **Go 1.24.1** - Primary language
* **Goa v3** - API framework with code generation
* **PostgreSQL 16** - Main database
* **SQLC** - Type-safe SQL code generation
* **pgx/v5** - High-performance PostgreSQL driver

### Authentication & Security

* **TOTP** - Time-based One-Time Password
* **bcrypt** - Secure password hashing
* **Sessions** - Managed with cookies
* **Rate limiting** - Attack mitigation

### Observability

* **OpenTelemetry** - Distributed tracing
* **Structured logging** - With goa/clue
* **Health checks** - Service monitoring

### Infrastructure

* **Docker** - Containerization
* **Docker Compose** - Local orchestration
* **Multi-stage builds** - Image optimization

## 🔧 Useful Commands

```bash
# Development
docker-compose -f docker-compose.dev.yml up --build

# Production
docker-compose up --build

# Generate code
./generate.sh

# Generate auth service code
cd services/auth && ./generate.sh

# Run tests
go test ./...
```

## 📊 Design Patterns

### Hexagonal Architecture

* **Ports**: Define contracts (interfaces)
* **Adapters**: Concrete implementations
* **Controllers**: Application logic
* **Repositories**: Data access

### Repository Pattern

* Abstracts data access
* Enables easy mocking for tests
* Supports transactions

### Dependency Injection

* Centralized configuration
* Facilitates testing and maintenance
* Inverts control of dependencies

## 🔐 Security

* **TOTP** for multifactor authentication
* **Backup codes** for recovery
* **Secure sessions** with HTTPOnly cookies
* **Rate limiting** on critical endpoints
* **Input validation** on all endpoints
* **Secure hashing** with bcrypt
