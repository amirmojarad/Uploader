package main

import (
	"Uploader/conf"
	"Uploader/database"
	"Uploader/internal/controller"
	"Uploader/internal/gateway"
	"Uploader/internal/jwt-handler"
	"Uploader/internal/logger"
	"Uploader/internal/repository"
	"Uploader/internal/service"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/pkg/errors"
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

	minioGateway, err := getMinioGateway(cfg)
	if err != nil {
		return nil, err
	}

	gin.SetMode(cfg.Gin.Mode)

	ginEngine := gin.Default()

	jwtService := getJwtService(cfg)
	middleware := controller.NewMiddleware(jwtService)

	uploaderSvc := getUploaderSvc(cfg, minioGateway)

	userRepo := repository.NewUser(psql)
	authSvc := getAuthService(userRepo, cfg)

	router := ginEngine.Group("/v1/uploader")
	controller.SetupAuthRoutes(router, getAuthController(authSvc, cfg, jwtService))
	controller.SetupUploadRoutes(router, middleware, getUploaderController(cfg, uploaderSvc))
	return ginEngine, nil
}

func getUploaderSvc(cfg *conf.AppConfig, minio *gateway.Minio) *service.Uploader {
	return service.NewUploader(cfg, logger.GetLogger().WithField("name", "service.uploader"), minio)
}

func getUploaderController(cfg *conf.AppConfig, uploaderSvc *service.Uploader) *controller.Uploader {
	return controller.NewUploader(cfg, logger.GetLogger().WithField("name", "uploaderSvc.controller"), uploaderSvc)
}

func getGormDB(_ *conf.AppConfig, postgresDb *sql.DB) (*gorm.DB, error) {
	psql, err := gorm.Open(postgres.New(postgres.Config{Conn: postgresDb}), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return psql, nil
}

func getAuthService(repo *repository.User, cfg *conf.AppConfig) *service.Auth {
	return service.NewAuth(cfg, logger.GetLogger().WithField("name", "auth.service"), repo)
}

func getAuthController(svc *service.Auth, cfg *conf.AppConfig, jwtService *jwt_handler.Jwt) *controller.Auth {
	return controller.NewAuth(svc, cfg, logger.GetLogger().WithField("name", "auth.controller"), jwtService)
}

func getJwtService(cfg *conf.AppConfig) *jwt_handler.Jwt {
	return jwt_handler.NewJwt(cfg.Jwt.Secret, cfg.Jwt.ExpireDuration)
}

func getMinioGateway(cfg *conf.AppConfig) (*gateway.Minio, error) {
	client, err := minio.New(cfg.Minio.Url, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.Minio.AccessKeyID, cfg.Minio.SecretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return gateway.NewMinio(client, cfg.Minio.BucketName), nil
}
