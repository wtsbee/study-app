package model

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	Title      string         `json:"title" gorm:"not null"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	User       User           `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	UserId     uint           `json:"user_id" gorm:"not null"`
	TaskList   TaskList       `json:"task_list" gorm:"foreignKey:TaskListId; constraint:OnDelete:CASCADE"`
	TaskListId uint           `json:"task_list_id" gorm:"not null"`
	Rank       uint           `json:"rank"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
}

type TaskResponse struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}

type TaskAndTaskListResponse struct {
	TaskListId   uint
	TaskListRank uint
	TaskListName string
	TaskId       uint
	TaskTitle    string
	UserId       uint
}
