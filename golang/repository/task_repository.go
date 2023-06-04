package repository

import (
	"study-app-api/model"

	"gorm.io/gorm"
)

// インターフェース
type ITaskRepository interface {
	GetOwnAllTasks(tasks *[]model.TaskAndTaskListResponse, userId uint) error
}

type taskRepository struct {
	db *gorm.DB
}

// コンストラクタ
func NewTaskRepository(db *gorm.DB) ITaskRepository {
	return &taskRepository{db}
}

func (tr *taskRepository) GetOwnAllTasks(ttl *[]model.TaskAndTaskListResponse, userId uint) error {
	// テーブル結合してデータ取得
	err := tr.db.Table("task_lists").
		Select("task_lists.id as task_list_id, task_lists.rank as task_list_rank, task_lists.name as task_list_name, tasks.id as task_id, tasks.title as task_title, tasks.user_id").
		Joins("JOIN tasks ON task_lists.id = tasks.task_list_id").
		Where("tasks.user_id = ?", userId).
		Order("task_lists.rank, tasks.rank").
		Find(&ttl).Error
	if err != nil {
		return err
	}
	return nil
}
