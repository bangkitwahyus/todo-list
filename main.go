package main

import (
	"todo-list/handler"
	"todo-list/repository"
	"todo-list/usecase"
)

func bootstrap() {
	repo := repository.Init()
	uc := usecase.Init(repo)
	hdl := handler.Init(uc)
	hdl.GetTodoApi()

}

func main() {
	bootstrap()
}
