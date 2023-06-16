package repository

import (
	"study-app-api/model"

	"gorm.io/gorm"
)

// インターフェース
type ITaskRepository interface {
	GetOwnAllTasks(tasks *[]model.TaskAndTaskListResponse, userId uint) error
	CreateTask(task *model.TaskRequest, userId uint) error
	UpdateOwnAllTasks(taskList *[]model.TaskListResponse, userId uint) error
	DeleteTaskList(taskListId uint, userId uint) error
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
		Joins("left outer join tasks on task_lists.id = tasks.task_list_id").
		Where("task_lists.user_id = ? AND task_lists.deleted_at IS NULL", userId).
		Order("task_lists.rank, tasks.rank").
		Find(&ttl).Error
	if err != nil {
		return err
	}
	return nil
}

func (tr *taskRepository) CreateTask(task *model.TaskRequest, userId uint) error {
	result := tr.db.Create(&model.Task{Title: task.Title, UserId: userId, TaskListId: task.TaskListId, Rank: task.Rank})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (tr *taskRepository) UpdateOwnAllTasks(taskList *[]model.TaskListResponse, userId uint) error {
	for index, value := range *taskList {
		rank := uint(index + 1)
		if value.ID == 0 {
			tr.db.Create(&model.TaskList{Name: value.Name, Rank: rank, UserId: userId})
		} else {
			tr.db.Updates(&model.TaskList{ID: value.ID, Rank: rank})
			for i, v := range value.Tasks {
				rank = uint(i + 1)
				tr.db.Updates(&model.Task{ID: v.ID, TaskListId: value.ID, Rank: rank})
			}
		}
	}
	return nil
}

func (tr *taskRepository) DeleteTaskList(taskListId uint, userId uint) error {
	result := tr.db.Where("id = ? AND user_id = ?", taskListId, userId).Delete(&model.TaskList{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
