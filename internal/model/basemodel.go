package model

import "time"

type BaseModel struct {
	Id          int64     `json:"id" gorm:"id"`
	CreatedTime time.Time `json:"created_time" gorm:"created_time"`
	UpdatedTime time.Time `json:"updated_time" gorm:"updated_time"`
	DeleteTime  time.Time `json:"delete_time" gorm:"delete_time"`
}
