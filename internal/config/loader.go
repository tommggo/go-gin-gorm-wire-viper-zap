package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

// LoadConfig 加载配置
func LoadConfig() *Config {
	v := viper.New() // 创建新的 viper 实例，避免全局状态

	// 1. 设置环境变量相关配置
	v.SetEnvPrefix("APP")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	// 2. 加载 .env 文件到系统环境变量
	if err := godotenv.Load(); err != nil {
		panic(fmt.Sprintf("load .env failed: %v", err))
	}

	// 3. 获取环境配置
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev" // 默认为开发环境
	}
	if env != "dev" && env != "prod" {
		panic(fmt.Sprintf("invalid APP_ENV value: %s, must be dev or prod", env))
	}

	// 4. 设置配置文件
	v.SetConfigName(fmt.Sprintf("config.%s", env))
	v.SetConfigType("yaml")
	v.AddConfigPath("./configs/")

	// 5. 读取配置文件
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("read config failed: %v", err))
	}

	// 6. 解析到结构体
	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		panic(fmt.Sprintf("unmarshal config failed: %v", err))
	}

	return &cfg
}
