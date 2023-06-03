package repository

import (
	"study-app-api/model"

	"gorm.io/gorm"
)

// インターフェース
type ITaskRepository interface {
	GetOwnAllTasks(tasks *[]model.TaskListResponse, userId uint) error
}

type taskRepository struct {
	db *gorm.DB
}

// コンストラクタ
func NewTaskRepository(db *gorm.DB) ITaskRepository {
	return &taskRepository{db}
}

func (tr *taskRepository) GetOwnAllTasks(tasks *[]model.TaskListResponse, userId uint) error {
	type Response struct {
		TaskListRank uint
		TaskListName string
		TaskRank     uint
		TaskTitle    string
		UserId       uint
	}

	res := []Response{}

	// テーブル結合してデータ取得
	err := tr.db.Table("task_lists").
		Select("task_lists.rank as task_list_rank, task_lists.name as task_list_name, tasks.rank as task_rank, tasks.title as task_title, tasks.user_id").
		Joins("JOIN tasks ON task_lists.id = tasks.task_list_id").
		Where("tasks.user_id = ?", userId).
		Order("task_lists.rank, tasks.rank").
		Find(&res).Error
	if err != nil {
		return err
	}

	var arr []model.TaskResponse
	for i, v := range res {
		taskRes := model.TaskResponse{
			Rank:  v.TaskRank,
			Title: v.TaskTitle,
		}
		arr = append(arr, taskRes)
		if i+1 == len(res) || res[i].TaskListRank != res[i+1].TaskListRank {
			taskListRes := model.TaskListResponse{
				Rank:  v.TaskListRank,
				Name:  v.TaskListName,
				Tasks: arr,
			}
			*tasks = append(*tasks, taskListRes)
			arr = nil
		}
	}

	return nil
}
