package model

import (
	"github.com/clz.skywalker/event.shop/kernal/pkg/constx"
	"gorm.io/gorm"
)

type UserModel struct {
	BaseModel
	Uid          string              `json:"uid,omitempty" gorm:"type:VARCHAR(36);index:udx_user_uid,unique"`
	NickName     string              `json:"nick_name,omitempty" gorm:"type:VARCHAR(255)"`
	MemberType   constx.UserType     `json:"member_type,omitempty" gorm:"type:INTEGER"`   // 用户类型
	RegisterType constx.RegisterTypt `json:"register_type,omitempty" gorm:"type:INTEGER"` // 注册方式
	Avatar       string              `json:"avatar,omitempty" gorm:"type:VARCHAR(255)"`
	Email        string              `json:"email,omitempty" gorm:"type:VARCHAR(30);index:udx_user_email,unique"`
	Phone        string              `json:"phone,omitempty" gorm:"type:VARCHAR(30);index:udx_user_phone,unique"`
	Pwd          string              `json:"pwd,omitempty" gorm:"type:VARCHAR"`
	Version      string              `json:"version,omitempty" gorm:"type:VARCHAR(30)"` // 最后一次登录的版本
}

type IUserModel interface {
	IBaseModel
	Insert(um *UserModel) (id uint, err error)
	QueryUser(um UserModel) (ru UserModel, err error)
	Update(um UserModel) (err error)
	Delete(id uint) (err error)
}

type defaultUserModel struct {
	conn  *gorm.DB
	table string
}

func NewDefaultUserModel(conn *gorm.DB) IUserModel {
	return &defaultUserModel{
		conn:  conn,
		table: UserTableName,
	}
}

func (m *defaultUserModel) TableName() string {
	return m.table
}

func (m *defaultUserModel) CreateTable() (err error) {
	err = m.conn.Table(m.table).AutoMigrate(&UserModel{})
	return
}

func (m *defaultUserModel) DropTable() (err error) {
	err = m.conn.Table(m.table).Migrator().DropTable(m.table)
	return
}

func (m *defaultUserModel) Insert(um *UserModel) (id uint, err error) {
	err = m.conn.Table(m.table).Create(um).Error
	id = um.Id
	return
}

func (m *defaultUserModel) QueryUser(um UserModel) (ru UserModel, err error) {
	err = m.conn.Table(m.table).Where(um).First(&ru).Error
	return
}

func (m *defaultUserModel) Update(um UserModel) (err error) {
	err = m.conn.Table(m.table).Updates(um).Error
	return
}

func (m *defaultUserModel) Delete(id uint) (err error) {
	err = m.conn.Table(m.table).Delete(UserModel{BaseModel: BaseModel{Id: id}}).Error
	return
}
