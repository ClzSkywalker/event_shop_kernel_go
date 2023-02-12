package model

import (
	"gorm.io/gorm"
)

type TaskModel struct {
	BaseModel
	Title           string `json:"title" gorm:"type:TEXT" validate:"required"`
	ClassifyId      int64  `json:"classify_id" gorm:"type:INTEGER" validate:"required"`
	ContentId       int64  `json:"content_id" gorm:"type:INTEGER"`
	TaskModeId      int64  `json:"task_mode_id" gorm:"type:INTEGER"`
	ComplemetedTime int64  `json:"complemeted_time" gorm:"type:INTEGER"`
	GiveUpTime      int64  `json:"give_up_time" gorm:"type:INTEGER"`
	StartTime       int64  `json:"start_time" gorm:"type:INTEGER;index"`
	EndTime         int64  `json:"end_time" gorm:"type:INTEGER;index"`
}

type ITaskModel interface {
	IBaseModel
	SelectByModel(TaskModel) ([]TaskModel, error)
	Insert(*TaskModel) (uint, error)
	Update(*TaskModel) error
	Delete(uint) error
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
func (m *defaultTaskModel) Insert(tm *TaskModel) (id uint, err error) {
	err = m.conn.Table(m.table).Create(tm).Error
	id = tm.Id
	return
}
func (m *defaultTaskModel) Update(tm *TaskModel) (err error) {
	err = m.conn.Table(m.table).Updates(tm).Error
	return
}
func (m *defaultTaskModel) Delete(id uint) (err error) {
	err = m.conn.Table(m.table).Delete(&TaskModel{BaseModel: BaseModel{Id: id}}).Error
	return
}
