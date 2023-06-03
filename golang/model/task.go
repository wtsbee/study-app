package model

import "time"

type Task struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	Title      string    `json:"title" gorm:"not null"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	User       User      `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	UserId     uint      `json:"user_id" gorm:"not null"`
	TaskList   TaskList  `json:"task_list" gorm:"foreignKey:TaskListId; constraint:OnDelete:CASCADE"`
	TaskListId uint      `json:"task_list_id" gorm:"not null"`
	Rank       uint      `json:"rank" gorm:"unique"`
}

type TaskResponse struct {
	Rank  uint   `json:"rank"`
	Title string `json:"title"`
}
