package config

import (
	"clean-code-structure/logger"
	"clean-code-structure/repository/pgsql"
	"clean-code-structure/scheduler"
)

type Application struct {
	GracefulShutdownTimeoutInSecond int  `koanf:"graceful_shutdown_timeout_in_seconds"`
	EnableProfiling                 bool `koanf:"enable_profiling"`
	ProfilingPort                   int  `koanf:"profiling_port"`
	TimeoutSeconds                  int  `koanf:"timeout_seconds"`
}

type HTTPServer struct {
	Port int `koanf:"port"`
}

type Config struct {
	Application Application      `koanf:"application"`
	HTTPServer  HTTPServer       `koanf:"http_server"`
	Scheduler   scheduler.Config `koanf:"scheduler"`
	Logger      logger.Config    `koanf:"logger"`
	PgSQL       pgsql.Config     `koanf:"pgsql"`
}
