package model

import (
	"time"

	"gorm.io/gorm"
)

type TaskList struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"not null"`
	Rank      uint           `json:"rank"`
	User      User           `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	UserId    uint           `json:"user_id" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type TaskListResponse struct {
	ID    uint           `json:"id"`
	Name  string         `json:"name"`
	Tasks []TaskResponse `json:"tasks"`
}
