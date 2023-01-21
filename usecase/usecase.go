package usecase

import (
	"todo-list/model"
	"todo-list/repository"
)

type UseCase struct {
	todoRepo *repository.Repository
}

func Init(todo *repository.Repository) *UseCase {
	return &UseCase{
		todoRepo: todo,
	}
}

func (u *UseCase) Get() model.Todo {
	return u.todoRepo.GetTodoStatic()
}
