package main

import (
	"github.com/VMironiuk/sportifight-server"
	"github.com/VMironiuk/sportifight-server/pkg/handler"
	"github.com/VMironiuk/sportifight-server/pkg/repository"
	"github.com/VMironiuk/sportifight-server/pkg/service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
	"os"
)

func main() {
	if err := gotenv.Load(); err != nil {
		logrus.Fatalf("an error occured while reading env variables: %s", err.Error())
	}

	if err := initConfig(); err != nil {
		logrus.Fatalf("an error occured while reading configs: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("an error occurred while initializing db: %s\n", err.Error())
	}

	repos := repository.New(db)
	services := service.New(repos)
	handlers := handler.New(services)
	srv := new(sportifight.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		logrus.Fatalf("an error occured during server start: %s\n", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
