package service

import "github.com/PopovVA/tic-tac-toe-API/pkg/repository"

type GameList interface{}

type GameItem interface{}

type Service struct {
	GameList
	GameItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
