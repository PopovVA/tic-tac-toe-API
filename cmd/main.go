package main

import (
	"log"

	tictactoeapi "github.com/PopovVA/tic-tac-toe-API"
	"github.com/PopovVA/tic-tac-toe-API/pkg/handler"
	"github.com/PopovVA/tic-tac-toe-API/pkg/repository"
	"github.com/PopovVA/tic-tac-toe-API/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	port := viper.GetString("port")
	if len(port) == 0 {
		log.Fatal("error reading the port")
	}

	srv := new(tictactoeapi.Server)
	if err := srv.Run(port, handlers.InitRoutes()); err != nil {
		log.Fatalf("unable to run the server, %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
