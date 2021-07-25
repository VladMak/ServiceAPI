package main

import (
	"github.com/VladMak/ServiceAPI"
	"github.com/VladMak/ServiceAPI/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(ServiceAPI.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
