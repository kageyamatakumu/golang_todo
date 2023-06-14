package usecase

import (
	"go-todo-api/model"
	"go-todo-api/repository"
)

// Interface
type ITodoUsecase interface {
	GetAllTodos(userId uint) ([]model.TodoResponse, error)
	GetTodoById(userId uint, todoId uint) (model.TodoResponse, error)
	CreateTodo(todo model.Todo) (model.TodoResponse, error)
	EditTodo(todo model.Todo, userId uint, todoId uint) (model.TodoResponse, error)
	DeleteTodo(userId uint, todoId uint) error
}

type todoUsecase struct {
	tr repository.ITodoRepository
}

// コンストラクタ
func NewTodoUsecase(tr repository.ITodoRepository) ITodoUsecase {
	return &todoUsecase{tr}
}

// Todo一覧取得
func (tu *todoUsecase) GetAllTodos(userId uint) ([]model.TodoResponse, error) {
	todos := []model.Todo{}
	if err := tu.tr.GetAllTodos(&todos, userId); err != nil {
		return nil, err
	}
	resTodos := []model.TodoResponse{}
	for _, v := range todos {
		t := model.TodoResponse{
			ID: v.ID,
			Title: v.Title,
			Description: v.Description,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		resTodos = append(resTodos, t)
	}
	return resTodos, nil
}

// Todo取得（単体）
func (tu *todoUsecase) GetTodoById(userId uint, todoId uint) (model.TodoResponse, error) {
	todo := model.Todo{}
	if err := tu.tr.GetTodoById(&todo, userId, todoId); err != nil {
		return model.TodoResponse{}, err
	}
	resTodo := model.TodoResponse{
		ID: todo.ID,
		Title: todo.Title,
		Description: todo.Description,
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
	}
	return resTodo, nil
}

// Todo作成
func (tu *todoUsecase) CreateTodo(todo model.Todo) (model.TodoResponse, error) {
	if err := tu.tr.CreateTodo(&todo); err != nil {
		return model.TodoResponse{}, err
	}
	resTodo := model.TodoResponse{
		ID: todo.ID,
		Title: todo.Title,
		Description: todo.Description,
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
	}
	return resTodo, nil
}

// Todo更新
func (tu *todoUsecase) EditTodo(todo model.Todo, userId uint, todoId uint) (model.TodoResponse, error) {
	if err := tu.tr.EditTodo(&todo, userId, todoId); err != nil {
		return model.TodoResponse{}, err
	}
	resTodo := model.TodoResponse{
		ID: todo.ID,
		Title: todo.Title,
		Description: todo.Description,
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
	}
	return resTodo, nil
}

// Todo削除
func (tu *todoUsecase) DeleteTodo(userId uint, todoId uint) error {
	if err := tu.tr.DeleteTodo(userId, todoId); err != nil {
		return err
	}
	return nil
}