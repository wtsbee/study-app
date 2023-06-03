package usecase

import (
	"study-app-api/model"
	"study-app-api/repository"
)

// // インターフェース
type ITaskUsecase interface {
	GetOwnAllTasks(userId uint) ([]model.TaskListResponse, error)
}

type taskUsecase struct {
	tr repository.ITaskRepository
}

// コンストラクタ
func NewTaskUsecase(tr repository.ITaskRepository) ITaskUsecase {
	return &taskUsecase{tr}
}

func (tu *taskUsecase) GetOwnAllTasks(userId uint) ([]model.TaskListResponse, error) {
	tasks := []model.TaskListResponse{}
	if err := tu.tr.GetOwnAllTasks(&tasks, userId); err != nil {
		return nil, err
	}
	return tasks, nil
}
