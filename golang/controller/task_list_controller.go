package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"study-app-api/model"
	"study-app-api/usecase"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// インターフェース
type ITaskListController interface {
	UpdateTaskList(c echo.Context) error
}

type taskListController struct {
	tlu usecase.ITaskListUsecase
}

// コンストラクタ
func NewTaskListController(tlu usecase.ITaskListUsecase) ITaskListController {
	return &taskListController{tlu}
}

func (tlc *taskListController) UpdateTaskList(c echo.Context) error {
	fmt.Println("aaaaaaaaaaaaaaaaaaa")
	user := c.Get("user").(*jwt.Token)
	fmt.Println("bbbbbbbbbbbb")
	claims := user.Claims.(jwt.MapClaims)
	fmt.Println("cccccccccccccc")
	userId := claims["user_id"]
	id := c.Param("id")
	taskId, _ := strconv.Atoi(id)

	taskList := model.TaskListRequest{}
	if err := c.Bind(&taskList); err != nil {
		log.Println("controller UpdateTaskList リクエストデータ取得エラー: ", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	err := tlc.tlu.UpdateTaskList(taskList, uint(userId.(float64)), uint(taskId))
	if err != nil {
		log.Println("controller UpdateTaskList タスクリスト更新エラー: ", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	log.Println("controller UpdateTaskList : タスクリスト更新成功")
	return nil
}
