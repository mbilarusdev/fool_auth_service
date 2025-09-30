package models

type Player struct {
	PlayerID string `json:"player_id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
