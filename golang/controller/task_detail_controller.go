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
type ITaskDetailController interface {
	GetTaskDetail(c echo.Context) error
	UpdateTaskDetail(c echo.Context) error
}

type taskDetailController struct {
	tdu usecase.ITaskDetailUsecase
}

// コンストラクタ
func NewTaskDetailController(tdu usecase.ITaskDetailUsecase) ITaskDetailController {
	return &taskDetailController{tdu}
}

func (tdc *taskDetailController) GetTaskDetail(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("id")
	taskId, _ := strconv.Atoi(id)

	taskDetail, err := tdc.tdu.GetTaskDetail(uint(taskId), uint(userId.(float64)))
	if err != nil {
		if err.Error() == "RecordNotFound" {
			log.Println("controller GetTaskDetail : RecordNotFound")
			return c.JSON(http.StatusOK, err.Error())
		}
		log.Println("controller GetTaskDetail タスク詳細取得エラー: ", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	log.Println("controller GetTaskDetail : タスク詳細取得成功")
	taskDetailRes := model.TaskDetailResponse{ID: taskDetail.ID, Detail: taskDetail.Detail}
	return c.JSON(http.StatusOK, taskDetailRes)
}

func (tdc *taskDetailController) UpdateTaskDetail(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	taskDetail := model.TaskDetailRequest{}
	fmt.Println("taskDetail:", taskDetail)
	if err := c.Bind(&taskDetail); err != nil {
		log.Println("controller UpdateTaskDetail リクエストデータ取得エラー: ", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	tdc.tdu.UpdateTaskDetail(taskDetail, uint(userId.(float64)))
	return nil
}
