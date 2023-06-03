package controller

import (
	"log"
	"net/http"
	"study-app-api/usecase"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// インターフェース
type ITaskController interface {
	GetOwnAllTasks(c echo.Context) error
}

type taskController struct {
	tu usecase.ITaskUsecase
}

func NewTaskController(ut usecase.ITaskUsecase) ITaskController {
	return &taskController{ut}
}

func (tc *taskController) GetOwnAllTasks(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	tasksRes, err := tc.tu.GetOwnAllTasks(uint(userId.(float64)))
	if err != nil {
		log.Println("controller GetOwnAllTasks タスク一覧取得エラー: ", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	log.Println("controller GetOwnAllTasks : タスク一覧取得成功")
	return c.JSON(http.StatusOK, tasksRes)
}
