package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
	"src/internal/api/handler"
	"src/internal/db"
	"src/internal/repository"
	"src/internal/service"
	"src/server"

	"github.com/spf13/viper"
)

// @openapi: 3.0.0
// @title        API документация
// @version      1.0
// @description  Это API для управления пользователями, товарами и поставщиками.
// @host         localhost:5000
// @BasePath     /api/v1
func main() {
	gin.SetMode(gin.ReleaseMode)

	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env file: %s", err.Error())
	}

	postgresDb, err := db.NewPostgresDB(db.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("PostgresPassword"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("failed initializing db: %s", err.Error())
	}

	repos := repository.NewRepositore(postgresDb)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
