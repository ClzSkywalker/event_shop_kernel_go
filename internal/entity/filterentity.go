package entity

import "github.com/clz.skywalker/event.shop/kernal/pkg/constx"

type TaskFilterParam struct {
	ClassifyId []string                `json:"classify_id,omitempty"`
	DevideIds  []string                `json:"devide_ids,omitempty"`
	Keyword    string                  `json:"keyword,omitempty"`
	BeginAt    int64                   `json:"begin_at,omitempty"`    // 时间范围-开始
	CloseAt    int64                   `json:"close_at,omitempty"`    //
	OrderType  constx.TaskOrderType    `json:"order_type,omitempty"`  // 排序规则
	TaskStatus constx.TaskCompleteType `json:"task_status,omitempty"` // 任务状态
}
