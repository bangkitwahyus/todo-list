package handler

import (
	"encoding/json"
	"fmt"
	"todo-list/usecase"
)

type Handler struct {
	todoUseCase *usecase.UseCase
}

func Init(todo *usecase.UseCase) *Handler {
	return &Handler{
		todoUseCase: todo,
	}
}

func (h *Handler) GetTodoApi() {
	result := h.todoUseCase.Get()
	b, err := json.Marshal(result)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	fmt.Println(string(b))
}
