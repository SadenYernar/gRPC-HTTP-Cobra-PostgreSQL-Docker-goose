package config

import (
	"fmt"
	"os"
	"runtime"
	"strconv"

	"github.com/SadenYernar/gRPC-HTTP-Cobra-PostgreSQL-Docker-goose.git/internal/logger"
	"github.com/joho/godotenv"
)

const (
	defaultHttpPort = ":80"
	defaultGrpcPort = ":81"
)

type Config struct {
	HttpPort           string
	GrpcPort           string
	DSN                string
	DbSchema           string
	MigrationPath      string
	MinConns           int32
	MaxConns           int32
	WorkerPoolSize     int
	ExternalServiceUrl string
}

func Get() *Config {
	if err := godotenv.Load(); err != nil {
		if check := os.IsNotExist(err); !check {
			logger.Errorf("failed to load env vars: %s", err)
		}
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s database=%s user=%s password=%s",
		getEnv("DB_HOST", ""),
		getEnv("DB_PORT", ""),
		getEnv("DB_DATABASE", ""),
		getEnv("DB_USER", ""),
		getEnv("DB_PASSWORD", ""),
	)

	return &Config{
		HttpPort:           getEnv("HTTP_PORT", defaultHttpPort),
		GrpcPort:           getEnv("GRPC_PORT", defaultGrpcPort),
		DSN:                dsn,
		DbSchema:           getEnv("DB_SCHEMA", ""),
		MinConns:           16,
		MaxConns:           16,
		MigrationPath:      getEnv("MIGRATION_PATH", "migrations"),
		WorkerPoolSize:     getEnvAsInt("WORKER_POOL_SIZE", runtime.NumCPU()),
		ExternalServiceUrl: getEnv("EXTERNAL_SERVICE_URL", ""),
	}

}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	value := getEnv(key, "")
	if v, e := strconv.Atoi(value); e == nil {
		return v
	}

	return defaultValue
}
