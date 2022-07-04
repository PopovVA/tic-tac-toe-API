package tictactoeapi

type Player struct {
	Id       int    `json:"-"`
	Username string `json:"username"`
}
