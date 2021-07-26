package main

import (
	"github.com/VladMak/ServiceAPI"
	"github.com/VladMak/ServiceAPI/pkg/handler"
	"github.com/VladMak/ServiceAPI/pkg/repository"
	"github.com/VladMak/ServiceAPI/pkg/service"
	"log"
)

/* Точка входа в программу
Определяем основные обработчики, абстракцию сервера HTTP, и запускаем все это дело на определенном порту. Подключаем обработку ошибок. Дальнейшая работа переходит в сервер HTTP, при этом туда передаются обработчики и инициализируются сразу же.
*/
func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)


	srv := new(ServiceAPI.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
