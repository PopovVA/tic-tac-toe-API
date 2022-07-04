package main

import (
	"log"

	tictactoeapi "github.com/PopovVA/tic-tac-toe-API"
	"github.com/PopovVA/tic-tac-toe-API/pkg/handler"
	"github.com/PopovVA/tic-tac-toe-API/pkg/repository"
	"github.com/PopovVA/tic-tac-toe-API/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(tictactoeapi.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("Unable to run the server, %s", err.Error())
	}
}
