package usecase

import (
	"study-app-api/model"
	"study-app-api/repository"
)

// // インターフェース
type ITaskDetailUsecase interface {
	GetTaskDetail(taskID uint, userId uint) (model.TaskDetail, error)
}

type taskDetailUsecase struct {
	tdr repository.ITaskDetailRepository
}

// コンストラクタ
func NewTaskDetailUsecase(tdr repository.ITaskDetailRepository) ITaskDetailUsecase {
	return &taskDetailUsecase{tdr}
}

func (tdu *taskDetailUsecase) GetTaskDetail(taskID uint, userId uint) (model.TaskDetail, error) {
	taskDetail := model.TaskDetail{}
	if err := tdu.tdr.GetTaskDetail(&taskDetail, taskID, userId); err != nil {
		return model.TaskDetail{}, err
	}
	return taskDetail, nil
}
