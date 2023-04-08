package model

import (
	"fmt"

	"github.com/clz.skywalker/event.shop/kernal/internal/entity"
	"github.com/clz.skywalker/event.shop/kernal/pkg/constx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/utils"
	"gorm.io/gorm"
)

type TaskModel struct {
	BaseModel
	OnlyCode    string `gorm:"column:oc;type:VARCHAR(26);index:udx_task_oc,unique"`
	CreatedBy   string `gorm:"column:created_by;type:VARCHAR(26);index:idx_task_uid"`
	Title       string `gorm:"column:title;type:VARCHAR(255)"`
	ClassifyId  string `gorm:"column:classify_id;type:VARCHAR(26);index:idx_task_cid"`
	ContentId   string `gorm:"column:content_id;type:VARCHAR(26)"`
	TaskModeId  string `gorm:"column:task_mode_id;type:VARCHAR(26)"`
	CompletedAt int64  `gorm:"column:completed_at;type:INTEGER"`
	GiveUpAt    int64  `gorm:"column:give_up_at;type:INTEGER"`
	StartAt     int64  `gorm:"column:start_at;type:INTEGER;index"`
	EndAt       int64  `gorm:"column:end_at;type:INTEGER;index"`
	ParentId    string `gorm:"parent_id;type:varchar(26)"`
}

type ITaskModel interface {
	IBaseModel
	InitData(lang, uid, cid, tid string) (err error)
	FindByClassifyId(classifyId string) (result []entity.TaskEntity, err error)
	FindByModel(TaskModel) ([]TaskModel, error)
	First(param TaskModel) (result TaskModel, err error)
	Insert(*TaskModel) (uint, error)
	InsertAll(tm []*TaskModel) (err error)
	Update(*TaskModel) error
	Delete(oc string) error
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

func (m *defaultTaskModel) GetTx() (tx *gorm.DB) {
	return m.conn
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
	err = m.InsertAll([]*TaskModel{t1, t2})
	return
}

func (m *defaultTaskModel) FindByModel(TaskModel) (result []TaskModel, err error) {
	return
}

func (m *defaultTaskModel) FindByClassifyId(classifyId string) (result []entity.TaskEntity, err error) {
	where := fmt.Sprintf("and t1.classify_id='%s' and t1.deleted_at=0", classifyId)
	err = m.conn.Raw(fmt.Sprintf(recursive(m.table, where)+`
	SELECT t1.*,tc.content 
	FROM all_folders t1  left join 
	%s tc on t1.content_id =tc.oc where t1.classify_id='%s' and t1.deleted_at=0 ;`, TaskContentModel{}.TableName(),
		classifyId)).Scan(&result).Error
	return
}

func (m *defaultTaskModel) First(param TaskModel) (result TaskModel, err error) {
	err = m.conn.Table(m.table).Where(param).First(&result).Error
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

func (m *defaultTaskModel) Delete(oc string) (err error) {
	err = m.conn.Table(m.table).Delete(&TaskModel{OnlyCode: oc}).Error
	return
}
