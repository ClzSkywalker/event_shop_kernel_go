package model

import (
	"time"

	"gorm.io/gorm"
)

type TaskModel struct {
	BaseModel
	Title           string    `json:"title" gorm:"title"`
	ClassifyId      int64     `json:"classify_id" gorm:"classify_id"`
	ContentId       int64     `json:"content_id" gorm:"content_id"`
	TaskModeId      int64     `json:"task_mode_id" gorm:"task_mode_id"`
	ComplemetedTime time.Time `json:"complemeted_time" gorm:"complemeted_time"`
	GiveUpTime      time.Time `json:"give_up_time" gorm:"give_up_time"`
	StartTime       time.Time `json:"start_time" gorm:"start_time"`
	EndTime         time.Time `json:"end_time" gorm:"end_time"`
}

type ITaskModel interface {
	GetTableName() string
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

func (m *defaultTaskModel) GetTableName() string {
	return m.table
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
