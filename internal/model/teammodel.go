package model

import (
	"github.com/clz.skywalker/event.shop/kernal/pkg/constx"
	"gorm.io/gorm"
)

type TeamModel struct {
	BaseModel
	TeamId      string `json:"team_id,omitempty" gorm:"type:VARCHAR(26);index:udx_team_tid,unique"`
	CreatedBy   string `json:"created_by,omitempty" gorm:"type:VARCHAR(26);index:idx_team_uid"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type ITeamModel interface {
	IBaseModel
	InitData(lang, tid, uid string) (err error)
}

type defaultTeamModel struct {
	conn  *gorm.DB
	table string
}

func NewDefaultTeamModel(conn *gorm.DB) ITeamModel {
	return &defaultTeamModel{
		conn:  conn,
		table: TeamTableName,
	}
}

func (m *defaultTeamModel) TableName() string {
	return m.table
}

func (m *defaultTeamModel) CreateTable() (err error) {
	err = m.conn.Table(m.table).AutoMigrate(&TeamModel{})
	return
}

func (m *defaultTeamModel) DropTable() (err error) {
	err = m.conn.Table(m.table).Migrator().DropTable(m.table)
	return
}

func (m *defaultTeamModel) InitData(lang, uid, tid string) (err error) {
	name := ""
	switch name {
	case constx.LangChinese:
		name = "新的起点"
	default:
		name = "new start"
	}
	tm := &TeamModel{TeamId: tid, CreatedBy: uid, Name: name}
	_, err = m.Insert(tm)
	return
}

func (m *defaultTeamModel) Insert(tm *TeamModel) (id uint, err error) {
	err = m.conn.Table(m.table).Create(tm).Error
	id = tm.Id
	return
}
