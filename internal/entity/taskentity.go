package entity

import "github.com/clz.skywalker/event.shop/kernal/pkg/constx"

type TaskFindByClassifyIdEntity struct {
	ClassifyId string                  `json:"classify_id,omitempty" binding:"required" validate:"required"`
	TaskStatus constx.TaskCompleteType `json:"task_status,omitempty"`  // 任务完成
	StartAt    int64                   `json:"start_at,omitempty"`     // 开始时间
	EndAt      int64                   `json:"end_at,omitempty"`       // 结束时间
	TaskModeId string                  `json:"task_mode_id,omitempty"` // 任务模式
	TaskOrder  constx.TaskOrderType    `json:"task_order,omitempty"`   // 排序模式
}

type TaskEntity struct {
	OnlyCode    string `json:"oc" gorm:"column:oc;"`
	CreatedBy   string `json:"created_by"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	ClassifyId  string `json:"classify_id"`
	TaskModeId  string `json:"task_mode_id"`
	CompletedAt int64  `json:"completed_at"`
	GiveUpAt    int64  `json:"give_up_at"`
	StartAt     int64  `json:"start_at"`
	EndAt       int64  `json:"end_at"`
	ParentId    string `json:"parent_id"`
}

type TaskInsertReq struct {
	Title      string `json:"title,omitempty" validate:"required"`
	Content    string `json:"content"`
	DevideId   string `json:"devide_id,omitempty" validate:"required"`
	TaskModeId string `json:"task_mode_id,omitempty"`
	StartAt    int64  `json:"start_at,omitempty"`
	EndAt      int64  `json:"end_at,omitempty"`
	ParentId   string `json:"parent_id,omitempty"`
}

type TaskUpdateReq struct {
	OnlyCode    string `json:"oc,omitempty" binding:"required" validate:"required"`
	Title       string `json:"title,omitempty" validate:"required"`
	Content     string `json:"content"`
	TaskModeId  string `json:"task_mode_id,omitempty"`
	DevideId    string `json:"devide_id,omitempty" validate:"required"`
	CompletedAt int64  `json:"completed_at,omitempty"`
	GiveUpAt    int64  `json:"give_up_at,omitempty"`
	StartAt     int64  `json:"start_at,omitempty"`
	EndAt       int64  `json:"end_at,omitempty"`
	ParentId    string `json:"parent_id,omitempty"`
}

type TaskDeleteReq struct {
	OnlyCode string `json:"oc,omitempty" binding:"required" validate:"required"`
}
