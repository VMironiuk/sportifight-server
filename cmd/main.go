package main

import (
	"log"
	"sportifight"
)

func main() {
	srv := new(sportifight.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatalf("an error occured during server start: %s\n", err.Error())
	}
}
