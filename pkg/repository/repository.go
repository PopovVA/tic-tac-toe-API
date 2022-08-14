package repository

type GameList interface{}

type GameItem interface{}

type Repository struct {
	GameList
	GameItem
}

func NewRepository() *Repository {
	return &Repository{}
}
