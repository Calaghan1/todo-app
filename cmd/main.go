package main

import (
	"log"

	"github.com/Calaghan1/todo-app"
	"github.com/Calaghan1/todo-app/pkg/handler"
	"github.com/Calaghan1/todo-app/pkg/repository"
	"github.com/Calaghan1/todo-app/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error when read config %s", err.Error())
	}
	repos := repository.NewRepositorye()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error when starting server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("../configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}