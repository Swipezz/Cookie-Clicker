package entity

type Player struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Score    int    `json:"score"`
}
