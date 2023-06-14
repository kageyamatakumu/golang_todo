package repository

import (
	"fmt"
	"go-todo-api/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Interface
type ITodoRepository interface {
	GetAllTodos(todos *[]model.Todo, userId uint) error
	GetTodoById(todo *model.Todo, userId uint, todoId uint) error
	CreateTodo(todo *model.Todo) error
	EditTodo(todo *model.Todo, userId uint, todoId uint) error
	DeleteTodo(userId uint, todoId uint) error
}

type todoRepository struct {
	db *gorm.DB
}

// コンストラクタ
func NewTodoRepository(db *gorm.DB) ITodoRepository {
	return &todoRepository{db}
}

// Todo一覧取得
func (tr *todoRepository) GetAllTodos(todos *[]model.Todo, userId uint) error {
	if err := tr.db.Joins("User").Where("user_id=?", userId).Find(todos).Error; err != nil {
		return err
	}
	return nil
}

// Todo取得（単体）
func (tr *todoRepository) GetTodoById(todo *model.Todo, userId uint, todoId uint) error {
	if err := tr.db.Where("id=? AND user_id=?", todoId, userId).Find(todo).Error; err != nil {
		return err
	}
	return nil
}

// Todo作成
func (tr *todoRepository) CreateTodo(todo *model.Todo) error {
	if err := tr.db.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

// Todo更新（タイトルと説明）
func (tr *todoRepository) EditTodo(todo *model.Todo, userId uint, todoId uint) error {
	result := tr.db.Model(todo).Clauses(clause.Returning{}).Where("id=? AND user_id=?", todoId, userId).Updates(model.Todo{Title: todo.Title, Description: todo.Description})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1{
		return fmt.Errorf("object does not exit")
	}
	return nil
}

// Todo削除
func (tr *todoRepository) DeleteTodo(userId uint, todoId uint) error {
	result := tr.db.Where("id=? AND user_id=?", todoId, userId).Delete(&model.Todo{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1{
		return fmt.Errorf("object does not exit")
	}
	return nil
}