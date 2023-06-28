package repository

import (
	"errors"
	"study-app-api/model"

	"gorm.io/gorm"
)

// インターフェース
type ITaskDetailRepository interface {
	GetTaskDetail(taskDetail *model.TaskDetail, taskId uint, userId uint) error
	CreateTaskDetail(taskId uint, userId uint) error
}

type taskDetailRepository struct {
	db *gorm.DB
}

// コンストラクタ
func NewTaskDetailRepository(db *gorm.DB) ITaskDetailRepository {
	return &taskDetailRepository{db}
}

func (tdr *taskDetailRepository) GetTaskDetail(taskDetail *model.TaskDetail, taskId uint, userId uint) error {
	err := tdr.db.Table("task_details").Where("task_id = ? AND user_id = ? AND deleted_at IS NULL", taskId, userId).First(&taskDetail).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("RecordNotFound")
	}
	if err != nil {
		return err
	}
	return nil
}

func (tdr *taskDetailRepository) CreateTaskDetail(taskId uint, userId uint) error {
	result := tdr.db.Create(&model.TaskDetail{TaskId: taskId, UserId: userId})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
