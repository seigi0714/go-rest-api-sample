package dto

type TodoRequest struct {
	Title   string `validate:"min=1,max=50"`
	Content string `validate:"min=1,max=500"`
}

type TodosResponse struct {
	Todos []map[string]interface{} `json:"todos"`
}
