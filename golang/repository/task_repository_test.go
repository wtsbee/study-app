package repository

import (
	"fmt"
	"study-app-api/model"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/stretchr/testify/suite"
)

// const (
// 	testDBUser     = "root"
// 	testDBPassword = "password"
// 	testDBName     = "study_app_db_test"
// 	testDBHost     = "localhost"
// 	testDBPort     = "3306"
// )

type TaskRepositoryTestSuite struct {
	suite.Suite
	db *gorm.DB
}

func (s *TaskRepositoryTestSuite) SetupSuite() {
	s.setupTestDB()

	// テーブルの削除
	s.deleteTables()

	// テーブルの作成
	s.createTables()

	// データの作成
	s.createTestData()
}

func (s *TaskRepositoryTestSuite) SetupTest() {}

func (s *TaskRepositoryTestSuite) TearDownTest() {}

// テストスイート終了後の処理
func (s *TaskRepositoryTestSuite) TearDownSuite() {
	// テーブルの削除
	s.deleteTables()
}

func TestTaskRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(TaskRepositoryTestSuite))
}

func (s *TaskRepositoryTestSuite) setupTestDB() {
	// テスト用の接続情報
	PROTOCOL := fmt.Sprintf("tcp(db:%s)", testDBPort)
	dsn := testDBUser + ":" + testDBPassword + "@" + PROTOCOL + "/" + testDBName + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	s.db = db

}

func (s *TaskRepositoryTestSuite) createTables() {
	err := s.db.AutoMigrate(&model.User{}, &model.Task{}, &model.TaskList{}, &model.TaskDetail{})
	if err != nil {
		panic("failed to migrate table: " + err.Error())
	}
}

func (s *TaskRepositoryTestSuite) deleteTables() {
	s.db.Migrator().DropTable(&model.TaskDetail{})
	s.db.Migrator().DropTable(&model.Task{})
	s.db.Migrator().DropTable(&model.TaskList{})
	s.db.Migrator().DropTable(&model.User{})
}

var (
	User     *model.User
	TaskList *model.TaskList
	Task1    *model.Task
	Task2    *model.Task
)

func (s *TaskRepositoryTestSuite) createTestData() {
	// usersテーブルに対応するデータを追加
	User = &model.User{
		Email:    "test@example.com",
		Password: "test_password",
	}
	if err := s.db.Create(User).Error; err != nil {
		panic("failed to create test user: " + err.Error())
	}

	// task_listsテーブルに対応するデータを追加
	TaskList = &model.TaskList{
		Name:   "Test Task List",
		UserId: User.ID,
	}
	if err := s.db.Create(TaskList).Error; err != nil {
		panic("failed to create test task list: " + err.Error())
	}

	// tasksテーブルに対応するデータを追加
	Task1 = &model.Task{
		Title:      "Test Task 1",
		UserId:     User.ID,
		TaskListId: TaskList.ID,
	}
	if err := s.db.Create(Task1).Error; err != nil {
		panic("failed to create test task: " + err.Error())
	}

	Task2 = &model.Task{
		Title:      "Test Task 2",
		UserId:     User.ID,
		TaskListId: TaskList.ID,
	}
	if err := s.db.Create(Task2).Error; err != nil {
		panic("failed to create test task: " + err.Error())
	}

}

func (s *TaskRepositoryTestSuite) TestGetTask() {
	// トランザクション処理
	tx := s.db.Begin()
	defer tx.Rollback()

	taskRepo := NewTaskRepository(tx)
	task := &model.Task{}
	err := taskRepo.GetTask(task, Task1.ID, User.ID)

	s.Assert().NoError(err)
	s.Assert().Equal("Test Task 1", task.Title)
}

