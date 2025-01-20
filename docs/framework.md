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
│   ├── cron/        # Cron jobs
│   │   ├── cron.go
│   │   └── jobs/
│   │       └── signal.go
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
    ├── cache/       # Cache utilities
    │   ├── cache.go
    │   └── redis/
    │       └── redis.go
    ├── cron/        # Cron utilities
    │   └── cron.go
    ├── database/    # Database utilities
    │   ├── database.go
    │   └── mysql/
    │       └── mysql.go
    ├── http/        # HTTP utilities
    │   ├── middleware/
    │   │   ├── error.go
    │   │   └── logger.go
    │   └── server.go
    ├── logger/      # Logging utilities
    │   └── logger.go
    └── utils/       # Common utilities
        ├── randomutil/
        │   └── randomutil.go
        └── timeutil/
            └── timeutil.go
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

### 2.3 Cache Framework
- Selection rationale: Flexible interface design, support multiple implementations
- Location: pkg/cache/
- Interface design:
  ```go
  type Cache interface {
      Get(ctx context.Context, key string) ([]byte, error)
      Set(ctx context.Context, key string, value []byte) error
      SetEX(ctx context.Context, key string, value []byte, expiration time.Duration) error
      SetNX(ctx context.Context, key string, value []byte, expiration time.Duration) (bool, error)
      Del(ctx context.Context, key string) error
      GetObject(ctx context.Context, key string, value interface{}) error
      SetObject(ctx context.Context, key string, value interface{}) error
      SetObjectEX(ctx context.Context, key string, value interface{}, expiration time.Duration) error
      Close() error
  }
  ```
- Main features:
  - Basic key-value operations
  - Object serialization support
  - Distributed lock capability (SetNX)
  - TTL support
  - Context support for timeout/cancellation
- Implementation: Redis (pkg/cache/redis)

### 2.4 Task Scheduling Framework
- Selection rationale: 
  - Based on robfig/cron: Go生态中最流行的定时任务库
  - 支持秒级调度精度（通过 WithSeconds 选项）
  - 支持标准的 cron 表达式
  - 提供了优雅的启动和停止机制
- Enhanced features:
  - Redis分布式锁防止任务并发执行
  - 自动的任务执行日志（开始时间、执行时长、错误信息）
  - 统一的任务注册机制（通过 Registrar 接口）
  - 任务执行状态监控（通过日志）
- Location: pkg/cron/
- Core components:
  - Job interface for task definition:
    ```go
    type Job interface {
        Run(ctx context.Context) error
    }
    ```
  - Registrar interface for job registration:
    ```go
    type Registrar interface {
        Register(c *Cron) error
    }
    ```
  - Cron manager for job scheduling and lifecycle management
  - Redis lock for preventing concurrent execution
- Usage:
  - Define jobs implementing the Job interface
  - Register jobs through Registrar
  - Support cron expressions (e.g., "*/5 * * * * *" for every 5 seconds)
  - Jobs are automatically protected by Redis locks
  - Execution logs with duration tracking

### 2.5 Dependency Injection - Wire
- Selection rationale: Compile-time DI, type-safe
- Main features: Dependency management, interface binding
- Usage: Component initialization
- Location: internal/di/
- Provider Structure:
  ```
  internal/di/
  ├── provider/
  │   └── provider.go    # Provider sets definitions
  ├── wire.go           # Wire injection specifications
  └── wire_gen.go       # Generated wire code
  ```
- Provider Categories:
  1. InfraProvider: Infrastructure dependencies
     - MySQL database
     - Redis cache
     - HTTP server
  2. RepositoryProvider: Data access layer
     - Signal repository
  3. ServiceProvider: Business service layer
     - Signal service
  4. HandlerProvider: Web handling layer
     - Router setup
  5. TaskProvider: Scheduled tasks
     - Cron manager
     - Cron jobs

### 2.6 Configuration - Viper
- Selection rationale: Standard configuration library
- Main features: Config file reading, environment variable support
- Configuration methods:
  - Environment: .env file (local development)
  - Config files: config.{env}.yaml
  - Environments: dev/prod

### 2.7 Logging - Zap
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
- Provider organization:
  1. Define providers in internal/di/provider/provider.go
  2. Group related providers into provider sets
  3. Combine all sets in ProviderSet for application initialization
- Wire usage:
  1. Define dependencies in wire.go
  2. Run wire command to generate wire_gen.go
  3. Use generated InitializeApp function to bootstrap application
