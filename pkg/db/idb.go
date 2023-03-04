package db

import (
	"github.com/clz.skywalker/event.shop/kernal/pkg/constx"
	"gorm.io/gorm"
)

type IOriginDb interface {
	GetVersion() (err error)
	SetVersion() (err error)
	SetDb(*gorm.DB)
	SetCreateFunc(...constx.CreateTableFunc)
	SetDropFunc(...constx.DropTableFunc)
	OnInitDb(mode string, ch chan<- constx.DbInitStateType) (err error)
	onCreate() (err error)
	onUpgrade() (err error)
	onInitData() (err error)
	onDrop() (err error)
}
