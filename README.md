# Sustainable Classrooms - Microservices Architecture

A comprehensive microservices architecture for managing sustainable educational content and learning experiences. Built with modern technologies and robust design patterns for scalability, maintainability, and performance.

## 🎯 Project Overview

The Sustainable Classrooms platform provides a complete ecosystem for educational content management, including:

- **User Authentication & Authorization** - Secure multi-factor authentication
- **User Profile Management** - Student and teacher profiles with role-based access
- **Text Content Management** - Creation and management of educational text content
- **Video Learning Platform** - Video content delivery and learning progress tracking
- **Knowledge Management** - Structured knowledge base and documentation
- **Interactive Code Labs** - Hands-on coding environments for learning
- **Email Communication** - Automated notifications and communication
- **Frontend Application** - Modern web interface built with Next.js

## 🏗️ Architecture Overview

This project implements a microservices architecture using:

- **Go 1.24.4** - Primary backend language with modern features
- **Next.js 15** - Frontend framework with React 19
- **PostgreSQL 16** - Primary database with multi-database architecture
- **Redis** - Caching and session management
- **MinIO** - Object storage for media files
- **Kong Gateway** - API Gateway for routing and security
- **Docker** - Containerization for all services
- **gRPC** - Inter-service communication
- **HTTP/REST** - Client-facing APIs
- **Goa Framework v3** - Code generation and API design

## 📁 Project Structure

```
infra-sustainable-classrooms/
├── README.md                       # Main project documentation
├── go.mod                         # Go dependencies
├── go.sum                         # Go dependency checksums
├── package.json                   # Node.js dependencies for tooling
├── docker-compose.yml             # Production service orchestration
├── docker-compose.dev.yml         # Development configuration
├── generate.sh                    # Global code generation script
├── sqlc.json                      # SQLC configuration
├── commitlint.config.ts           # Commit message linting
├── flake.nix                      # Nix development environment
│
├── db/                            # Database Configuration
│   ├── init/                      # Database initialization scripts
│   │   ├── init_databases.sh      # Multi-database setup script
│   │   ├── auth_db.sql           # Authentication database schema
│   │   ├── profiles_db.sql       # User profiles database schema
│   │   ├── text_db.sql           # Text content database schema
│   │   ├── video_learning_db.sql # Video learning database schema
│   │   ├── knowledge_db.sql      # Knowledge management database schema
│   │   └── codelab_db.sql        # Code lab database schema
│   └── mock/                     # Mock data for testing
│
├── kong/                         # API Gateway Configuration
│   ├── README.md                 # Kong setup and configuration guide
│   └── kong.yml                  # Kong Gateway routing rules
│
├── mailing/                      # Email Infrastructure
│   ├── README.md                 # Email infrastructure documentation
│   ├── compose.yaml              # Mail server docker compose
│   ├── mailing-setup.md          # Detailed mail server setup guide
│   ├── minimal-mailing-setup.md  # Quick setup guide
│   └── mailserver.env            # Mail server environment variables
│
├── minio/                        # Object Storage
│   ├── README.md                 # MinIO setup and configuration guide
│   ├── init_buckets.sh           # MinIO bucket initialization
│   └── data/                     # Storage buckets structure
│       ├── video-learning-thumbnails-confirmed/
│       ├── video-learning-thumbnails-staging/
│       ├── video-learning-videos-confirmed/
│       └── video-learning-videos-staging/
│
└── services/                     # Microservices
    ├── auth/                     # Authentication Service
    ├── profiles/                 # User Profile Management Service
    ├── text/                     # Text Content Management Service
    ├── video_learning/           # Video Learning Platform Service
    ├── knowledge/                # Knowledge Management Service
    ├── codelab/                  # Interactive Code Labs Service
    ├── mailing/                  # Email Communication Service
    └── frontend/                 # Web Frontend Application
```

Each service follows a consistent structure with its own documentation. See individual service README files for detailed information.

## 🚀 Microservices Overview

The platform consists of the following microservices, each with dedicated documentation:

### Core Services

| Service                    | Port | Description                                       |
| -------------------------- | ---- | ------------------------------------------------- |
| **Auth Service**           | 8081 | Multi-factor authentication, session management   |
| **Profiles Service**       | 8082 | User profile management for students and teachers |
| **Text Service**           | 8083 | Text-based educational content management         |
| **Video Learning Service** | 8084 | Video content delivery and progress tracking      |
| **Knowledge Service**      | 8085 | Knowledge base and documentation management       |
| **Codelab Service**        | 8086 | Interactive coding environments                   |
| **Mailing Service**        | 8087 | Email notifications and communication             |

### Frontend & Infrastructure

| Component          | Port               | Description                                | Documentation                      |
| ------------------ | ------------------ | ------------------------------------------ | ---------------------------------- |
| **Frontend App**   | 3000               | Next.js web application with modern UI     | -                                  |
| **Kong Gateway**   | 8000               | API Gateway for routing and load balancing | [📖 Configuration](kong/README.md) |
| **PostgreSQL**     | 5432               | Multi-database setup for service isolation | [📖 Schema](db/README.md)          |
| **Mailing Server** | 25,143,465,587,993 | SMTP server for email communication        | [📖 Setup](mailing/README.md)      |
| **Redis**          | 6379               | Caching and session storage                | -                                  |
| **MinIO**          | 9000               | Object storage for media files             | [📖 Setup](minio/README.md)        |

## 🛠️ Technology Stack

### Backend Technologies

