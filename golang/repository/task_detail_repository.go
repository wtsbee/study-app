package repository

import (
	"study-app-api/model"

	"gorm.io/gorm"
)

// インターフェース
type ITaskDetailRepository interface {
	CreateTaskDetail(taskId uint, userId uint) error
}

type taskDetailRepository struct {
	db *gorm.DB
}

// コンストラクタ
func NewTaskDetailRepository(db *gorm.DB) ITaskDetailRepository {
	return &taskDetailRepository{db}
}

func (tdr *taskDetailRepository) CreateTaskDetail(taskId uint, userId uint) error {
	result := tdr.db.Create(&model.TaskDetail{TaskId: taskId, UserId: userId})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
