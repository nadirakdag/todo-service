# Todo Service

A modern, production-ready Todo microservice built with Go, featuring JWT authentication, PostgreSQL database, Redis caching, and message queue integration.

## Features

- RESTful API endpoints for todo management
- JWT-based authentication
- PostgreSQL for persistent storage
- Redis for caching
- RabbitMQ for event messaging
- Clean architecture
- Docker support
- API documentation

## Project Structure

```
todo-service/
├── cmd/
│   └── api/
│       └── main.go                 # Application entry point
├── internal/
│   ├── config/
│   │   └── config.go              # Configuration management
│   ├── handler/
│   │   ├── todo.go               # HTTP request handlers for todos
│   │   └── auth.go               # Authentication handlers
│   ├── middleware/
│   │   └── auth.go               # JWT authentication middleware
│   ├── model/
│   │   └── todo.go               # Domain models
│   ├── repository/
│   │   ├── repository.go         # Repository interfaces
│   │   ├── postgres/
│   │   │   └── todo.go           # PostgreSQL implementations
│   │   └── cache/
│   │       └── redis.go          # Redis cache implementation
│   ├── service/
│   │   └── todo.go               # Business logic layer
│   └── messaging/
│       ├── publisher.go          # Message queue publisher
│       └── events.go             # Event definitions
├── pkg/
│   ├── jwt/
│   │   └── token.go              # JWT utilities
│   └── database/
│       └── postgres.go           # Database utilities
├── docker-compose.yml            # Docker compose configuration
├── Dockerfile                    # Docker build file
├── go.mod                        # Go module file
└── README.md                     # This file
```

## Prerequisites

- Go 1.21 or higher
- PostgreSQL 15 or higher
- Redis 7.0 or higher
- RabbitMQ 3.12 or higher
- Docker & Docker Compose (optional)

## Environment Variables

```env
# Server
SERVER_PORT=8080
SERVER_HOST=0.0.0.0

# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=todo_db

# Redis
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=

# RabbitMQ
RABBITMQ_URL=amqp://guest:guest@localhost:5672/

# JWT
JWT_SECRET=your-secret-key
JWT_EXPIRATION=24h
```

## Getting Started

1. Clone the repository:
```bash
git clone https://github.com/yourusername/todo-service.git
cd todo-service
```

2. Install dependencies:
```bash
go mod download
```

3. Set up the environment variables:
```bash
cp .env.example .env
# Edit .env with your configurations
```

4. Run with Docker:
```bash
docker-compose up -d
```

Or run locally:
```bash
go run cmd/api/main.go
```

## API Endpoints

### Authentication
- POST /api/v1/auth/register - Register a new user
- POST /api/v1/auth/login - Login user

### Todos
- GET /api/v1/todos - Get all todos for authenticated user
- GET /api/v1/todos/:id - Get specific todo
- POST /api/v1/todos - Create new todo
- PUT /api/v1/todos/:id - Update todo
- DELETE /api/v1/todos/:id - Delete todo

## Testing

Run all tests:
```bash
go test ./...
```

Run specific package tests:
```bash
go test ./internal/service
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details

## Directory Structure Explanation

- `cmd/`: Contains the main applications of the project
- `internal/`: Contains the private application and library code
    - `config/`: Application configuration
    - `handler/`: HTTP request handlers
    - `middleware/`: HTTP middleware components
    - `model/`: Domain models and business logic interfaces
    - `repository/`: Data access layer implementations
    - `service/`: Business logic implementation
    - `messaging/`: Message queue integration
- `pkg/`: Library code that's ok to use by external applications
    - `jwt/`: JWT token handling
    - `database/`: Database utilities

## Architecture Overview

This project follows clean architecture principles:

1. **Models Layer** (`internal/model`)
    - Contains enterprise/domain logic
    - Defines core business rules and interfaces

2. **Repository Layer** (`internal/repository`)
    - Handles data persistence
    - Implements database operations
    - Manages caching strategies

3. **Service Layer** (`internal/service`)
    - Implements business logic
    - Orchestrates the flow of data
    - Handles complex operations

4. **Handler Layer** (`internal/handler`)
    - Manages HTTP requests/responses
    - Handles input validation
    - Routes to appropriate services

This architecture ensures:
- Separation of concerns
- Testability
- Maintainability
- Scalability