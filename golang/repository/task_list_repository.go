package repository

import (
	"study-app-api/model"

	"gorm.io/gorm"
)

// インターフェース
type ITaskListRepository interface {
	UpdateTaskList(taskList *model.TaskListRequest, userId uint, taskListId uint) error
}

type taskListRepository struct {
	db *gorm.DB
}

// コンストラクタ
func NewTaskListRepository(db *gorm.DB) ITaskListRepository {
	return &taskListRepository{db}
}

func (tlr *taskListRepository) UpdateTaskList(taskList *model.TaskListRequest, userId uint, taskListId uint) error {
	result := tlr.db.Model(&model.TaskList{}).Where("id=? AND user_id=?", taskListId, userId).Update("name", taskList.Name)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
