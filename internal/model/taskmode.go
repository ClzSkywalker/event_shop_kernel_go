package model

import "gorm.io/gorm"

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
	ModeId TaskModeType        `json:"mode_id" gorm:"type:INTEGER"` // 重复模式 TaskModeEnum
	Config TaskModeConfigModel `json:"config" gorm:"type:BLOB"`
}

type TaskModeConfigModel struct {
	Day []int `json:"day" gorm:"day"`
}

type ITaskModeModel interface {
	IBaseModel
	SelectByModel(TaskModeModel) ([]TaskModeModel, error)
	Insert(TaskModeModel) (int64, error)
	Update(TaskModeModel) error
	Delete(int64) error
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
func (m *defaultTaskModeModel) Insert(TaskModeModel) (id int64, err error) {
	return
}
func (m *defaultTaskModeModel) Update(TaskModeModel) (err error) {
	return
}
func (m *defaultTaskModeModel) Delete(int64) (err error) {
	return
}
