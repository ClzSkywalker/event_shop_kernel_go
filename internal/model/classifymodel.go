package model

import (
	"fmt"

	"github.com/clz.skywalker/event.shop/kernal/pkg/constx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/utils"
	"gorm.io/gorm"
)

// 分类
type ClassifyModel struct {
	BaseModel
	OnlyCode  string               `gorm:"column:oc;type:VARCHAR(26);index:udx_classify_oc,unique"`
	CreatedBy string               `gorm:"column:created_by;type:VARCHAR(26);index:idx_cm_uid_tid_add,priority:1"`
	TeamId    string               `gorm:"column:team_id;type:VARCHAR(26);index:idx_cm_uid_tid_add,priority:2"`
	Title     string               `gorm:"column:title;type:varchar"`
	Color     string               `gorm:"column:color;type:varchar"`
	ShowType  constx.TaskShowType  `gorm:"column:show_type;type:INTEGER"` // 展示模式
	OrderType constx.TaskOrderType `gorm:"column:order_type;type:INTEGER"`
	Sort      int                  `gorm:"column:sort;type:INTEGER"`
	ParentId  string               `gorm:"parent_id;type:varchar(26)"`
}

type IClassifyModel interface {
	IBaseModel
	InitData(lang, uid, tid, cid string) (err error)
	FindByTeamId(teamId string) ([]ClassifyModel, error)
	First(p ClassifyModel) (result ClassifyModel, err error)
	Insert(*ClassifyModel) (uint, error)
	InsertAll(cm []*ClassifyModel) (err error)
	Update(ClassifyModel) error
	Delete(tid, oc string) error
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

func (m *defaultClassifyModel) GetTx() (tx *gorm.DB) {
	return m.conn
}

func (m *defaultClassifyModel) InitData(lang, uid, tid, cid string) (err error) {
	cm1 := &ClassifyModel{OnlyCode: cid, CreatedBy: uid, TeamId: tid,
		Title: "", Color: "#fd8f80", Sort: 0}
	cm2 := &ClassifyModel{OnlyCode: utils.NewUlid(), CreatedBy: uid, TeamId: tid,
		Title: "", Color: "#a0cb62", Sort: 1}
	cm3 := &ClassifyModel{OnlyCode: utils.NewUlid(), CreatedBy: uid, TeamId: tid,
		Title: "", Color: "#4ac0e4", Sort: 2}
	cm4 := &ClassifyModel{OnlyCode: utils.NewUlid(), CreatedBy: uid, TeamId: tid,
		Title: "", Color: "#b4b4b4", Sort: 3}
	cm5 := &ClassifyModel{OnlyCode: utils.NewUlid(), CreatedBy: uid, TeamId: tid,
		Title: "", Color: "#b4b4b4", Sort: 4}
	switch lang {
	case constx.LangChinese:
		cm1.Title = "紧急&重要"
		cm2.Title = "紧急&不重要"
		cm3.Title = "不紧急&重要"
		cm4.Title = "不紧急&不重要"
		cm5.Title = "未分类"
	default:
		cm1.Title = "Important & Urgent"
		cm2.Title = "Important & Not Urgent"
		cm3.Title = "Not Important & Urgent"
		cm4.Title = "Not Important & Not Urgent"
		cm5.Title = "unclassified"
	}
	err = m.InsertAll([]*ClassifyModel{cm1, cm2, cm3, cm4})
	return
}

func (m *defaultClassifyModel) FindByTeamId(teamId string) (cms []ClassifyModel, err error) {
	where := fmt.Sprintf("and t1.team_id='%s' and t1.deleted_at=0", teamId)
	err = m.conn.Raw(fmt.Sprintf(recursive(m.table, where)+`
	  SELECT *
	  FROM all_folders where team_id='%s' and deleted_at=0;`, m.table, m.table, m.table, teamId)).Scan(&cms).Error
	return
}

func (m *defaultClassifyModel) First(p ClassifyModel) (result ClassifyModel, err error) {
	err = m.conn.Table(m.table).Where(p).First(&result).Error
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

func (m *defaultClassifyModel) Delete(tid, oc string) (err error) {
	err = m.conn.Table(m.table).Delete(ClassifyModel{OnlyCode: oc, TeamId: tid}).Error
	return
}
