package model

import "gorm.io/gorm"

// 分类
type ClassifyModel struct {
	BaseModel
	Title string `json:"title" gorm:"title"`
	Color string `json:"color" gorm:"color"`
	Sort  int    `json:"sort" gorm:"sort"`
}

type IClassifyModel interface {
	GetTableName() string
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

func (m *defaultClassifyModel) GetTableName() string {
	return m.table
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
