package repository

type GameList interface{}

type GameItem interface{}

type HistoryList interface{}

type HistoryItem interface{}

type Repository struct {
	GameList
	GameItem
	HistoryList
	HistoryItem
}

func NewRepository() *Repository {
	return &Repository{}
}
