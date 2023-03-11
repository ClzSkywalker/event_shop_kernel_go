package model

import "gorm.io/gorm"

type UserToTeamModel struct {
	BaseModel
	Uid  string `gorm:"column:uid;type:VARCHAR(26);index:udx_utt_uid_tid_add,priority:1,unique"`
	Tid  string `gorm:"column:tid;type:VARCHAR(26);index:udx_utt_uid_tid_add,priority:2,unique"`
	Sort int    `gorm:"column:sort;type:INTEGER"`
}

type IUserToTeamModel interface {
	IBaseModel
	InitData(uid, tid string) (err error)
	Insert(p *UserToTeamModel) (id uint, err error)
}

type defaultUserToTeamModel struct {
	conn  *gorm.DB
	table string
}

func NewDefaultUserToTeamModel(conn *gorm.DB) IUserToTeamModel {
	return &defaultUserToTeamModel{
		conn:  conn,
		table: UserToTeamTableName,
	}
}

func (m *defaultUserToTeamModel) TableName() string {
	return m.table
}

func (m *defaultUserToTeamModel) CreateTable() (err error) {
	err = m.conn.Table(m.table).AutoMigrate(&UserToTeamModel{})
	return
}

func (m *defaultUserToTeamModel) DropTable() (err error) {
	err = m.conn.Table(m.table).Migrator().DropTable(m.table)
	return
}

func (m *defaultUserToTeamModel) GetTx() (tx *gorm.DB) {
	return m.conn
}

func (m *defaultUserToTeamModel) InitData(uid, tid string) (err error) {
	um := &UserToTeamModel{Uid: uid, Tid: tid}
	_, err = m.Insert(um)
	return
}

func (m *defaultUserToTeamModel) Insert(p *UserToTeamModel) (id uint, err error) {
	err = m.conn.Table(m.table).Create(p).Error
	if err != nil {
		return
	}
	id = p.Id
	return
}

func (m *defaultUserToTeamModel) Create(utt *UserToTeamModel) (err error) {
	err = m.conn.Table(m.table).Create(utt).Error
	return
}

func (m *defaultUserToTeamModel) Delete(utt *UserToTeamModel) (err error) {
	err = m.conn.Table(m.table).Where(utt).Delete(utt).Error
	return
}