| Category      | Technology | Version   | Purpose                            |
| ------------- | ---------- | --------- | ---------------------------------- |
| **Language**  | Go         | 1.24.4    | Primary backend language           |
| **Framework** | Goa v3     | 3.21.1    | API framework with code generation |
| **Database**  | PostgreSQL | 16-alpine | Primary data storage               |
| **Cache**     | Redis      | alpine    | Session and data caching           |
| **Storage**   | MinIO      | latest    | Object storage for media files     |
| **DB Driver** | pgx/v5     | 5.7.5     | High-performance PostgreSQL driver |
| **Code Gen**  | SQLC       | latest    | Type-safe SQL code generation      |

### Frontend Technologies

| Category             | Technology   | Version | Purpose                         |
| -------------------- | ------------ | ------- | ------------------------------- |
| **Framework**        | Next.js      | 15.3.5  | React-based web framework       |
| **Runtime**          | React        | 19.0.0  | UI component library            |
| **Language**         | TypeScript   | latest  | Type-safe JavaScript            |
| **Styling**          | Tailwind CSS | latest  | Utility-first CSS framework     |
| **UI Components**    | Radix UI     | latest  | Accessible component primitives |
| **State Management** | Zustand      | 5.0.2   | Lightweight state management    |
| **Data Fetching**    | SWR          | 2.3.3   | Data fetching and caching       |

### Infrastructure & DevOps

| Category             | Technology              | Purpose                              |
| -------------------- | ----------------------- | ------------------------------------ |
| **Containerization** | Docker & Docker Compose | Service orchestration                |
| **API Gateway**      | Kong Gateway            | Request routing and load balancing   |
| **Code Generation**  | Goa DSL                 | API design and code generation       |
| **Development**      | Nix Flakes              | Reproducible development environment |
| **Linting**          | ESLint, Prettier        | Code formatting and quality          |
| **Git Hooks**        | Husky, Commitlint       | Commit message validation            |

### Security & Authentication

| Technology           | Purpose                                 |
| -------------------- | --------------------------------------- |
| **TOTP**             | Time-based One-Time Password for MFA    |
| **bcrypt**           | Secure password hashing                 |
| **Session Cookies**  | HTTPOnly cookies for session management |
| **Rate Limiting**    | Request throttling and abuse prevention |
| **Input Validation** | Request validation with Goa             |

## Quick Start

### Prerequisites

- **Docker** and **Docker Compose** installed
- **Go 1.24.4** or later (for development)
- **Node.js 18+** and **pnpm** (for frontend development)
- **Git** for version control

### Development Setup

1. **Clone the repository**

   ```bash
   git clone https://github.com/ynoacamino/infra-sustainable-classrooms.git
   cd infra-sustainable-classrooms
   ```

2. **Set up environment files**

   ```bash
   # Copy example environment files for each service
   cp services/auth/.env.example services/auth/.env.dev
   cp services/profiles/.env.example services/profiles/.env.dev
   # ... repeat for other services
   ```

3. **Start all services in development mode**

   ```bash
   docker-compose -f docker-compose.dev.yml up --build
   ```

4. **Generate code (if making changes to API design)**
   ```bash
   ./generate.sh
   ```

### Production Deployment

1. **Configure production environment**

   ```bash
   # Set up production environment files
   cp services/auth/.env.example services/auth/.env.prod
   # Configure production values
   ```

2. **Deploy services**
   ```bash
   docker-compose up --build -d
   ```

## 📊 Architecture Patterns

### Hexagonal Architecture (Ports & Adapters)

- **Ports**: Define contracts and interfaces for external interactions
- **Adapters**: Concrete implementations of ports (HTTP, gRPC, Database)
- **Controllers**: Application logic and business rules
- **Repositories**: Data access abstraction layer

### Repository Pattern

- **Abstraction**: Repository interfaces define data access contracts
- **Implementation**: Concrete repositories handle specific storage mechanisms
- **Testing**: Easy mocking and testing through interfaces
- **Transactions**: Centralized transaction management

### Dependency Injection

- **Configuration**: Centralized service configuration
- **Testing**: Facilitates unit testing and mocking
- **Maintenance**: Loose coupling between components
- **Flexibility**: Easy to swap implementations

### Code Generation

- **API Design**: Goa DSL for defining APIs declaratively
- **Type Safety**: SQLC generates type-safe Go code from SQL
- **Consistency**: Automated generation ensures consistency
- **Documentation**: Auto-generated OpenAPI specifications

### Security Features

- **Multi-Factor Authentication**: TOTP-based 2FA with backup codes
- **Session Management**: Secure, HTTPOnly cookies with Redis backend
- **Rate Limiting**: Request throttling to prevent abuse
- **Input Validation**: Comprehensive request validation
- **Secure Headers**: Security headers in all responses
- **Database Security**: Parameterized queries prevent SQL injection

## 🚀 Development

### Prerequisites

- Go 1.24.4+
- PostgreSQL 16+
- Redis (latest)
- Docker (optional)

### Local Development

1. **Set up environment**

   ```bash
   cp .env.example .env.dev
   # Edit .env.dev with your configuration
   ```

2. **Start dependencies**

   ```bash
   # Using Docker Compose
   docker-compose -f docker-compose.dev.yml up postgres redis
   ```

3. **Generate code**

   ```bash
   ./generate.sh
   ```

4. **Run the service**
   ```bash
   go run cmd/main.go
   ```

### Code Generation

The service uses Goa for code generation:

```bash
# Generate all code
./generate.sh

# Generate code per microservice
goa gen github.com/ynoacamino/infra-sustainable-classrooms/services/auth/design
```

## 📄 License

This project is licensed under the [MIT License](LICENSE) - see the LICENSE file for details.
