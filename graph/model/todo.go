package model

type Todo struct {
	ID     int    `json:"id"`
	Text   string `json:"text"`
	Done   bool   `json:"done"`
	UserId int    `json:"userId"`
	User   *User  `json:"user"`
}
