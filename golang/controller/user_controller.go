package controller

import (
	"log"
	"net/http"
	"os"
	"study-app-api/model"
	"study-app-api/usecase"
	"time"

	"github.com/labstack/echo/v4"
)

type IUserController interface {
	SignUp(c echo.Context) error
	LogIn(c echo.Context) error
	LogOut(c echo.Context) error
	CsrfToken(c echo.Context) error
}

type userController struct {
	uu usecase.IUserUsecase
}

func NewUserController(uu usecase.IUserUsecase) IUserController {
	return &userController{uu}
}

func (uc *userController) SignUp(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		log.Println("controller SignUp リクエストデータ取得エラー: ", err)
		return c.JSON(http.StatusBadRequest, err)
	}
	userRes, err := uc.uu.SignUp(user)
	if err != nil {
		log.Println("controller SignUp サインアップエラー: ", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	log.Println("controller SignUp サインアップ成功")
	return c.JSON(http.StatusCreated, userRes)
}

func (uc *userController) LogIn(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		log.Println("controller LogIn リクエストデータ取得エラー: ", err)
		return c.JSON(http.StatusBadRequest, err)
	}
	tokenString, err := uc.uu.Login(user)
	if err != nil {
		log.Println("controller LogIn ログインエラー: ", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
	log.Println("controller LogIn ログイン成功")
	return c.NoContent(http.StatusOK)
}

func (uc *userController) LogOut(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Expires = time.Now()
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
	log.Println("controller LogOut ログアウト成功")
	return c.NoContent(http.StatusOK)
}

func (uc *userController) CsrfToken(c echo.Context) error {
	token := c.Get("csrf").(string)
	return c.JSON(http.StatusOK, echo.Map{
		"csrf_token": token,
	})
}
