package model

import "time"

type TaskModel struct {
	BaseModel
	Title           string    `json:"title" gorm:"title"`
	ClassifyId      int       `json:"classify_id" gorm:"classify_id"`
	ContentId       int       `json:"content_id" gorm:"content_id"`
	TaskModeId      int       `json:"task_mode_id" gorm:"task_mode_id"`
	ComplemetedTime time.Time `json:"complemeted_time" gorm:"complemeted_time"`
	GiveUpTime      time.Time `json:"give_up_time" gorm:"give_up_time"`
	StartTime       time.Time `json:"start_time" gorm:"start_time"`
	EndTime         time.Time `json:"end_time" gorm:"end_time"`
}
