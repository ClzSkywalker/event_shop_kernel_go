package model

import "gorm.io/gorm"

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

type ITaskContentModel interface {
	GetTableName() string
	SelectByModel(TaskContentModel) ([]TaskContentModel, error)
	Insert(TaskContentModel) (int64, error)
	Update(TaskContentModel) error
	Delete(int64) error
}

type defaultTaskContentModel struct {
	conn  *gorm.DB
	table string
}

func NewDefaultTaskContentModel(conn *gorm.DB) *defaultTaskContentModel {
	return &defaultTaskContentModel{
		conn:  conn,
		table: "task_content",
	}
}

func (m *defaultTaskContentModel) GetTableName() string {
	return m.table
}
func (m *defaultTaskContentModel) SelectByModel(TaskContentModel) (result []TaskContentModel, err error) {
	return
}
func (m *defaultTaskContentModel) Insert(TaskContentModel) (id int64, err error) {
	return
}
func (m *defaultTaskContentModel) Update(TaskContentModel) (err error) {
	return
}
func (m *defaultTaskContentModel) Delete(int64) (err error) {
	return
}
