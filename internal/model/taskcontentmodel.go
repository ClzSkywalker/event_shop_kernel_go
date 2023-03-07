package model

import "gorm.io/gorm"

type FileUpType int

const (
	Local FileUpType = iota
	Github
)

type TaskContentModel struct {
	BaseModel
	TaskId   string          `json:"task_id" gorm:"type:VARCHAR(26);index:idx_task_content_tid,unique"`
	Content  string          `json:"content" gorm:"type:varchar"`
	FileList []TaskFileModel `json:"file_list" gorm:"type:varchar"`
}

type TaskFileModel struct {
	Url    string `json:"url" gorm:"url"`
	UpType int    `json:"up_type" gorm:"up_type"`
}

type ITaskContentModel interface {
	IBaseModel
	InitData() (err error)
	SelectByModel(TaskContentModel) ([]TaskContentModel, error)
	Insert(TaskContentModel) (int64, error)
	Update(TaskContentModel) error
	Delete(int64) error
}

type defaultTaskContentModel struct {
	conn  *gorm.DB
	table string
}

func NewDefaultTaskContentModel(conn *gorm.DB) ITaskContentModel {
	return &defaultTaskContentModel{
		conn:  conn,
		table: TaskContentTableName,
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
	err = m.conn.Table(m.table).Migrator().DropTable(m.table)
	return
}

func (m *defaultTaskContentModel) InitData() (err error) {
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
