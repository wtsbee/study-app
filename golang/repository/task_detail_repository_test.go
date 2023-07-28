package repository

import (
	"fmt"
	"testing"

	"study-app-api/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/DATA-DOG/go-txdb"
)

const (
	testDBUser     = "root"
	testDBPassword = "password"
	testDBName     = "study_app_db_test"
	testDBHost     = "localhost"
	testDBPort     = "3306"
)

func setupTestDB() *gorm.DB {
	// テスト用の接続情報
	PROTOCOL := fmt.Sprintf("tcp(db:%s)", testDBPort)
	// dsn := testDBUser + ":" + testDBPassword + "@tcp(" + testDBHost + ":" + testDBPort + ")/" + testDBName + "?parseTime=true"
	dsn := testDBUser + ":" + testDBPassword + "@" + PROTOCOL + "/" + testDBName + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"

	txdb.Register("txdb", "mysql", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	// テーブルの作成
	err = db.AutoMigrate(&model.User{}, &model.Task{}, &model.TaskList{}, &model.TaskDetail{})
	if err != nil {
		panic("failed to migrate table: " + err.Error())
	}

	// usersテーブルに対応するデータを追加
	user := &model.User{
		Email:    "test@example.com",
		Password: "test_password",
	}
	if err := db.Create(user).Error; err != nil {
		panic("failed to create test user: " + err.Error())
	}

	// task_listsテーブルに対応するデータを追加
	taskList := &model.TaskList{
		Name:   "Test Task List",
		UserId: user.ID,
	}
	if err := db.Create(taskList).Error; err != nil {
		panic("failed to create test task list: " + err.Error())
	}

	// tasksテーブルに対応するデータを追加
	task := &model.Task{
		Title:      "Test Task",
		UserId:     user.ID,
		TaskListId: taskList.ID,
	}
	if err := db.Create(task).Error; err != nil {
		panic("failed to create test task: " + err.Error())
	}

	// テスト用のデータベースを返す
	return db
}

func TestTaskDetailRepository(t *testing.T) {
	// テスト用のデータベースをセットアップ
	db := setupTestDB()

	defer db.Migrator().DropTable(&model.TaskDetail{})
	defer db.Migrator().DropTable(&model.Task{})
	defer db.Migrator().DropTable(&model.TaskList{})
	defer db.Migrator().DropTable(&model.User{})

	// テスト用のリポジトリを作成
	taskDetailRepo := NewTaskDetailRepository(db)

	t.Run("CreateTaskDetail", func(t *testing.T) {
		// テスト用データ
		taskId := uint(1)
		userId := uint(1)

		// タスク詳細を作成
		err := taskDetailRepo.CreateTaskDetail(taskId, userId)
		if err != nil {
			t.Fatalf("Error creating task detail: %v", err)
		}

		// タスク詳細を取得して確認
		var taskDetail model.TaskDetail
		err = db.Table("task_details").Where("task_id = ? AND user_id = ? AND deleted_at IS NULL", taskId, userId).First(&taskDetail).Error
		if err != nil {
			t.Fatalf("Error fetching task detail: %v", err)
		}

		// 最後にタスク詳細を物理削除
		defer db.Session(&gorm.Session{AllowGlobalUpdate: true}).Unscoped().Delete(&model.TaskDetail{})

		// 確認
		if taskDetail.TaskId != taskId {
			t.Errorf("Expected TaskId to be %d, but got %d", taskId, taskDetail.TaskId)
		}
		if taskDetail.UserId != userId {
			t.Errorf("Expected UserId to be %d, but got %d", userId, taskDetail.UserId)
		}
	})

	t.Run("UpdateTaskDetail", func(t *testing.T) {
		// テスト用データ
		taskId := uint(1)
		userId := uint(1)

		// タスク詳細を作成
		td := &model.TaskDetail{TaskId: taskId, UserId: userId, Detail: "Initial detail"}
		err := db.Create(td).Error
		if err != nil {
			t.Fatalf("Error creating task detail: %v", err)
		}

		// 最後にタスク詳細を物理削除
		defer db.Session(&gorm.Session{AllowGlobalUpdate: true}).Unscoped().Delete(&model.TaskDetail{})

		// タスク詳細を更新
		newDetail := "Updated detail"
		err = taskDetailRepo.UpdateTaskDetail(&model.TaskDetailRequest{ID: td.ID, Detail: newDetail}, userId)
		if err != nil {
			t.Fatalf("Error updating task detail: %v", err)
		}

		// 更新後のタスク詳細を取得して確認
		var updatedDetail model.TaskDetail
		err = db.Table("task_details").Where("id = ? AND deleted_at IS NULL", td.ID).First(&updatedDetail).Error
		if err != nil {
			t.Fatalf("Error fetching updated task detail: %v", err)
		}

		// 確認
		if updatedDetail.Detail != newDetail {
			t.Errorf("Expected Detail to be %s, but got %s", newDetail, updatedDetail.Detail)
		}
	})

	t.Run("GetTaskDetail_NotFound", func(t *testing.T) {
		// テスト用データ
		taskId := uint(999) // 存在しないタスクID
		userId := uint(1)

		// タスク詳細を取得
		var taskDetail model.TaskDetail
		err := taskDetailRepo.GetTaskDetail(&taskDetail, taskId, userId)

		// エラーがRecordNotFoundであることを確認
		if err.Error() != "RecordNotFound" {
			t.Fatalf("Expected RecordNotFound error, but got: %v", err)
		}
	})

	t.Run("GetTaskDetail_Found", func(t *testing.T) {
		// テスト用データ
		taskId := uint(1)
		userId := uint(1)

		// タスク詳細を作成
		td := &model.TaskDetail{TaskId: taskId, UserId: userId, Detail: "Initial detail"}
		err := db.Create(td).Error
		if err != nil {
			t.Fatalf("Error creating task detail: %v", err)
		}

		// タスク詳細を取得
		var taskDetail model.TaskDetail
		err = taskDetailRepo.GetTaskDetail(&taskDetail, td.TaskId, userId)
		if err != nil {
			t.Fatalf("Error fetching task detail: %v", err)
		}

		// 確認
		if taskDetail.ID != td.ID {
			t.Errorf("Expected ID to be %d, but got %d", td.ID, taskDetail.ID)
		}
		if taskDetail.TaskId != td.TaskId {
			t.Errorf("Expected TaskId to be %d, but got %d", td.TaskId, taskDetail.TaskId)
		}
		if taskDetail.UserId != td.UserId {
			t.Errorf("Expected UserId to be %d, but got %d", td.UserId, taskDetail.UserId)
		}
		if taskDetail.Detail != td.Detail {
			t.Errorf("Expected Detail to be %s, but got %s", td.Detail, taskDetail.Detail)
		}
	})
}
