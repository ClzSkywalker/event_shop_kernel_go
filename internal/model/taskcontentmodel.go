package model

type FileUpType int

const (
	Local FileUpType = iota
	Github
)

type TaskContentModel struct {
	BaseModel
	TaskId       int    `json:"task_id" gorm:"task_id"`
	Content      string `json:"content" gorm:"content"`
	FileListByte []byte `json:"file_list" gorm:"file_list"`
	FileList     []TaskFileModel
}

type TaskFileModel struct {
	Url    string `json:"url" gorm:"url"`
	UpType int    `json:"up_type" gorm:"up_type"`
}
