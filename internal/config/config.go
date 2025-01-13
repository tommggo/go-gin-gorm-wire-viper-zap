package config

import "time"

// Config 总配置结构
type Config struct {
	App      AppConfig      `mapstructure:"app"`
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	Log      LogConfig      `mapstructure:"log"`
}

// AppConfig 应用配置
type AppConfig struct {
	Name    string `mapstructure:"name"`    // 应用名称
	Version string `mapstructure:"version"` // 应用版本
	Env     string `mapstructure:"env"`     // 环境标识：dev/prod
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Port               int           `mapstructure:"port"`              // 端口
	Mode               string        `mapstructure:"mode"`              // 运行模式：debug/release
	MaxMultipartMemory int64         `mapstructure:"max_multipart_mem"` // 上传文件最大内存
	ReadTimeout        time.Duration `mapstructure:"read_timeout"`      // 读取超时
	WriteTimeout       time.Duration `mapstructure:"write_timeout"`     // 写入超时
	MaxHeaderBytes     int           `mapstructure:"max_header_bytes"`  // 请求头最大字节数
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	DSN             string        `mapstructure:"dsn"`
	MaxOpenConns    int           `mapstructure:"max_open_conns"`
	MaxIdleConns    int           `mapstructure:"max_idle_conns"`
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime"`
}

// LogConfig 日志配置
type LogConfig struct {
	Level      string `mapstructure:"level"`       // 日志级别
	Filename   string `mapstructure:"filename"`    // 日志文件
	MaxSize    int    `mapstructure:"max_size"`    // 单个文件最大尺寸，MB
	MaxBackups int    `mapstructure:"max_backups"` // 保留旧文件最大个数
	MaxAge     int    `mapstructure:"max_age"`     // 保留旧文件最大天数
	Compress   bool   `mapstructure:"compress"`    // 是否压缩
}
