package entity

import (
	"github.com/clz.skywalker/event.shop/kernal/pkg/constx"
)

type TaskModeEntity struct {
	OnlyCode string               `gorm:"column:oc;type:VARCHAR(26);index:udx_task_mode_oc,unique"`
	ModeType constx.TaskModeType  `gorm:"column:mode_type;type:INTEGER"` // 重复模式 TaskModeEnum
	TeamId   string               `gorm:"column:team_id;type:VARCHAR(26);index:idx_task_mode_tid"`
	Config   TaskModeConfigEntity `gorm:"column:config;type:varchar"`
}

type TaskModeConfigEntity struct {
	Days []int `json:"days"`
}
