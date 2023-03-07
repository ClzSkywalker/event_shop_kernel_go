package model

import (
	"github.com/clz.skywalker/event.shop/kernal/pkg/constx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/utils"
	"gorm.io/gorm"
)

type TaskModel struct {
	BaseModel
	OnlyCode        string `json:"only_code" gorm:"type:VARCHAR(26);index:udx_task_oc,unique"`
	CreatedBy       string `json:"created_by" gorm:"type:VARCHAR(26);index:idx_task_uid"`
	Title           string `json:"title" gorm:"type:VARCHAR(255)"`
	ClassifyId      string `json:"classify_id" gorm:"type:VARCHAR(26);index:idx_task_cid"`
	ContentId       string `json:"content_id" gorm:"type:VARCHAR(26)"`
	TaskModeId      string `json:"task_mode_id" gorm:"type:VARCHAR(26)"`
	ComplemetedTime int64  `json:"complemeted_time" gorm:"type:INTEGER"`
	GiveUpTime      int64  `json:"give_up_time" gorm:"type:INTEGER"`
	StartTime       int64  `json:"start_time" gorm:"type:INTEGER;index"`
	EndTime         int64  `json:"end_time" gorm:"type:INTEGER;index"`
}

type ITaskModel interface {
	IBaseModel
	InitData(lang, uid, cid, tid string) (err error)
	SelectByModel(TaskModel) ([]TaskModel, error)
	Insert(*TaskModel) (uint, error)
	InsertAll(tm []*TaskModel) (err error)
	Update(*TaskModel) error
	Delete(uint) error
}

type defaultTaskModel struct {
	conn  *gorm.DB
	table string
}

func NewDefaultTaskModel(conn *gorm.DB) ITaskModel {
	return &defaultTaskModel{
		conn:  conn,
		table: TaskTableName,
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
	err = m.conn.Table(m.table).Migrator().DropTable(m.table)
	return
}

func (m *defaultTaskModel) InitData(lang, uid, cid, tid string) (err error) {
	t1 := &TaskModel{OnlyCode: utils.NewUlid(), CreatedBy: uid, ClassifyId: cid, TaskModeId: tid}
	t2 := &TaskModel{OnlyCode: utils.NewUlid(), CreatedBy: uid, ClassifyId: cid, TaskModeId: tid}
	switch lang {
	case constx.LangChinese:
		t1.Title = "第一步"
		t2.Title = "测试"
	default:
		t1.Title = "The first step"
		t2.Title = "test"
	}
	err = m.InsertAll([]*TaskModel{t1})
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

func (m *defaultTaskModel) InsertAll(tm []*TaskModel) (err error) {
	err = m.conn.Table(m.table).Create(tm).Error
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
