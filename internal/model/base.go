package model

import "time"

type BaseModel struct {
	Id          int64     `json:"id" gorm:"primaryKey"`
	CreatedTime time.Time `json:"created_time" gorm:"type:timestamp"`
	UpdatedTime time.Time `json:"updated_time" gorm:"type:timestamp"`
	DeleteTime  time.Time `json:"delete_time" gorm:"type:timestamp;index"`
}

type IBaseModel interface {
	TableName() string
	CreateTable() (err error) // 创建表
	DropTable() (err error)   // 删除表
}
