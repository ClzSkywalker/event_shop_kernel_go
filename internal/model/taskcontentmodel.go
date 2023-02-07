package model

import "gorm.io/gorm"

type FileUpType int

const (
	Local FileUpType = iota
	Github
)

type TaskContentModel struct {
	BaseModel
	TaskId   int             `json:"task_id" gorm:"type:INTEGER"`
	Content  string          `json:"content" gorm:"type:TEXT"`
	FileList []TaskFileModel `json:"file_list" gorm:"type:BLOB"`
}

type TaskFileModel struct {
	Url    string `json:"url" gorm:"url"`
	UpType int    `json:"up_type" gorm:"up_type"`
}

type ITaskContentModel interface {
	IBaseModel
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

func (m *defaultTaskContentModel) TableName() string {
	return m.table
}

func (m *defaultTaskContentModel) CreateTable() (err error) {
	err = m.conn.Table(m.table).AutoMigrate(&TaskContentModel{})
	return
}

func (m *defaultTaskContentModel) DropTable() (err error) {
	err = m.conn.Table(m.table).Migrator().DropTable(m)
	return
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
