// Package config 提供从环境变量加载服务配置的最小实现（骨架）。
package config

import "os"

// Config 是所有服务共用的基础配置。
type Config struct {
	ServiceName  string
	GRPCAddr     string
	HTTPAddr     string
	OTLPEndpoint string
}

// Load 从环境变量读取配置，未设置时使用默认值。
func Load(serviceName string) Config {
	return Config{
		ServiceName:  serviceName,
		GRPCAddr:     Getenv("GRPC_ADDR", ":50051"),
		HTTPAddr:     Getenv("HTTP_ADDR", ":8080"),
		OTLPEndpoint: Getenv("OTLP_ENDPOINT", "localhost:4317"),
	}
}

// Getenv 返回环境变量值，缺省时返回 def。
func Getenv(key, def string) string {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		return v
	}
	return def
}
