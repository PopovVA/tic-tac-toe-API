package service

import "github.com/PopovVA/tic-tac-toe-API/pkg/repository"

type GameList interface{}

type GameItem interface{}

type HistoryList interface{}

type HistoryItem interface{}

type Service struct {
	GameList
	GameItem
	HistoryList
	HistoryItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
