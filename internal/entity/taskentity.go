package entity

type TaskFindByClassifyIdEntity struct {
	ClassifyId string `json:"classify_id" binding:"required"`
}

type TaskEntity struct {
	OnlyCode    string `json:"only_code"`
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
	Title       string `json:"title,omitempty" validate:"required"`
	Content     string `json:"content"`
	ClassifyId  string `json:"classify_id,omitempty" validate:"required"`
	TaskModeId  string `json:"task_mode_id,omitempty"`
	CompletedAt int64  `json:"completed_at,omitempty"`
	GiveUpAt    int64  `json:"give_up_at,omitempty"`
	StartAt     int64  `json:"start_at,omitempty"`
	EndAt       int64  `json:"end_at,omitempty"`
	ParentId    string `json:"parent_id,omitempty"`
}

type TaskUpdateReq struct {
	OnlyCode    string `json:"only_code,omitempty" binding:"required" validate:"required"`
	Title       string `json:"title,omitempty" validate:"required"`
	Content     string `json:"content"`
	TaskModeId  string `json:"task_mode_id,omitempty"`
	CompletedAt int64  `json:"completed_at,omitempty"`
	GiveUpAt    int64  `json:"give_up_at,omitempty"`
	StartAt     int64  `json:"start_at,omitempty"`
	EndAt       int64  `json:"end_at,omitempty"`
	ParentId    string `json:"parent_id,omitempty"`
}

type TaskDeleteReq struct {
	OnlyCode string `json:"only_code,omitempty" binding:"required" validate:"required"`
}
