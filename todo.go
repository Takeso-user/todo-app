package todo_app

type TodoList struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UsersLIst struct {
	ID     int
	UserId int
	ListId int
}

type TodoItem struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type ListsItems struct {
	ID     int
	ListId int
	ItemId int
}
