package main

import (
	"github.com/VladMak/ServiceAPI"
)

func main() {
	srv := new(todo.Server)
	if err := srv.Run("8080"); err != nil {
		log.Fatalf(format: "error occured while running http server: %s", err.Error())
	}
}