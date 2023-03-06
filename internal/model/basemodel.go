package model

import (
	"gorm.io/plugin/soft_delete"
)

const (
	UserTableName        = "user"
	TeamTableName        = "team"
	UserToTeamTableName  = "user_to_team"
	ClassifyTableName    = "classify"
	TaskTableName        = "task"
	TaskModeTableName    = "task_mode"
	TaskChildTableName   = "task_child"
	TaskContentTableName = "task_content"
)

type BaseModel struct {
	Id        uint                  `gorm:"primarykey"`
	CreatedAt int64                 `json:"created_at" gorm:"autoUpdateTime"`
	UpdatedAt int64                 `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt soft_delete.DeletedAt `json:"deleted_at" gorm:"softDelete"`
}

type IBaseModel interface {
	TableName() string
	CreateTable() (err error) // 创建表
	DropTable() (err error)   // 删除表
}
