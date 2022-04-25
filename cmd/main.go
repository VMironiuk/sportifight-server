package main

import (
	"log"

	"github.com/VMironiuk/sportifight-server"
	"github.com/VMironiuk/sportifight-server/pkg/handler"
	"github.com/VMironiuk/sportifight-server/pkg/service"
)

func main() {
	services := service.New()
	handlers := handler.New(services)
	srv := new(sportifight.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("an error occured during server start: %s\n", err.Error())
	}
}
