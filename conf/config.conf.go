package conf

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type AppConfig struct {
	App struct {
		Port string
		Name string
		Env  string
	}

	Gin struct {
		Mode string
	}

	PostgresDatabase struct {
		Database
	}

	Jwt struct {
		Secret         string
		ExpireDuration int
	}

	Minio struct {
		AccessKeyID     string
		BucketName      string
		SecretAccessKey string
		Url             string
		Ssl             bool
	}
}

type Database struct {
	Username              string
	Password              string
	Host                  string
	Name                  string
	Port                  uint64
	SslMode               string
	Timezone              string
	MaxIdleConnections    int
	MaxOpenConnections    int
	ConnectionMaxLifetime time.Duration
	MigrationPath         string
}

func NewAppConfig() (*AppConfig, error) {
	cfg := AppConfig{}

	setGinConfigs(&cfg)
	setAppConfigs(&cfg)

	if err := setPostgresDbConfig(&cfg); err != nil {
		return nil, err
	}

	if err := setJwtConfigs(&cfg); err != nil {
		return nil, err
	}

	setMinioConfigs(&cfg)

	return &cfg, nil
}

func setGinConfigs(cfg *AppConfig) {
	cfg.Gin.Mode = os.Getenv("GIN_MODE")
}

func setAppConfigs(cfg *AppConfig) {
	cfg.App.Port = os.Getenv("APP_PORT")
	cfg.App.Name = os.Getenv("APP_NAME")
	cfg.App.Env = os.Getenv("APP_ENV")
}

func setJwtConfigs(cfg *AppConfig) error {
	expireDuration, err := envConvertor("JWT_EXPIRE_DURATION", func(v string) (int64, error) {
		return strconv.ParseInt(v, 10, 32)
	})

	if err != nil {
		return err
	}

	cfg.Jwt.ExpireDuration = int(expireDuration)
	cfg.Jwt.Secret = os.Getenv("JWT_SECRET")

	return nil
}

func setPostgresDbConfig(cfg *AppConfig) error {
	cfg.PostgresDatabase.Username = os.Getenv("POSTGRES_DATABASE_USERNAME")
	cfg.PostgresDatabase.Password = os.Getenv("POSTGRES_DATABASE_PASSWORD")
	cfg.PostgresDatabase.Host = os.Getenv("POSTGRES_DATABASE_HOST")
	cfg.PostgresDatabase.Name = os.Getenv("POSTGRES_DATABASE_NAME")

	port, err := envConvertor("POSTGRES_DATABASE_PORT", func(v string) (uint64, error) {
		return strconv.ParseUint(v, 10, 32)
	})
	if err != nil {
		return err
	}

	cfg.PostgresDatabase.Port = port

	cfg.PostgresDatabase.SslMode = os.Getenv("POSTGRES_DATABASE_SSLMODE")
	cfg.PostgresDatabase.Timezone = os.Getenv("POSTGRES_DATABASE_TIMEZONE")

	maxConn, err := envConvertor("POSTGRES_DATABASE_MAX_OPEN_CONN", strconv.Atoi)
	if err != nil {
		return err
	}

	cfg.PostgresDatabase.MaxOpenConnections = maxConn

	maxIdle, err := envConvertor("POSTGRES_DATABASE_MAX_IDLE_CONN", strconv.Atoi)
	if err != nil {
		return err
	}

	cfg.PostgresDatabase.MaxIdleConnections = maxIdle

	connMaxLif, err := envConvertor("POSTGRES_DATABASE_CONN_MAX_LIFETIME", time.ParseDuration)
	if err != nil {
		return err
	}

	cfg.PostgresDatabase.ConnectionMaxLifetime = connMaxLif

	cfg.PostgresDatabase.MigrationPath = os.Getenv("POSTGRES_DATABASE_MIGRATION_PATH")

	return nil
}

func setMinioConfigs(cfg *AppConfig) {
	cfg.Minio.AccessKeyID = os.Getenv("GATEWAY_MINIO_ACCESS_KEY_ID")
	cfg.Minio.Url = os.Getenv("GATEWAY_MINIO_URL")
	cfg.Minio.SecretAccessKey = os.Getenv("GATEWAY_MINIO_SECRET_ACCESS_KEY")
	cfg.Minio.BucketName = os.Getenv("GATEWAY_MINIO_BUCKET_NAME")
}

func envConvertor[T any](envKey string, converter func(v string) (T, error)) (T, error) {
	value := os.Getenv(envKey)

	result, err := converter(value)
	if err != nil {
		var noop T

		return noop, fmt.Errorf("%s is not a valid value for %s, %w", value, envKey, err)
	}

	return result, nil
}
