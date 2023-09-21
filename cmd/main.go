package main

import (
	"fmt"

	"Uploader/conf"
	"Uploader/database"
	"Uploader/internal/controller"
	"Uploader/internal/jwt-handler"
	"Uploader/internal/logger"
	"Uploader/internal/repository"
	"Uploader/internal/service"
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	_ "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	if err := runServer(); err != nil {
		logger.GetLogger().Fatalf("failed to run server %v", err)
	}
}

func runServer() error {
	cfg, err := conf.NewAppConfig()
	if err != nil {
		return err
	}

	postgresDb, err := database.ConnectToPostgres(cfg)
	if err != nil {
		return err
	}

	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	if err := goose.Up(postgresDb, cfg.PostgresDatabase.MigrationPath); err != nil {
		return err
	}

	routerEngine, err := setupRouter(cfg, postgresDb)
	if err != nil {
		return err
	}

	logger.GetLogger().Infof("running server on port %s ...", cfg.App.Port)

	return routerEngine.Run(fmt.Sprintf(":%s", cfg.App.Port))
}

func setupRouter(cfg *conf.AppConfig, postgresDb *sql.DB) (*gin.Engine, error) {
	psql, err := getGormDB(cfg, postgresDb)
	if err != nil {
		return nil, err
	}

	gin.SetMode(cfg.Gin.Mode)

	ginEngine := gin.Default()

	jwtService := getJwtService(cfg)

	userRepo := repository.NewUser(psql)
	authSvc := getAuthService(userRepo, *jwtService, cfg)

	router := ginEngine.Group("/v1/uploader")
	controller.SetupAuthRoutes(router, getAuthController(authSvc, cfg))

	return ginEngine, nil
}

func getGormDB(_ *conf.AppConfig, postgresDb *sql.DB) (*gorm.DB, error) {
	psql, err := gorm.Open(postgres.New(postgres.Config{Conn: postgresDb}), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return psql, nil
}

func getAuthService(repo *repository.User, jwtService jwt_handler.Jwt, cfg *conf.AppConfig) *service.Auth {
	return service.NewAuth(cfg, logger.GetLogger().WithField("name", "auth.service"), repo, jwtService)
}

func getAuthController(svc *service.Auth, cfg *conf.AppConfig) *controller.Auth {
	return controller.NewAuth(svc, cfg, logger.GetLogger().WithField("name", "auth.controller"))
}

func getJwtService(cfg *conf.AppConfig) *jwt_handler.Jwt {
	return jwt_handler.NewJwt(cfg.Jwt.Secret, cfg.Jwt.ExpireDuration)
}
