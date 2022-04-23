package main

import (
	"log"
	"sportifight"
	"sportifight/handler"
)

func main() {
	handlers := handler.New()
	srv := new(sportifight.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("an error occured during server start: %s\n", err.Error())
	}
}
