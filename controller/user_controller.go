package controller

import (
	"go-todo-api/model"
	"go-todo-api/usecase"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// Interface
type IUserController interface {
	SignUp(e echo.Context) error
	LogIn(e echo.Context) error
	LogOut(e echo.Context) error
	ReName(e echo.Context) error
}

type userController struct {
	uu usecase.IUserUsecase
}

// コンストラクタ
func NewUserController(uu usecase.IUserUsecase) IUserController {
	return &userController{uu}
}

// サインアップ
func (uc *userController) SignUp(e echo.Context) error {
	user := model.User{}
	if err := e.Bind(&user); err != nil {
		return err
	}
	userRes, err := uc.uu.SignUp(user)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusCreated, userRes)
}

// Userの名前更新
func (uc *userController) ReName(e echo.Context) error {
	userReq := e.Get("user").(*jwt.Token)
	claims := userReq.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	user := model.User{}
	if err := e.Bind(&user); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	userRes, err := uc.uu.ReNameUserName(user, uint(userId.(float64)))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, userRes)
}

// ログイン
func (uc *userController) LogIn(e echo.Context) error {
	user := model.User{}
	if err := e.Bind(&user); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	tokenString, err := uc.uu.Login(user)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	// cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteDefaultMode
	// cookie.SameSite = http.SameSiteNoneMode
	e.SetCookie(cookie)

	return e.NoContent(http.StatusOK)
}

// ログアウト
func (uc *userController) LogOut(e echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Expires = time.Now()
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	// cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteDefaultMode
	// cookie.Sa = http.SameSiteNoneMode
	e.SetCookie(cookie)

	return e.NoContent(http.StatusOK)
}

