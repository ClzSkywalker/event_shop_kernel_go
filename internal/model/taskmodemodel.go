package model

import (
	"github.com/clz.skywalker/event.shop/kernal/pkg/utils"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type TaskModeType int

const (
	TaskModeNormal TaskModeType = iota
	TaskModeDay
	TaskModeWeek
	TaskModeMonth
	TaskModeYear
	TaskModeWorkday          // 工作日(周一-周五)
	TaskModeLegalWorkingDay  // 法定工作日
	TaskModeStatutoryHoliday // 法定节假日
	TaskModeCustome          // 自定义
)

type TaskModeModel struct {
	BaseModel
	OnlyCode string         `gorm:"column:oc;type:VARCHAR(26);index:udx_task_mode_oc,unique"`
	ModeType TaskModeType   `gorm:"column:mode_type;type:INTEGER"` // 重复模式 TaskModeEnum
	TeamId   string         `gorm:"column:team_id;type:VARCHAR(26);index:idx_task_mode_tid"`
	Config   datatypes.JSON `gorm:"column:config;type:varchar"`
}

type TaskModeConfigModel struct {
	Days []int `json:"days"`
}

type ITaskModeModel interface {
	IBaseModel
	InitData(tmid, tid string) (err error)
	SelectByModel(TaskModeModel) ([]TaskModeModel, error)
	Insert(*TaskModeModel) (uint, error)
	InsertAll(tmList []*TaskModeModel) (err error)
	Update(*TaskModeModel) error
	Delete(uint) error
}

type defaultTaskModeModel struct {
	conn  *gorm.DB
	table string
}

func NewDefaultTaskModeModel(conn *gorm.DB) ITaskModeModel {
	return &defaultTaskModeModel{
		conn:  conn,
		table: TaskModeTableName,
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
	err = m.conn.Table(m.table).Migrator().DropTable(m.table)
	return
}

func (m *defaultTaskModeModel) InitData(tmid, tid string) (err error) {
	tm1 := &TaskModeModel{OnlyCode: tmid, TeamId: tid, ModeType: TaskModeNormal}
	tm2 := &TaskModeModel{OnlyCode: utils.NewUlid(), TeamId: tid, ModeType: TaskModeDay}
	tm3 := &TaskModeModel{OnlyCode: utils.NewUlid(), TeamId: tid, ModeType: TaskModeWeek}
	tm4 := &TaskModeModel{OnlyCode: utils.NewUlid(), TeamId: tid, ModeType: TaskModeMonth}
	tm5 := &TaskModeModel{OnlyCode: utils.NewUlid(), TeamId: tid, ModeType: TaskModeYear}
	tm6 := &TaskModeModel{OnlyCode: utils.NewUlid(), TeamId: tid, ModeType: TaskModeWorkday}
	tm7 := &TaskModeModel{OnlyCode: utils.NewUlid(), TeamId: tid, ModeType: TaskModeLegalWorkingDay}
	err = m.InsertAll([]*TaskModeModel{tm1, tm2, tm3, tm4, tm5, tm6, tm7})
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

func (m *defaultTaskModeModel) InsertAll(tmList []*TaskModeModel) (err error) {
	err = m.conn.Table(m.table).Create(tmList).Error
	return
}

func (m *defaultTaskModeModel) Update(tmm *TaskModeModel) (err error) {
	return
}

func (m *defaultTaskModeModel) Delete(id uint) (err error) {
	err = m.conn.Table(m.table).Delete(&TaskModeModel{BaseModel: BaseModel{Id: id}}).Error
	return
}
