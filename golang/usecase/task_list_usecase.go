package usecase

import (
	"study-app-api/model"
	"study-app-api/repository"
)

// // インターフェース
type ITaskListUsecase interface {
	UpdateTaskList(taskList model.TaskListRequest, userId uint, taskListId uint) error
}

type taskListUsecase struct {
	tlr repository.ITaskListRepository
}

// コンストラクタ
func NewTaskListUsecase(tlr repository.ITaskListRepository) ITaskListUsecase {
	return &taskListUsecase{tlr}
}

func (tlu *taskListUsecase) UpdateTaskList(taskList model.TaskListRequest, userId uint, taskListId uint) error {
	if err := tlu.tlr.UpdateTaskList(&taskList, userId, taskListId); err != nil {
		return err
	}
	return nil
}
