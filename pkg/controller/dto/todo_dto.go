package dto

type TodoResponse struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type TodoRequest struct {
	Title   string `validate:"min=1,max=50"`
	Content string `validate:"min=1,max=500"`
}

type TodosResponse struct {
	Todos []TodoResponse `json:"todos"`
}
