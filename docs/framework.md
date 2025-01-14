# Framework Documentation

## 1. Directory Structure
```
.
├── .env                  # Environment variables
├── .env.example         # Environment template
├── .gitignore           # Git ignore rules
├── LICENSE             # License file
├── README.md           # Project documentation
├── api/                # API definitions
│   └── openapi/       # OpenAPI/Swagger docs
│       └── swagger.yaml
├── cmd/               # Main applications
│   └── main.go       # Main entry point
├── configs/           # Configuration files
│   ├── config.dev.yaml   # Development config
│   └── config.prod.yaml  # Production config
├── docs/             # Documentation files
│   ├── database.sql  # Database scripts
│   └── framework.md  # Framework documentation
├── internal/         # Private application code
│   ├── api/         # API handlers
│   │   ├── request.go   # Request models
│   │   ├── response.go  # Response models
│   │   └── signal.go    # Business handlers
│   ├── config/      # Configuration
│   │   ├── config.go    # Config structures
│   │   └── loader.go    # Config loader
│   ├── di/          # Dependency injection
│   │   ├── provider/    # DI providers
│   │   │   └── provider.go
│   │   ├── wire.go      # Wire definitions
│   │   └── wire_gen.go  # Generated wire code
│   ├── errors/      # Error handling
│   │   ├── code.go     # Error codes
│   │   └── errors.go   # Error definitions
│   ├── model/       # Data models
│   │   └── signal.go
│   ├── repository/  # Data access
│   │   ├── base.go     # Base repository
│   │   └── signal.go   # Business repository
│   ├── router/      # HTTP routing
│   │   ├── health.go   # Health check
│   │   └── router.go   # Router setup
│   └── service/     # Business logic
│       └── signal.go
├── logs/            # Log files
│   └── app.log
└── pkg/             # Public libraries
    ├── database/    # Database utilities
    │   └── mysql.go
    ├── http/        # HTTP utilities
    │   ├── middleware/
    │   │   ├── error.go
    │   │   └── logger.go
    │   └── server.go
    ├── logger/      # Logging utilities
    │   └── logger.go
    └── utils/       # Common utilities
        └── random.go
```

## 2. Core Frameworks

### 2.1 Web Framework - Gin
- Selection rationale: High performance, lightweight, active community
- Main features: Routing, middleware support, parameter validation
- Usage: API endpoints, middleware processing

### 2.2 Database - GORM
- Selection rationale: Feature-rich, active community, comprehensive documentation
- Main features: ORM mapping, transaction management, associations
- Usage: Database operations, model relationships

### 2.3 Dependency Injection - Wire
- Selection rationale: Compile-time DI, type-safe
- Main features: Dependency management, interface binding
- Usage: Component initialization
- Location: internal/di/

### 2.4 Configuration - Viper
- Selection rationale: Standard configuration library
- Main features: Config file reading, environment variable support
- Configuration methods:
  - Environment: .env file (local development)
  - Config files: config.{env}.yaml
  - Environments: dev/prod

### 2.5 Logging - Zap
- Selection rationale: High performance, structured logging
- Main features: Log levels, file rotation, structured output
- Usage: System logs, business logs, error tracking
- Location: pkg/logger/

## 3. Key Standards

### 3.1 Project Standards
- Follow standard Go project layout
- Business code in internal directory
- Public code in pkg directory
- Configurations in configs directory

### 3.2 API Standards
- RESTful style
- Unified error handling
- Unified response format
- OpenAPI/Swagger documentation

### 3.3 Error Handling Standards
- Error codes defined in internal/errors
- Unified error response format
- Middleware error handling
- Business error logging

### 3.4 Logging Standards
- Use structured logging
- Unified log format
- Log levels: DEBUG/INFO/WARN/ERROR/FATAL
- Log files in logs directory

### 3.5 Configuration Standards
- Environment variables:
  - Development: .env file
  - Production: System environment variables
- Configuration files:
  - Development: config.dev.yaml
  - Production: config.prod.yaml
- Sensitive information via environment variables

## 4. Development Process

### 4.1 Local Development
1. Copy .env.example to .env
2. Modify .env configuration
3. Ensure configs/config.dev.yaml is properly configured
4. Run `go run cmd/main.go`

### 4.2 Production Deployment
1. Prepare config.prod.yaml
2. Set necessary environment variables
3. Start service with APP_ENV=prod

### 4.3 Dependency Management
- Use go mod for dependency management
- Generate dependency injection code:
  ```bash
  cd internal/di && wire
  ```
