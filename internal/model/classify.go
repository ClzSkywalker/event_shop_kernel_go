package model

import "gorm.io/gorm"

// 分类
type ClassifyModel struct {
	BaseModel
	Title string `json:"title" gorm:"type:TEXT"`
	Color string `json:"color" gorm:"type:TEXT"`
	Sort  int    `json:"sort" gorm:"type:INTEGER"`
}

type IClassifyModel interface {
	IBaseModel
	SelectByModel(ClassifyModel) ([]ClassifyModel, error)
	Insert(ClassifyModel) (int64, error)
	Update(ClassifyModel) error
	Delete(int64) error
}

type defaultClassifyModel struct {
	conn  *gorm.DB
	table string
}

func NewDefaultClassifyModel(conn *gorm.DB) *defaultClassifyModel {
	return &defaultClassifyModel{
		conn:  conn,
		table: "classify",
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

func (m *defaultClassifyModel) SelectByModel(ClassifyModel) (result []ClassifyModel, err error) {
	return
}
func (m *defaultClassifyModel) Insert(ClassifyModel) (id int64, err error) {
	return
}
func (m *defaultClassifyModel) Update(ClassifyModel) (err error) {
	return
}
func (m *defaultClassifyModel) Delete(int64) (err error) {
	return
}
