package model

import (
	"github.com/clz.skywalker/event.shop/kernal/pkg/constx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/utils"
	"gorm.io/gorm"
)

// 分类
type ClassifyModel struct {
	BaseModel
	OnlyCode  string `gorm:"column:oc;type:VARCHAR(26);index:udx_classify_oc,unique"`
	CreatedBy string `gorm:"column:created_by;type:VARCHAR(26);index:idx_cm_uid_tid_add,priority:1"`
	TeamId    string `gorm:"column:team_id;type:VARCHAR(26);index:idx_cm_uid_tid_add,priority:2"`
	Title     string `gorm:"column:title;type:varchar"`
	Color     string `gorm:"column:color;type:varchar"`
	Sort      int    `gorm:"column:sort;type:INTEGER"`
}

type IClassifyModel interface {
	IBaseModel
	InitData(lang, uid, tid, cid string) (err error)
	QueryAll() ([]ClassifyModel, error)
	QueryById(id uint) (result ClassifyModel, err error)
	QueryByUid(uid string) ([]ClassifyModel, error)
	QueryByTitle(uid, title string) (ClassifyModel, error)
	QueryByUidAndTitle(uid, title string) (cm ClassifyModel, err error)
	Insert(*ClassifyModel) (uint, error)
	InsertAll(cm []*ClassifyModel) (err error)
	Update(ClassifyModel) error
	Delete(id uint, uid string) error
}

type defaultClassifyModel struct {
	conn  *gorm.DB
	table string
}

func NewDefaultClassifyModel(conn *gorm.DB) IClassifyModel {
	return &defaultClassifyModel{
		conn:  conn,
		table: ClassifyTableName,
	}
}

func (m *defaultClassifyModel) TableName() string {
	return m.table
}

func (m *defaultClassifyModel) CreateTable() (err error) {
	err = m.conn.Table(m.table).AutoMigrate(&ClassifyModel{})
	return
}

func (m *defaultClassifyModel) DropTable() (err error) {
	err = m.conn.Table(m.table).Migrator().DropTable(m.table)
	return
}

func (m *defaultClassifyModel) InitData(lang, uid, tid, cid string) (err error) {
	cm1 := &ClassifyModel{OnlyCode: cid, CreatedBy: uid, TeamId: tid,
		Title: "", Color: "", Sort: 0}
	cm2 := &ClassifyModel{OnlyCode: utils.NewUlid(), CreatedBy: uid, TeamId: tid,
		Title: "", Color: "", Sort: 1}
	cm3 := &ClassifyModel{OnlyCode: utils.NewUlid(), CreatedBy: uid, TeamId: tid,
		Title: "", Color: "", Sort: 2}
	cm4 := &ClassifyModel{OnlyCode: utils.NewUlid(), CreatedBy: uid, TeamId: tid,
		Title: "", Color: "", Sort: 3}
	switch lang {
	case constx.LangChinese:
		cm1.Title = "紧急&重要"
		cm1.Color = "#fd8f80"
		cm2.Title = "紧急&不重要"
		cm2.Color = "#a0cb62"
		cm3.Title = "不紧急&重要"
		cm3.Color = "#4ac0e4"
		cm4.Title = "不紧急&不重要"
		cm4.Color = "#b4b4b4"
	default:
		cm1.Title = "Important & Urgent"
		cm1.Color = "#fd8f80"
		cm2.Title = "Important & Not Urgent"
		cm2.Color = "#a0cb62"
		cm3.Title = "Not Important & Urgent"
		cm3.Color = "#4ac0e4"
		cm4.Title = "Not Important & Not Urgent"
		cm4.Color = "#b4b4b4"
	}
	err = m.InsertAll([]*ClassifyModel{cm1, cm2, cm3, cm4})
	return
}

func (m *defaultClassifyModel) QueryAll() (cms []ClassifyModel, err error) {
	err = m.conn.Table(m.table).Find(&cms).Error
	return
}

func (m *defaultClassifyModel) QueryById(id uint) (result ClassifyModel, err error) {
	err = m.conn.Table(m.table).Where(ClassifyModel{BaseModel: BaseModel{Id: id}}).First(&result).Error
	return
}

func (m *defaultClassifyModel) QueryByUid(uid string) (result []ClassifyModel, err error) {
	err = m.conn.Table(m.table).Where(ClassifyModel{CreatedBy: uid}).Find(&result).Error
	return
}

func (m *defaultClassifyModel) QueryByTitle(uid, title string) (result ClassifyModel, err error) {
	err = m.conn.Table(m.table).Where(ClassifyModel{CreatedBy: uid, Title: title}).Find(&result).Error
	return
}

func (m *defaultClassifyModel) QueryByUidAndTitle(uid, title string) (cm ClassifyModel, err error) {
	err = m.conn.Table(m.table).Where(ClassifyModel{CreatedBy: uid, Title: title}).First(&cm).Error
	return
}

func (m *defaultClassifyModel) Insert(cm *ClassifyModel) (id uint, err error) {
	err = m.conn.Table(m.table).Create(cm).Error
	id = cm.Id
	return
}

func (m *defaultClassifyModel) InsertAll(cm []*ClassifyModel) (err error) {
	err = m.conn.Table(m.table).Create(cm).Error
	return
}

func (m *defaultClassifyModel) Update(cm ClassifyModel) (err error) {
	err = m.conn.Table(m.table).Updates(cm).Error
	return
}

func (m *defaultClassifyModel) Delete(id uint, uid string) (err error) {
	err = m.conn.Table(m.table).Delete(ClassifyModel{BaseModel: BaseModel{Id: id}, CreatedBy: uid}).Error
	return
}
