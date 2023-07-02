package usecase

import (
	"study-app-api/model"
	"study-app-api/repository"
)

// // インターフェース
type ITaskDetailUsecase interface {
	GetTaskDetail(taskID uint, userId uint) (model.TaskDetail, error)
	CreateTaskDetail(taskDetailId uint, userId uint) error
	UpdateTaskDetail(taskDetail model.TaskDetailRequest, userId uint) error
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

func (tdu *taskDetailUsecase) CreateTaskDetail(taskDetailId uint, userId uint) error {
	if err := tdu.tdr.CreateTaskDetail(taskDetailId, userId); err != nil {
		return err
	}
	return nil
}

func (tdu *taskDetailUsecase) UpdateTaskDetail(taskDetail model.TaskDetailRequest, userId uint) error {
	if err := tdu.tdr.UpdateTaskDetail(&taskDetail, userId); err != nil {
		return err
	}
	return nil
}
