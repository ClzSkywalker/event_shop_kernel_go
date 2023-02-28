package model

import (
	"time"

	"gorm.io/gorm"
)

type TaskChildModel struct {
	BaseModel
	ParentId      int64     `json:"parent_id" gorm:"type:INTEGER;index:idx_task_pid"`
	Title         string    `json:"title" gorm:"type:varchar"`
	CompletedTime time.Time `json:"completed_time" gorm:"type:timestamp"`
	GiveUpTime    time.Time `json:"give_up_time" gorm:"type:timestamp"`
}

type ITaskChildModel interface {
	IBaseModel
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
