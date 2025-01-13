# 框架说明文档

## 1. 目录结构
```
.
├── .env                  # 环境变量配置文件
├── .env.example         # 环境变量模板
├── .gitignore           # Git忽略配置
├── LICENSE             # 开源许可证
├── README.md           # 项目说明文档
├── api/                # API 定义和文档
│   └── openapi/       # OpenAPI/Swagger 文档
├── cmd/               # 主程序入口
│   └── main.go       # 主程序
├── configs/           # 配置文件目录
│   ├── config.dev.yaml   # 开发环境配置
│   └── config.prod.yaml  # 生产环境配置
├── docs/             # 文档目录
│   ├── database.sql  # 数据库脚本
│   └── framework.md  # 框架说明文档
├── internal/         # 私有应用代码
│   ├── api/         # HTTP API 处理
│   │   ├── health.go    # 健康检查
│   │   ├── request.go   # 请求模型定义
│   │   ├── response.go  # 响应处理
│   │   ├── router.go    # 路由定义
│   │   └── signal.go    # 业务处理器
│   ├── config/      # 配置处理
│   │   ├── config.go    # 配置结构定义
│   │   └── loader.go    # 配置加载器
│   ├── di/          # 依赖注入
│   │   ├── provider/    # 依赖提供者
│   │   ├── wire.go      # Wire 依赖定义
│   │   └── wire_gen.go  # Wire 生成的代码
│   ├── errors/      # 错误处理
│   │   ├── code.go     # 错误码定义
│   │   └── errors.go   # 错误处理
│   ├── model/       # 数据模型
│   ├── repository/  # 数据库操作
│   │   ├── base.go     # 基础仓储接口
│   │   └── signal.go   # 业务仓储实现
│   └── service/     # 业务逻辑
│       └── signal.go    # 业务服务实现
├── logs/            # 日志文件目录
└── pkg/             # 公共代码包
    ├── database/    # 数据库工具
    │   └── mysql.go    # MySQL 连接管理
    ├── http/        # HTTP 服务
    │   ├── middleware/ # HTTP 中间件
    │   └── server.go   # HTTP 服务器
    ├── logger/      # 日志工具
    └── utils/       # 通用工具
```

## 2. 核心框架

### 2.1 Web 框架 - Gin
- 选型原因：高性能、轻量级、社区活跃
- 主要功能：路由管理、中间件支持、参数验证
- 使用场景：API 接口、中间件处理

### 2.2 数据库框架 - GORM
- 选型原因：功能完善、社区活跃、文档丰富
- 主要功能：ORM 映射、事务管理、关联处理
- 使用场景：数据库操作、模型关系管理

### 2.3 依赖注入 - Wire
- 选型原因：编译时依赖注入、类型安全
- 主要功能：依赖管理、接口绑定
- 使用场景：组件初始化、依赖管理
- 文件位置：internal/di/

### 2.4 配置管理 - Viper
- 选型原因：配置管理标准库
- 主要功能：配置文件读取、环境变量支持
- 配置方式：
  - 环境变量：.env 文件（本地开发）
  - 配置文件：config.{env}.yaml
  - 支持环境：dev/prod

### 2.5 日志框架 - Zap
- 选型原因：高性能、结构化日志
- 主要功能：日志分级、文件轮转、结构化输出
- 使用场景：系统日志、业务日志、错误追踪
- 配置位置：pkg/logger/

## 3. 关键规范

### 3.1 项目规范
- 遵循 Go 项目标准布局
- 业务代码位于 internal 目录
- 公共代码位于 pkg 目录
- 配置文件位于 configs 目录

### 3.2 API 规范
- RESTful 风格
- 统一错误处理
- 统一响应格式
- OpenAPI/Swagger 文档

### 3.3 错误处理规范
- 错误码统一在 internal/errors 定义
- 统一错误响应格式
- 中间件统一处理错误
- 业务错误需要记录日志

### 3.4 日志规范
- 使用结构化日志
- 统一日志格式
- 按级别分类：DEBUG/INFO/WARN/ERROR/FATAL
- 日志文件位于 logs 目录

### 3.5 配置规范
- 环境变量：
  - 开发环境使用 .env 文件
  - 生产环境使用系统环境变量
- 配置文件：
  - 开发环境：config.dev.yaml
  - 生产环境：config.prod.yaml
- 敏感信息通过环境变量注入

## 4. 开发流程

### 4.1 本地开发
1. 复制 .env.example 到 .env
2. 修改 .env 中的配置
3. 确保 configs/config.dev.yaml 配置正确
4. 执行 `go run cmd/main.go`

### 4.2 生产部署
1. 准备生产环境配置 config.prod.yaml
2. 设置必要的环境变量
3. 使用 APP_ENV=prod 启动服务

### 4.3 依赖管理
- 使用 go mod 管理依赖
- 执行 wire 生成依赖注入代码：
  ```bash
  cd internal/di && wire
  ```
