package model

import (
	"github.com/clz.skywalker/event.shop/kernal/pkg/constx"
	"gorm.io/gorm"
)

//
// Author         : ClzSkywalker
// Date           : 2023-04-10
// Description    : task 分组
// return          {*}
//
type DevideModel struct {
	BaseModel
	OnlyCode   string `json:"oc" gorm:"column:oc;type:VARCHAR(26);index:udx_devide_oc,unique"`
	ClassifyId string `gorm:"column:classify_id;type:VARCHAR(26);index:idx_devide_cid"`
	CreatedBy  string `gorm:"column:created_by;type:VARCHAR(26);index:idx_devide_uid"`
	Title      string `gorm:"column:title;type:VARCHAR(255)"`
	Sort       int    `gorm:"column:sort;type:INTEGER"`
}

func (r DevideModel) TableName() string {
	return DevideTableName
}

type IDevideModel interface {
	IBaseModel
	InitData(lang, uid, cid, did string) (err error)
	Insert(tm *DevideModel) (id uint, err error)
	InsertAll(tm []*DevideModel) (err error)
}

type defaultDevideModel struct {
	conn  *gorm.DB
	table string
}

func NewDefaultDevideModel(conn *gorm.DB) IDevideModel {
	return &defaultDevideModel{
		conn:  conn,
		table: DevideTableName,
	}
}
func (r *defaultDevideModel) TableName() string {
	return r.table
}

func (r *defaultDevideModel) CreateTable() (err error) {
	err = r.conn.Table(r.table).AutoMigrate(&DevideModel{})
	return
}

func (r *defaultDevideModel) DropTable() (err error) {
	err = r.conn.Table(r.table).Migrator().DropTable(r.table)
	return
}

func (r *defaultDevideModel) GetTx() (tx *gorm.DB) {
	return r.conn
}

func (r *defaultDevideModel) InitData(lang, uid, cid, did string) (err error) {
	m := &DevideModel{
		OnlyCode:   did,
		CreatedBy:  uid,
		ClassifyId: cid,
		Sort:       0,
	}
	switch lang {
	case constx.LangChinese:
		m.Title = "未分类"
	default:
		m.Title = "ungrouped"
	}
	_, err = r.Insert(m)
	return
}

func (r *defaultDevideModel) Insert(tm *DevideModel) (id uint, err error) {
	err = r.conn.Table(r.table).Create(tm).Error
	id = tm.Id
	return
}

func (r *defaultDevideModel) InsertAll(tm []*DevideModel) (err error) {
	err = r.conn.Table(r.table).Create(tm).Error
	return
}
