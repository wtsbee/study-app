package controller

import (
	"fmt"
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
type ITaskController interface {
	GetOwnAllTasks(c echo.Context) error
	GetTask(c echo.Context) error
	CreateTask(c echo.Context) error
	UpdateTask(c echo.Context) error
	UpdateOwnAllTasks(c echo.Context) error
	DeleteTaskList(c echo.Context) error
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

// コンストラクタ
func NewTaskController(tu usecase.ITaskUsecase) ITaskController {
	return &taskController{tu}
}

func (tc *taskController) GetTask(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("id")
	taskId, _ := strconv.Atoi(id)

	task, err := tc.tu.GetTask(uint(taskId), uint(userId.(float64)))
	if err != nil {
		log.Println("controller GetTask タスク取得エラー: ", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	log.Println("controller GetTask : タスク取得成功")
	taskRes := model.TaskResponse{ID: task.ID, Title: task.Title}
	return c.JSON(http.StatusOK, taskRes)
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

func (tc *taskController) CreateTask(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	task := model.TaskRequest{}
	if err := c.Bind(&task); err != nil {
		log.Println("controller CreateTask リクエストデータ取得エラー: ", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	err := tc.tu.CreateTask(task, uint(userId.(float64)))
	if err != nil {
		log.Println("controller CreateTask タスク作成エラー: ", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return nil
}

func (tc *taskController) UpdateTask(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("id")
	taskId, _ := strconv.Atoi(id)

	task := model.TaskRequest{}
	if err := c.Bind(&task); err != nil {
		log.Println("controller CreateTask リクエストデータ取得エラー: ", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	err := tc.tu.UpdateTask(task, uint(userId.(float64)), uint(taskId))
	if err != nil {
		log.Println("controller CreateTask タスク更新エラー: ", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	log.Println("controller CreateTask : タスク更新成功")
	return nil
}

func (tc *taskController) UpdateOwnAllTasks(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	taskList := []model.TaskListResponse{}
	if err := c.Bind(&taskList); err != nil {
		log.Println("controller UpdateOwnAllTasks リクエストデータ取得エラー: ", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	tc.tu.UpdateOwnAllTasks(taskList, uint(userId.(float64)))
	return nil
}

func (tc *taskController) DeleteTaskList(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("id")
	taskListId, _ := strconv.Atoi(id)

	err := tc.tu.DeleteTaskList(uint(taskListId), uint(userId.(float64)))
	if err != nil {
		log.Println("controller DeleteTaskList タスクリスト削除エラー: ", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	log.Println("controller DeleteTaskList : タスクリスト削除成功")
	return c.NoContent(http.StatusNoContent)
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
		fmt.Println("message: ", message)

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
