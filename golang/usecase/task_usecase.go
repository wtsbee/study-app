package usecase

import (
	"study-app-api/model"
	"study-app-api/repository"
)

// // インターフェース
type ITaskUsecase interface {
	GetOwnAllTasks(userId uint) ([]model.TaskListResponse, error)
	CreateTask(task model.TaskRequest, userId uint) error
	UpdateTask(task model.TaskRequest, userId uint, taskId uint) error
	UpdateOwnAllTasks(taskList []model.TaskListResponse, userId uint) error
	DeleteTaskList(taskListId uint, userId uint) error
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
		taskRes := model.TaskResponse{}
		if v.TaskId != 0 {
			taskRes = model.TaskResponse{
				ID:    v.TaskId,
				Title: v.TaskTitle,
			}
			arr = append(arr, taskRes)
		}
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

func (tu *taskUsecase) CreateTask(task model.TaskRequest, userId uint) error {
	if err := tu.tr.CreateTask(&task, userId); err != nil {
		return err
	}
	return nil
}

func (tu *taskUsecase) UpdateTask(task model.TaskRequest, userId uint, taskId uint) error {
	if err := tu.tr.UpdateTask(&task, userId, taskId); err != nil {
		return err
	}
	return nil
}

func (tu *taskUsecase) UpdateOwnAllTasks(taskList []model.TaskListResponse, userId uint) error {
	if err := tu.tr.UpdateOwnAllTasks(&taskList, userId); err != nil {
		return err
	}
	return nil
}

func (tu *taskUsecase) DeleteTaskList(taskListId uint, userId uint) error {
	if err := tu.tr.DeleteTaskList(taskListId, userId); err != nil {
		return err
	}
	return nil
}
