package controller

import (
	"go-todo-api/model"
	"go-todo-api/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// Interface
type ITodoController interface {
	GetAllTodos(c echo.Context) error
	GetTodoById(c echo.Context) error
	CreateTodo(c echo.Context) error
	EditTodo(c echo.Context) error
	DeleteTodo(c echo.Context) error
}

type todoController struct {
	tu usecase.ITodoUsecase
}

// コンストラクタ
func NewTodoController(tu usecase.ITodoUsecase) ITodoController {
	return &todoController{tu}
}

// Todo一覧取得
func (tc *todoController) GetAllTodos(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	todosRes, err := tc.tu.GetAllTodos(uint(userId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, todosRes)
}

// Todo取得（単体）
func (tc *todoController) GetTodoById(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("todoId")
	todoId, _ := strconv.Atoi(id)
	todoRes, err := tc.tu.GetTodoById(uint(userId.(float64)), uint(todoId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, todoRes)
}

// Todo作成
func (tc *todoController) CreateTodo(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	todo := model.Todo{}
	if err := c.Bind(&todo); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	todo.UserId = uint(userId.(float64))
	todoRes, err := tc.tu.CreateTodo(todo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, todoRes)
}

// Todo更新
func (tc *todoController) EditTodo(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("todoId")
	todoId, _ := strconv.Atoi(id)

	todo := model.Todo{}
	if err := c.Bind(&todo); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	todoRes, err := tc.tu.EditTodo(todo, uint(userId.(float64)), uint(todoId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, todoRes)
}

// Todo削除
func (tc *todoController) DeleteTodo(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("todoId")
	todoId, _ := strconv.Atoi(id)

	err := tc.tu.DeleteTodo(uint(userId.(float64)), uint(todoId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

