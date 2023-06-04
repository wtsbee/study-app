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
	ttl := []model.TaskAndTaskListResponse{}
	if err := tu.tr.GetOwnAllTasks(&ttl, userId); err != nil {
		return nil, err
	}

	tasks := []model.TaskListResponse{}
	arr := []model.TaskResponse{}
	for i, v := range ttl {
		taskRes := model.TaskResponse{
			ID:    v.TaskId,
			Title: v.TaskTitle,
		}
		arr = append(arr, taskRes)
		if i+1 == len(ttl) || ttl[i].TaskListRank != ttl[i+1].TaskListRank {
			taskListRes := model.TaskListResponse{
				ID:    v.TaskListId,
				Name:  v.TaskListName,
				Tasks: arr,
			}
			tasks = append(tasks, taskListRes)
			arr = nil
		}
	}
	return tasks, nil
}
