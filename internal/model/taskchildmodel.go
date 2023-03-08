package model

import (
	"gorm.io/gorm"
)

type TaskChildModel struct {
	BaseModel
	OnlyCode    string `gorm:"column:oc;type:VARCHAR(26);index:udx_task_child_oc,unique"`
	ParentId    string `gorm:"column:parent_id;type:VARCHAR(26);index:idx_task_pid"`
	Title       string `gorm:"column:title;type:varchar"`
	CompletedAt int64  `gorm:"column:completed_at;type:INTEGER"`
	GiveUpAt    int64  `gorm:"column:give_up_at;type:INTEGER"`
}

type ITaskChildModel interface {
	IBaseModel
	InitData() (err error)
	SelectByModel(TaskChildModel) ([]TaskChildModel, error)
	Insert(TaskChildModel) (int64, error)
	Update(TaskChildModel) error
	Delete(int64) error
}

type defaultTaskChildModel struct {
	conn  *gorm.DB
	table string
}

func NewDefaultTaskChildModel(conn *gorm.DB) ITaskChildModel {
	return &defaultTaskChildModel{
		conn:  conn,
		table: TaskChildTableName,
	}
}

func (m *defaultTaskChildModel) TableName() string {
	return m.table
}

func (m *defaultTaskChildModel) CreateTable() (err error) {
	err = m.conn.Table(m.table).AutoMigrate(&TaskChildModel{})
	return
}

func (m *defaultTaskChildModel) DropTable() (err error) {
	err = m.conn.Table(m.table).Migrator().DropTable(m.table)
	return
}

func (m *defaultTaskChildModel) InitData() (err error) {
	return
}

func (m *defaultTaskChildModel) SelectByModel(TaskChildModel) (result []TaskChildModel, err error) {
	return
}

func (m *defaultTaskChildModel) Insert(TaskChildModel) (id int64, err error) {
	return
}

func (m *defaultTaskChildModel) Update(TaskChildModel) (err error) {
	return
}

func (m *defaultTaskChildModel) Delete(int64) (err error) {
	return
}