func (s *TaskRepositoryTestSuite) TestGetOwnAllTasks() {
	// トランザクション処理
	tx := s.db.Begin()
	defer tx.Rollback()

	taskRepo := NewTaskRepository(tx)
	taskList := []model.TaskAndTaskListResponse{}
	err := taskRepo.GetOwnAllTasks(&taskList, User.ID)

	s.Assert().NoError(err)
	s.Assert().Equal("Test Task 1", taskList[0].TaskTitle)
	s.Assert().Equal("Test Task 2", taskList[1].TaskTitle)
	s.Assert().Equal(2, len(taskList))
}

func (s *TaskRepositoryTestSuite) TestCreateTask() {
	// トランザクション処理
	tx := s.db.Begin()
	defer tx.Rollback()

	taskRepo := NewTaskRepository(tx)
	task := &model.Task{
		Title:      "New Task",
		TaskListId: TaskList.ID,
		UserId:     User.ID,
	}
	err := taskRepo.CreateTask(task, User.ID)

	s.Assert().NoError(err)
	s.Assert().NotZero(task.ID)
}

func (s *TaskRepositoryTestSuite) TestUpdateTask() {
	// トランザクション処理
	tx := s.db.Begin()
	defer tx.Rollback()

	taskRepo := NewTaskRepository(tx)
	taskReq := &model.TaskRequest{
		Title: "Updated Task",
	}
	err := taskRepo.UpdateTask(taskReq, User.ID, Task1.ID)

	s.Assert().NoError(err)

	task := &model.Task{}
	err = tx.Where("id = ?", Task1.ID).First(task).Error
	s.Assert().NoError(err)
	s.Assert().Equal("Updated Task", task.Title)
}

func (s *TaskRepositoryTestSuite) TestUpdateOwnAllTasks() {
	// トランザクション処理
	tx := s.db.Begin()
	defer tx.Rollback()

	taskRepo := NewTaskRepository(tx)
	taskList := []model.TaskListResponse{
		{
			ID:   TaskList.ID,
			Name: TaskList.Name,
			Tasks: []model.TaskResponse{
				{
					ID:    Task2.ID,
					Title: Task2.Title,
				},
				{
					ID:    Task1.ID,
					Title: Task1.Title,
				},
			},
		},
	}
	err := taskRepo.UpdateOwnAllTasks(&taskList, User.ID)

	s.Assert().NoError(err)

	var tasks []model.Task
	err = tx.Where("task_list_id = ?", TaskList.ID).Find(&tasks).Error
	s.Assert().NoError(err)
	s.Assert().Equal("Test Task 1", tasks[0].Title)
	s.Assert().Equal(uint(2), tasks[0].Rank)
	s.Assert().Equal("Test Task 2", tasks[1].Title)
	s.Assert().Equal(uint(1), tasks[1].Rank)
	s.Assert().Equal(2, len(tasks))
}

func (s *TaskRepositoryTestSuite) TestDeleteTask() {
	// トランザクション処理
	tx := s.db.Begin()
	defer tx.Rollback()

	taskRepo := NewTaskRepository(tx)
	err := taskRepo.DeleteTask(Task1.ID, User.ID)

	s.Assert().NoError(err)

	task := &model.Task{}
	err = tx.Where("id = ?", Task1.ID).First(task).Error
	fmt.Println("err:", err)
	s.Assert().Error(err)
	s.Assert().Equal("record not found", err.Error())
}

func (s *TaskRepositoryTestSuite) TestDeleteTaskList() {
	// トランザクション処理
	tx := s.db.Begin()
	defer tx.Rollback()

	taskRepo := NewTaskRepository(tx)
	err := taskRepo.DeleteTaskList(TaskList.ID, User.ID)

	s.Assert().NoError(err)

	taskList := &model.TaskList{}
	err = tx.Where("id = ?", TaskList.ID).First(taskList).Error
	s.Assert().Error(err)
	s.Assert().Equal("record not found", err.Error())

	tasks := []model.Task{}
	err = tx.Where("task_list_id = ?", TaskList.ID).Find(&tasks).Error
	s.Assert().NoError(err)
	s.Assert().Empty(tasks) // 関連するタスクも削除されていることを期待
}
