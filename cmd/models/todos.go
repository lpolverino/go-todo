package models

type Todo struct {
	Id      string `json:"id"`
	Status  bool   `json:"status"`
	Author  string `json:"author"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
