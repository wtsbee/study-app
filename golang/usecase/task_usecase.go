package usecase

import (
	"study-app-api/model"
	"study-app-api/repository"
)

// // インターフェース
type ITaskUsecase interface {
	GetOwnAllTasks(userId uint) ([]model.TaskResponse, error)
}

type taskUsecase struct {
	tr repository.ITaskRepository
}

// コンストラクタ
func NewTaskUsecase(tr repository.ITaskRepository) ITaskUsecase {
	return &taskUsecase{tr}
}

func (tu *taskUsecase) GetOwnAllTasks(userId uint) ([]model.TaskResponse, error) {
	tasks := []model.Task{}
	if err := tu.tr.GetOwnAllTasks(&tasks, userId); err != nil {
		return nil, err
	}
	resTasks := []model.TaskResponse{}
	for _, v := range tasks {
		t := model.TaskResponse{
			ID:        v.ID,
			Title:     v.Title,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		resTasks = append(resTasks, t)
	}
	return resTasks, nil
}
