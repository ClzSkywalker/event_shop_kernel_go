package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type TaskModeType int

const (
	Normal TaskModeType = iota
	Day
	Week
	Month
	Year
	Workday          // 工作日(周一-周五)
	LegalWorkingDay  // 法定工作日
	StatutoryHoliday // 法定节假日
)

type TaskModeModel struct {
	BaseModel
	ModeId int            `json:"mode_id" gorm:"type:INTEGER" validate:"required"` // 重复模式 TaskModeEnum
	Config datatypes.JSON `json:"config" gorm:"type:JSON"`
}

type TaskModeConfigModel struct {
	Day []int `json:"day" gorm:"day"`
}

type ITaskModeModel interface {
	IBaseModel
	SelectByModel(TaskModeModel) ([]TaskModeModel, error)
	Insert(*TaskModeModel) (uint, error)
	Update(*TaskModeModel) error
	Delete(uint) error
}

type defaultTaskModeModel struct {
	conn  *gorm.DB
	table string
}

func NewDefaultTaskModeModel(conn *gorm.DB) *defaultTaskModeModel {
	return &defaultTaskModeModel{
		conn:  conn,
		table: "task_mode",
	}
}

func (m *defaultTaskModeModel) TableName() string {
	return m.table
}

func (m *defaultTaskModeModel) CreateTable() (err error) {
	err = m.conn.Table(m.table).AutoMigrate(&TaskModeModel{})
	return
}

func (m *defaultTaskModeModel) DropTable() (err error) {
	err = m.conn.Table(m.table).Migrator().DropTable(m)
	return
}
func (m *defaultTaskModeModel) SelectByModel(TaskModeModel) (result []TaskModeModel, err error) {
	return
}
func (m *defaultTaskModeModel) Insert(tmm *TaskModeModel) (id uint, err error) {
	err = m.conn.Table(m.table).Create(tmm).Error
	id = tmm.Id
	return
}
func (m *defaultTaskModeModel) Update(tmm *TaskModeModel) (err error) {
	return
}
func (m *defaultTaskModeModel) Delete(id uint) (err error) {
	err = m.conn.Table(m.table).Delete(&TaskModeModel{BaseModel: BaseModel{Id: id}}).Error
	return
}
