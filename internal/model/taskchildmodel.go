package model

import (
	"time"

	"gorm.io/gorm"
)

type TaskChildModel struct {
	BaseModel
	ParentId      int64     `json:"parent" gorm:"parent"`
	Title         string    `json:"title" gorm:"title"`
	CompletedTime time.Time `json:"completed_time" gorm:"completed_time"`
	GiveUpTime    time.Time `json:"give_up" gorm:"give_up"`
}

type ITaskChildModel interface {
	GetTableName() string
	SelectByModel(TaskChildModel) ([]TaskChildModel, error)
	Insert(TaskChildModel) (int64, error)
	Update(TaskChildModel) error
	Delete(int64) error
}

type defaultTaskChildModel struct {
	conn  *gorm.DB
	table string
}

func NewDefaultTaskChildModel(conn *gorm.DB) *defaultTaskChildModel {
	return &defaultTaskChildModel{
		conn:  conn,
		table: "task_child",
	}
}

func (m *defaultTaskChildModel) GetTableName() string {
	return m.table
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
