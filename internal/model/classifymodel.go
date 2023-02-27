package model

import "gorm.io/gorm"

// 分类
type ClassifyModel struct {
	BaseModel
	Uid   string `json:"uid,omitempty" gorm:"type:TEXT;index:idx_uid,unique"`
	Title string `json:"title" gorm:"type:TEXT"`
	Color string `json:"color" gorm:"type:TEXT"`
	Sort  int    `json:"sort" gorm:"type:INTEGER"`
}

type IClassifyModel interface {
	IBaseModel
	QueryAll() ([]ClassifyModel, error)
	QueryByUid(uid string) ([]ClassifyModel, error)
	QueryByUidAndTitle(uid, title string) (cm ClassifyModel, err error)
	Insert(*ClassifyModel) (uint, error)
	Update(ClassifyModel) error
	Delete(id uint) error
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
	err = m.conn.Table(m.table).Migrator().DropTable(m)
	return
}

func (m *defaultClassifyModel) QueryAll() (cms []ClassifyModel, err error) {
	err = m.conn.Table(m.table).Find(&cms).Error
	return
}

func (m *defaultClassifyModel) QueryByUid(uid string) (result []ClassifyModel, err error) {
	err = m.conn.Table(m.table).Where(ClassifyModel{Uid: uid}).Find(&result).Error
	return
}

func (m *defaultClassifyModel) QueryByUidAndTitle(uid, title string) (cm ClassifyModel, err error) {
	err = m.conn.Table(m.table).Where(ClassifyModel{Uid: uid, Title: title}).First(&cm).Error
	return
}

func (m *defaultClassifyModel) Insert(cm *ClassifyModel) (id uint, err error) {
	err = m.conn.Table(m.table).Create(cm).Error
	id = cm.Id
	return
}

func (m *defaultClassifyModel) Update(cm ClassifyModel) (err error) {
	err = m.conn.Table(m.table).Updates(cm).Error
	return
}

func (m *defaultClassifyModel) Delete(id uint) (err error) {
	err = m.conn.Table(m.table).Delete(ClassifyModel{BaseModel: BaseModel{Id: id}}).Error
	return
}
