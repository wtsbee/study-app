package model

import "time"

type TaskList struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	Rank      uint      `json:"rank" gorm:"unique"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TaskListResponse struct {
	ID    uint           `json:"id"`
	Name  string         `json:"name"`
	Tasks []TaskResponse `json:"tasks"`
}
