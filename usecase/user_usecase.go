package usecase

import (
	"go-todo-api/model"
	"go-todo-api/repository"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// Interface
type IUserUsecase interface {
	SignUp(user model.User) (model.UserResponse, error)
	Login(user model.User) (string, error)
	ReNameUserName(user model.User, userId uint) (model.UserResponse, error)
}

type userUsecase struct {
	ur repository.IUserRepository
}

// コンストラクタ
func NewUserUsecase(ur repository.IUserRepository) IUserUsecase {
	return &userUsecase{ur}
}

// サインアップ → User作成
func (uu *userUsecase) SignUp(user model.User) (model.UserResponse, error) {
	newUser := model.User{Email: user.Email, Password: user.Password, Name: user.Name}
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return model.UserResponse{}, err
	}
	resUser := model.UserResponse{
		ID: newUser.ID,
		Email: newUser.Email,
		Name: newUser.Name,
	}
	return resUser, nil
}

// ログイン
func (uu *userUsecase) Login(user model.User) (string, error) {
	storedUser := model.User{}
	if err := uu.ur.GetUserByEmail(&storedUser, user.Email); err != nil {
		return "", err
	}
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		return "", nil
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp": time.Now().Add(time.Hour * 12).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", nil
	}
	return tokenString, nil
}

// Userの名前更新
func (uc *userUsecase) ReNameUserName(user model.User, userId uint) (model.UserResponse, error) {
	if err := uc.ur.ReNameUserName(&user, userId); err != nil {
		return model.UserResponse{}, err
	}
	resUser := model.UserResponse{
		ID: user.ID,
		Email: user.Email,
		Name: user.Name,
	}
	return resUser, nil
}