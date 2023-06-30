package model

import (
	"time"

	"gorm.io/gorm"
)

type TaskDetail struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Detail    string         `json:"detail"`
	User      User           `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	UserId    uint           `json:"user_id" gorm:"not null"`
	Task      Task           `json:"task" gorm:"foreignKey:TaskId; constraint:OnDelete:CASCADE"`
	TaskId    uint           `json:"task_id" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type TaskDetailRequest struct {
	ID     uint   `json:"id"`
	Detail string `json:"detail"`
	TaskId uint   `json:"task_id"`
}

type TaskDetailResponse struct {
	ID     uint   `json:"id"`
	Detail string `json:"detail"`
	TaskId uint   `json:"task_id"`
}
