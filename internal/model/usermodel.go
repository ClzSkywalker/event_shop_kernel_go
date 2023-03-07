package model

import (
	"github.com/clz.skywalker/event.shop/kernal/pkg/constx"
	"gorm.io/gorm"
)

type UserModel struct {
	BaseModel
	CreatedBy    string              `json:"uid,omitempty" gorm:"type:VARCHAR(26);index:udx_user_uid,unique"`
	NickName     string              `json:"nick_name,omitempty" gorm:"type:VARCHAR(255)"`
	MemberType   constx.UserType     `json:"member_type,omitempty" gorm:"type:INTEGER"`   // 用户类型
	RegisterType constx.RegisterTypt `json:"register_type,omitempty" gorm:"type:INTEGER"` // 注册方式
	Avatar       string              `json:"avatar,omitempty" gorm:"type:VARCHAR(255)"`
	Email        string              `json:"email,omitempty" gorm:"type:VARCHAR(30);index:idx_user_email"`
	Phone        string              `json:"phone,omitempty" gorm:"type:VARCHAR(30);index:idx_user_phone"`
	Pwd          string              `json:"pwd,omitempty" gorm:"type:VARCHAR"`
	Version      string              `json:"version,omitempty" gorm:"type:VARCHAR(30)"` // 最后一次登录的版本
}

type IUserModel interface {
	IBaseModel
	InitData(lang, uid string) (err error)
	Insert(um *UserModel) (id uint, err error)
	QueryUser(um UserModel) (ru UserModel, err error)
	CheckRegisterRepeat(um UserModel) (ru UserModel, err error)
	CheckBind(um UserModel) (ru UserModel, err error)
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

func (m *defaultUserModel) InitData(lang, uid string) (err error) {
	nickName := ""
	switch lang {
	case constx.LangChinese:
		nickName = "未名"
	default:
		nickName = "sunshine"
	}
	um := UserModel{CreatedBy: uid, NickName: nickName, Version: constx.KernelVersion}
	err = m.Update(um)
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

func (m *defaultUserModel) CheckRegisterRepeat(um UserModel) (ru UserModel, err error) {
	err = m.conn.Table(m.table).Or(UserModel{Email: um.Email}).
		Or(UserModel{Phone: um.Phone}).Or(UserModel{CreatedBy: um.CreatedBy}).First(&ru).Error
	return
}

func (m *defaultUserModel) CheckBind(um UserModel) (ru UserModel, err error) {
	err = m.conn.Table(m.table).Or(UserModel{Email: um.Email}).
		Or(UserModel{Phone: um.Phone}).First(&ru).Error
	return
}

func (m *defaultUserModel) Update(um UserModel) (err error) {
	err = m.conn.Table(m.table).Where(UserModel{CreatedBy: um.CreatedBy}).Updates(um).Error
	return
}

func (m *defaultUserModel) Delete(id uint) (err error) {
	err = m.conn.Table(m.table).Delete(UserModel{BaseModel: BaseModel{Id: id}}).Error
	return
}
