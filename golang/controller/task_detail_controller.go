package controller

import (
	"log"
	"net/http"
	"strconv"
	"study-app-api/model"
	"study-app-api/usecase"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

// インターフェース
type ITaskDetailController interface {
	GetTaskDetail(c echo.Context) error
	CreateTaskDetail(c echo.Context) error
	UpdateTaskDetail(c echo.Context) error
	WebSocketHandlerForTaskDetail(c echo.Context) error
}

type taskDetailController struct {
	tdu usecase.ITaskDetailUsecase
}

var (
	upgraderForTaskDetail = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	clientsForTaskDetail = make(map[*websocket.Conn]bool)
)

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
	taskDetailRes := model.TaskDetailResponse{ID: taskDetail.ID, Detail: taskDetail.Detail, TaskId: taskDetail.TaskId}
	return c.JSON(http.StatusOK, taskDetailRes)
}

func (tdc *taskDetailController) CreateTaskDetail(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	taskId, _ := strconv.Atoi(c.Param("id"))
	err := tdc.tdu.CreateTaskDetail(uint(taskId), uint(userId.(float64)))
	if err != nil {
		log.Println("controller CreateTaskDetail タスク詳細作成エラー: ", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	log.Println("controller CreateTaskDetail タスク詳細作成成功")
	return nil
}

func (tdc *taskDetailController) UpdateTaskDetail(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	taskDetail := model.TaskDetailRequest{}
	if err := c.Bind(&taskDetail); err != nil {
		log.Println("controller UpdateTaskDetail リクエストデータ取得エラー: ", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	err := tdc.tdu.UpdateTaskDetail(taskDetail, uint(userId.(float64)))
	if err != nil {
		log.Println("controller UpdateTaskDetail タスク詳細更新エラー: ", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	log.Println("controller UpdateTaskDetail : タスク詳細更新成功")
	return nil
}

func (tdc *taskDetailController) WebSocketHandlerForTaskDetail(c echo.Context) error {
	conn, err := upgraderForTaskDetail.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Println("WebSocketHandlerForTaskDetail Upgrade error:", err)
		return err
	}
	defer conn.Close()

	clientsForTaskDetail[conn] = true

	for {
		// クライアントからメッセージを受信
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("WebSocketHandlerForTaskDetail Read error:", err)
			break
		}
		log.Println("message: ", message)

		// 受信したメッセージを全てのクライアントに送信
		for client := range clientsForTaskDetail {
			err = client.WriteMessage(messageType, message)
			if err != nil {
				log.Println("WebSocketHandlerForTaskDetail Write error:", err)
				client.Close()
				delete(clientsForTaskDetail, client)
			}
		}
	}
	return nil
}
