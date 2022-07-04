package main

import (
	"log"

	tictactoeapi "github.com/PopovVA/tic-tac-toe-API"
	"github.com/PopovVA/tic-tac-toe-API/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(tictactoeapi.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("Unable to run the server, %s", err.Error())
	}
}
