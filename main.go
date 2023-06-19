package main

import (
	"go-todo-api/controller"
	"go-todo-api/db"
	"go-todo-api/repository"
	"go-todo-api/router"
	"go-todo-api/usecase"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserReposiotry(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)
	todoRepository := repository.NewTodoRepository(db)
	todoUsercase := usecase.NewTodoUsecase(todoRepository)
	todoController := controller.NewTodoController(todoUsercase)
	e := router.NewRouter(userController, todoController)
	e.Logger.Fatal(e.Start(":8080"))
}