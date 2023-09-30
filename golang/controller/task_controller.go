package controller

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"study-app-api/model"
	"study-app-api/usecase"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// インターフェース
type ITaskController interface {
	GetOwnAllTasks(c echo.Context) error
	GetTask(c echo.Context) error
	CreateTask(c echo.Context) error
	UpdateTask(c echo.Context) error
	UpdateOwnAllTasks(c echo.Context) error
	DeleteTask(c echo.Context) error
	DeleteTaskList(c echo.Context) error
	WebSocketHandler(c echo.Context) error
	UploadImageHandler(c echo.Context) error
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
	log.Println("controller CreateTask タスク作成成功")
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

func (tc *taskController) DeleteTask(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("id")
	taskId, _ := strconv.Atoi(id)

	err := tc.tu.DeleteTask(uint(taskId), uint(userId.(float64)))
	if err != nil {
		log.Println("controller DeleteTask タスク削除エラー: ", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	log.Println("controller DeleteTask : タスク削除成功")
	return c.NoContent(http.StatusNoContent)
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
		log.Println("message: ", message)

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

func (tc *taskController) UploadImageHandler(c echo.Context) error {
	// 画像ファイルを取得する
	file, err := c.FormFile("image")
	if err != nil {
		return c.String(http.StatusBadRequest, "Failed to retrieve image")
	}

	// 画像を保存する
	savePath := SaveImage(file)

	// 画像の保存先パスを返す
	return c.String(http.StatusOK, savePath)
}

func SaveImage(fileHeader *multipart.FileHeader) string {
	file, err := fileHeader.Open()
	if err != nil {
		fmt.Println("SaveImage error", err)
		return ""
	}
	defer file.Close()

	// 保存先のフォルダを作成する
	saveDir := "images" // 保存フォルダのパス
	os.MkdirAll(saveDir, os.ModePerm)

	// ファイルの保存先のパスを作成する
	saveFilePath := filepath.Join(saveDir, fmt.Sprintf("image_%d%s", time.Now().UnixNano(), filepath.Ext(fileHeader.Filename)))

	creds := credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), "") // AWSに接続

	sess, err := session.NewSession(&aws.Config{
		Credentials:      creds,
		Region:           aws.String("ap-northeast-1"),
		Endpoint:         aws.String(os.Getenv("S3_URL")),
		S3ForcePathStyle: aws.Bool(true),
	})
	if err != nil {
		fmt.Println("SaveImage error", err)
		return ""
	}

	defer file.Close()

	// S3に画像ファイルをアップロードします。
	uploader := s3manager.NewUploader(sess) // S3にアップロード
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(os.Getenv("S3_BUCKET_NAME")),
		Key:    aws.String(saveFilePath),
		Body:   file,
	})
	if err != nil {
		fmt.Println("SaveImage error", err)
		return ""
	}

	// ファイルを保存する
	saveFile, err := os.Create(saveFilePath)
	if err != nil {
		fmt.Println("SaveImage error", err)
		return ""
	}
	defer saveFile.Close()

	// ファイルの内容を保存する
	_, err = io.Copy(saveFile, file)
	if err != nil {
		fmt.Println("SaveImage error", err)
		return ""
	}

	return saveFilePath
}
