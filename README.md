# go-gin-gorm-wire-viper-zap

A Go web project template integrating mainstream frameworks and best practices.

## Features

- 🚀 Web framework based on [Gin](https://github.com/gin-gonic/gin)
- 📦 Database operations with [GORM](https://gorm.io/)
- 🔧 Dependency injection using [Wire](https://github.com/google/wire)
- ⚙️ Configuration management with [Viper](https://github.com/spf13/viper)
- 📝 Logging with [Zap](https://github.com/uber-go/zap)
- 💾 Caching with Redis
- ⏰ Task scheduling with [Cron](https://github.com/robfig/cron)
  - Distributed lock support
  - Automatic logging
  - Job interface abstraction
  - Graceful start/stop
- 🔐 Separated configurations for development and production
- 📚 Standard Go project layout
- 🎯 Unified error handling and response format
- 🔄 Graceful shutdown support
- 📊 Built-in health check
- 🛠️ Rich utility packages

## Getting Started

### Prerequisites

- Go 1.20 or higher
- MySQL 8.0 or higher
- Redis 6.0 or higher
- Wire: `go install github.com/google/wire/cmd/wire@latest`

### Installation

```bash
# Clone the repository
git clone https://github.com/tommggo/go-gin-gorm-wire-viper-zap.git

# Change directory
cd go-gin-gorm-wire-viper-zap

# Initialize module (for new projects)
go mod init go-gin-gorm-wire-viper-zap

# Install dependencies
go mod tidy

# Generate dependency injection code
cd internal/di && wire
```

### Configuration

1. Copy the environment template:
```bash
cp .env.example .env
```

2. Modify the configuration in `.env`:
```env
APP_ENV=dev  # Environment: dev/prod
```

3. Update the configuration file for your environment:
- Development: `configs/config.dev.yaml`
- Production: `configs/config.prod.yaml`

### Running

```bash
# Development
go run cmd/main.go

# Production
APP_ENV=prod go run cmd/main.go
```

## Project Structure

For detailed project structure and documentation, please refer to [Framework Documentation](docs/framework.md).

## Development Guide

### Adding New API

1. Add new handler in `internal/api`
2. Register route in `internal/router/router.go`
3. Implement business logic in `internal/service`
4. Implement data access in `internal/repository`
5. Add necessary models in `internal/model`

### Adding New Cron Job

1. Implement the Job interface in `internal/cron/jobs`:
```go
type Job interface {
    Run(ctx context.Context) error
}
```

2. Register the job in cron manager with cron expression:
```go
// Support cron expressions like: "*/5 * * * *" (every 5 minutes)
// See https://pkg.go.dev/github.com/robfig/cron/v3 for expression format
cron.AddJob("job-name", "*/5 * * * *", job)
```

3. Jobs will automatically:
   - Use distributed locks to prevent concurrent execution
   - Log execution details (start/end time, duration, errors)
   - Handle panic recovery

### Error Handling

- Define error codes in `internal/errors/code.go`
- Create new errors using `errors.New`
- Errors are handled uniformly by middleware

### Logging

```go
import "go-gin-gorm-wire-viper-zap/pkg/logger"

// Examples
logger.Info("message", logger.String("key", "value"))
logger.Error("error message", logger.Err(err))
```

### Caching

```go
import "go-gin-gorm-wire-viper-zap/pkg/cache"

// Examples
cache.Set(ctx, key, value)
cache.SetEX(ctx, key, value, expiration)
cache.SetNX(ctx, key, value, expiration) // For distributed locks
```

## Deployment

### Build

```bash
# Build binary
go build -o app cmd/main.go
```

### Run

```bash
# Set environment
export APP_ENV=prod

# Run service
./app
```

## Contributing

Issues and Pull Requests are welcome.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

