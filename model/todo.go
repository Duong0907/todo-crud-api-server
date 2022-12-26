package todo

type Todo struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Status  string `json:"status"`
}