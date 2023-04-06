package entity

type TaskFindByClassifyIdEntity struct {
	ClassifyId string `json:"classify_id" binding:"required"`
}

type TaskEntity struct {
	OnlyCode    string `json:"only_code,omitempty"`
	CreatedBy   string `json:"created_by,omitempty"`
	Title       string `json:"title,omitempty"`
	ClassifyId  string `json:"classify_id,omitempty"`
	ContentId   string `json:"content_id,omitempty"`
	TaskModeId  string `json:"task_mode_id,omitempty"`
	CompletedAt int64  `json:"completed_at,omitempty"`
	GiveUpAt    int64  `json:"give_up_at,omitempty"`
	StartAt     int64  `json:"start_at,omitempty"`
	EndAt       int64  `json:"end_at,omitempty"`
	ParentId    string `json:"parent_id,omitempty"`
}
