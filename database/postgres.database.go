package database

import (
	"Uploader/conf"
	"database/sql"
	"fmt"
)

func ConnectToPostgres(cfg *conf.AppConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.PostgresDatabase.Host,
		cfg.PostgresDatabase.Port,
		cfg.PostgresDatabase.Username,
		cfg.PostgresDatabase.Password,
		cfg.PostgresDatabase.Name,
		cfg.PostgresDatabase.SslMode,
	)

	sqlDB, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(cfg.PostgresDatabase.MaxIdleConnections)
	sqlDB.SetMaxOpenConns(cfg.PostgresDatabase.MaxOpenConnections)
	sqlDB.SetConnMaxLifetime(cfg.PostgresDatabase.ConnectionMaxLifetime)

	return sqlDB, nil
}
