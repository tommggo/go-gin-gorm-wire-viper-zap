# go-gin-gorm-wire-viper-zap

A Go web project template integrating mainstream frameworks and best practices.

## Features

- 🚀 Web framework based on [Gin](https://github.com/gin-gonic/gin)
- 📦 Database operations with [GORM](https://gorm.io/)
- 🔧 Dependency injection using [Wire](https://github.com/google/wire)
- ⚙️ Configuration management with [Viper](https://github.com/spf13/viper)
- 📝 Logging with [Zap](https://github.com/uber-go/zap)
- 🔐 Separated configurations for development and production
- 📚 Standard Go project layout
- 🎯 Unified error handling and response format

## Getting Started

### Prerequisites

- Go 1.20 or higher
- MySQL 8.0 or higher
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
2. Register route in `internal/api/router.go`
3. Implement business logic in `internal/service`
4. Implement data access in `internal/repository`

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

