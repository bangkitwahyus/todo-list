package repository

import "todo-list/model"

type Repository struct {
}

func Init() *Repository {
	return &Repository{}
}

func (r *Repository) GetTodoStatic() model.Todo {
	return model.Todo{
		Name: "Beli PC",
		Des:  "PC yang bagus harus RGB",
	}
}
