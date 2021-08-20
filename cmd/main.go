package main

import (
	"github.com/VladMak/ServiceAPI"
	"github.com/VladMak/ServiceAPI/pkg/handler"
	"github.com/VladMak/ServiceAPI/pkg/repository"
	"github.com/VladMak/ServiceAPI/pkg/service"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	_"log"
	"github.com/sirupsen/logrus"
	"os"
	_"fmt"
)

const (
	DEVELOPMENT bool = true
)

var (
	port string
)

/* Точка входа в программу
Определяем основные обработчики, абстракцию сервера HTTP, и запускаем все это дело на определенном порту. Подключаем обработку ошибок. Дальнейшая работа переходит в сервер HTTP, при этом туда передаются обработчики и инициализируются сразу же.
*/
func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	// viper.GetString("port")
	// os.Getenv("PORT")
	if port = os.Getenv("PORT"); port == "" {
		port = viper.GetString("port")
	}

	srv := new(ServiceAPI.Server)
	if err := srv.Run(port, handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	if DEVELOPMENT {
		viper.AddConfigPath("configs")
	} else {
		viper.AddConfigPath("configs")
		//viper.AddConfigPath("prod_configs")
	}
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}