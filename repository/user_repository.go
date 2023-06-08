package repository

import (
	"fmt"
	"go-todo-api/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Interface
type IUserRepository interface {
	CreateUser(user *model.User) error
	GetUserByEmail(user *model.User, email string) error
	ReNameUserName(user *model.User, userId uint) error
}

type userRepository struct {
	db *gorm.DB
}

// コンストラクタ
func NewUserReposiotry(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

// User作成
func (ur *userRepository) CreateUser(user *model.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// User検索
func (ur *userRepository) GetUserByEmail(user *model.User, email string) error {
	if err := ur.db.Where("email=?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

// Userの名前更新
func (ur *userRepository) ReNameUserName(user *model.User, userId uint) error {
	result := ur.db.Model(user).Clauses(clause.Returning{}).Where("user_id=?", userId).Update("name", user.Name)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}

	return nil
}

