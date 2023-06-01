package repository

import (
	"study-app-api/model"

	"gorm.io/gorm"
)

// インターフェース
type ITaskRepository interface {
	GetOwnAllTasks(tasks *[]model.Task, userId uint) error
}

type taskRepository struct {
	db *gorm.DB
}

// コンストラクタ
func NewTaskRepository(db *gorm.DB) ITaskRepository {
	return &taskRepository{db}
}

func (tr *taskRepository) GetOwnAllTasks(tasks *[]model.Task, userId uint) error {
	if err := tr.db.Joins("User").Where("user_id=?", userId).Order("created_at").Find(tasks).Error; err != nil {
		return err
	}
	return nil
}
