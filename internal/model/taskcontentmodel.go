package model

import "gorm.io/gorm"

type FileUpType int

const (
	Local FileUpType = iota
	Github
)

type TaskContentModel struct {
	BaseModel
	OnlyCode string `gorm:"column:oc;type:VARCHAR(26);index:idx_task_content_oc,unique"`
	Content  string `gorm:"column:content;type:varchar"`
	// FileList []TaskFileModel `gorm:"column:file_list;type:varchar"`
}

func (r TaskContentModel) TableName() string {
	return TaskContentTableName
}

type TaskFileModel struct {
	Url    string `json:"url" gorm:"url"`
	UpType int    `json:"up_type" gorm:"up_type"`
}

type ITaskContentModel interface {
	IBaseModel
	InitData() (err error)
	FindByModel(TaskContentModel) ([]TaskContentModel, error)
	FindByOC(oc string) (result TaskContentModel, err error)
	FindByOcList(ocs string) (result []TaskContentModel, err error)
	Insert(TaskContentModel) (int64, error)
	Update(TaskContentModel) error
	Delete(string) error
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

func (m *defaultTaskContentModel) GetTx() (tx *gorm.DB) {
	return m.conn
}

func (m *defaultTaskContentModel) InitData() (err error) {
	return
}

func (m *defaultTaskContentModel) FindByModel(p TaskContentModel) (result []TaskContentModel, err error) {
	err = m.conn.Table(m.table).Where(p).Find(&result).Error
	return
}

func (m *defaultTaskContentModel) FindByOC(oc string) (result TaskContentModel, err error) {
	err = m.conn.Table(m.table).Where(TaskContentModel{OnlyCode: oc}).First(&result).Error
	return
}

func (m *defaultTaskContentModel) FindByOcList(ocs string) (result []TaskContentModel, err error) {
	return
}

func (m *defaultTaskContentModel) Insert(p TaskContentModel) (id int64, err error) {
	m.conn.Table(m.table).Create(p)
	return
}

func (m *defaultTaskContentModel) Update(p TaskContentModel) (err error) {
	err = m.conn.Table(m.table).Where(TaskContentModel{OnlyCode: p.OnlyCode}).Updates(p).Error
	return
}

func (m *defaultTaskContentModel) Delete(oc string) (err error) {
	m.conn.Table(m.table).Where(TaskContentModel{OnlyCode: oc}).Delete(TaskContentModel{})
	return
}
