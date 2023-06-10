package controller

import (
	"log"
	"net/http"
	"study-app-api/model"
	"study-app-api/usecase"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

// インターフェース
type ITaskController interface {
	GetOwnAllTasks(c echo.Context) error
	UpdateOwnAllTasks(c echo.Context) error
	WebSocketHandler(c echo.Context) error
}

type taskController struct {
	tu usecase.ITaskUsecase
}

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	clients = make(map[*websocket.Conn]bool)
)

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

func (tc *taskController) UpdateOwnAllTasks(c echo.Context) error {
	taskList := []model.TaskListResponse{}
	if err := c.Bind(&taskList); err != nil {
		log.Println("controller UpdateOwnAllTasks リクエストデータ取得エラー: ", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	tc.tu.UpdateOwnAllTasks(taskList)
	return nil
}

func (tc *taskController) WebSocketHandler(c echo.Context) error {
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Println("websocketHandler Upgrade error:", err)
		return err
	}
	defer conn.Close()

	clients[conn] = true

	for {
		// クライアントからメッセージを受信
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("websocketHandler Read error:", err)
			break
		}

		// 受信したメッセージを全てのクライアントに送信
		for client := range clients {
			err = client.WriteMessage(messageType, message)
			if err != nil {
				log.Println("websocketHandler Write error:", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
	return nil
}
