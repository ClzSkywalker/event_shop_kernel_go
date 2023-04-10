package model

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

const (
	UserTableName        = "user"
	TeamTableName        = "team"
	UserToTeamTableName  = "user_to_team"
	ClassifyTableName    = "classify"
	DevideTableName      = "devide"
	TaskTableName        = "task"
	TaskModeTableName    = "task_mode"
	TaskContentTableName = "task_content"
)

const (
	// 父子文件夹递归查询
	recursiveSql = `WITH RECURSIVE all_folders AS (
		SELECT t1.* 
		FROM %s t1
		WHERE t1.parent_id = '' %s
		UNION ALL
		SELECT t2.*
		FROM %s t1
		JOIN %s t2 ON t2.parent_id = t1.oc where 1=1 %s
	  )`
)

type BaseModel struct {
	Id        uint                  `gorm:"primarykey"`
	CreatedAt int64                 `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt int64                 `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt soft_delete.DeletedAt `json:"deleted_at" gorm:"softDelete"`
}

type IBaseModel interface {
	TableName() string
	CreateTable() (err error) // 创建表
	DropTable() (err error)   // 删除表
	GetTx() (tx *gorm.DB)
}

func recursive(table, where string) string {
	return fmt.Sprintf(recursiveSql, table, where, table, table, where)
}
