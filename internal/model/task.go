package model

import (
	"time"

	"gorm.io/gorm"
)

type TaskModel struct {
	BaseModel
	Title           string    `json:"title" gorm:"type:TEXT"`
	ClassifyId      int64     `json:"classify_id" gorm:"type:INTEGER"`
	ContentId       int64     `json:"content_id" gorm:"type:INTEGER"`
	TaskModeId      int64     `json:"task_mode_id" gorm:"type:INTEGER"`
	ComplemetedTime time.Time `json:"complemeted_time" gorm:"type:timestamp"`
	GiveUpTime      time.Time `json:"give_up_time" gorm:"type:timestamp"`
	StartTime       time.Time `json:"start_time" gorm:"type:timestamp;index"`
	EndTime         time.Time `json:"end_time" gorm:"type:timestamp;index"`
}

type ITaskModel interface {
	IBaseModel
	SelectByModel(TaskModel) ([]TaskModel, error)
	Insert(TaskModel) (int64, error)
	Update(TaskModel) error
	Delete(int64) error
}

type defaultTaskModel struct {
	conn  *gorm.DB
	table string
}

func NewDefaultTaskModel(conn *gorm.DB) *defaultTaskModel {
	return &defaultTaskModel{
		conn:  conn,
		table: "task",
	}
}

func (m *defaultTaskModel) TableName() string {
	return m.table
}

func (m *defaultTaskModel) CreateTable() (err error) {
	err = m.conn.Table(m.table).AutoMigrate(&TaskModel{})
	return
}

func (m *defaultTaskModel) DropTable() (err error) {
	err = m.conn.Table(m.table).Migrator().DropTable(m)
	return
}

func (m *defaultTaskModel) SelectByModel(TaskModel) (result []TaskModel, err error) {
	return
}
func (m *defaultTaskModel) Insert(TaskModel) (id int64, err error) {
	return
}
func (m *defaultTaskModel) Update(TaskModel) (err error) {
	return
}
func (m *defaultTaskModel) Delete(int64) (err error) {
	return
}
