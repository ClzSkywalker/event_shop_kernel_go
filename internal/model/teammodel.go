package model

import (
	"fmt"

	"github.com/clz.skywalker/event.shop/kernal/pkg/constx"
	"gorm.io/gorm"
)

type TeamModel struct {
	BaseModel
	OnlyCode    string `json:"oc" gorm:"column:oc;type:VARCHAR(26);index:udx_team_tid,unique"`
	CreatedBy   string `gorm:"column:created_by;type:VARCHAR(26);index:idx_team_uid"`
	Name        string `gorm:"column:name;type:VARCHAR"`
	Description string `gorm:"column:description;type:VARCHAR"`
}

func (m TeamModel) TableName() string {
	return TeamTableName
}

type ITeamModel interface {
	IBaseModel
	InitData(lang, tid, uid string) (err error)
	Find(p TeamModel) (result []TeamModel, err error)
	FindMyTeam(uid string) (result []TeamModel, err error)
	First(p TeamModel) (result TeamModel, err error)
	Create(p *TeamModel) (id uint, err error)
	Update(p TeamModel) (err error)
	Delete(tid string) (err error)
	FindMyTeamName(uid, tname string) (result TeamModel, err error)
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

func (m *defaultTeamModel) GetTx() (tx *gorm.DB) {
	return m.conn
}

func (m *defaultTeamModel) InitData(lang, uid, tid string) (err error) {
	name := ""
	switch name {
	case constx.LangChinese:
		name = "新的起点"
	default:
		name = "new start"
	}
	tm := &TeamModel{OnlyCode: tid, CreatedBy: uid, Name: name}
	_, err = m.Create(tm)
	return
}

func (m *defaultTeamModel) Find(p TeamModel) (result []TeamModel, err error) {
	err = m.conn.Table(m.table).Where(p).Find(&result).Error
	return
}

func (m *defaultTeamModel) FindMyTeam(uid string) (result []TeamModel, err error) {
	err = m.conn.Raw(fmt.Sprintf(`select
	t.*,
	utt.sort
from
	%s t
join %s utt on
	t.created_by = utt.uid
order by
	utt.sort asc`, TeamTableName, UserToTeamTableName)).Scan(&result).Error
	return
}

func (m *defaultTeamModel) FindMyTeamName(uid, tname string) (result TeamModel, err error) {
	err = m.conn.Raw(fmt.Sprintf(`select
	t.*,
	utt.sort
from
	%s t
join %s utt on
	t.created_by = utt.uid
where t.name='%s';
`, TeamTableName, UserToTeamTableName, tname)).Scan(&result).Error
	return
}

func (m *defaultTeamModel) First(p TeamModel) (result TeamModel, err error) {
	err = m.conn.Table(m.table).Where(p).First(&result).Error
	return
}

func (m *defaultTeamModel) Create(tm *TeamModel) (id uint, err error) {
	err = m.conn.Table(m.table).Create(tm).Error
	id = tm.Id
	return
}

func (m *defaultTeamModel) Update(p TeamModel) (err error) {
	err = m.conn.Table(m.table).Where(TeamModel{OnlyCode: p.OnlyCode}).Updates(p).Error
	return
}

func (m *defaultTeamModel) Delete(tid string) (err error) {
	err = m.conn.Table(m.table).Where(TeamModel{OnlyCode: tid}).Delete(TeamModel{}).Error
	return
}
