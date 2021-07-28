package main

import (
	_"github.com/lib/pq"
	"github.com/spf13/viper"
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
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host: viper.GetString("db.host"),
		Port: viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName: viper.GetString("db.dbname"),
		SSLMode: viper.GetString("db.sslmode"),
	})

	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)


	srv := new(ServiceAPI.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}


func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}