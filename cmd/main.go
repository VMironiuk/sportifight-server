package main

import (
	"github.com/VMironiuk/sportifight-server"
	"github.com/VMironiuk/sportifight-server/pkg/handler"
	"github.com/VMironiuk/sportifight-server/pkg/repository"
	"github.com/VMironiuk/sportifight-server/pkg/service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {
	db, err := repository.NewPostgresDB()
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
