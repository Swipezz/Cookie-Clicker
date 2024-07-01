package entity

type Player struct {
	name     string `json:"username"`
	password string `json:"password"`
	score    int    `json:"score"`
}
