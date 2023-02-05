package model

import "time"

type TaskChildModel struct {
	BaseModel
	ParentId      int64     `json:"parent" gorm:"parent"`
	Title         string    `json:"title" gorm:"title"`
	CompletedTime time.Time `json:"completed_time" gorm:"completed_time"`
	GiveUpTime    time.Time `json:"give_up" gorm:"give_up"`
}
