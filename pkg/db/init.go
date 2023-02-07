package db

import (
	"github.com/clz.skywalker/event.shop/kernal/pkg/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// 数据库初始化状态
type DbInitStateType int
type CreateTableFunc func() error
type DropTableFunc func() error
type autoMigrateFunc func() (err error)

const (
	lastVersion                   = 1
	DbInitFailure DbInitStateType = iota - 1
	DbCreating
	DbUpgrading
	DbInitSuccess
)

func InitDatabase(dbPath, mode string) (db *gorm.DB, idb iinitDb, err error) {
	utils.ZapLog.Info("init database start")
	db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return
	}
	idb = &sqliteDbStruct{
		db:          db,
		lastVersion: lastVersion,
		migrateList: make([]autoMigrateFunc, 0),
	}

	err = idb.GetVersion()
	if err != nil {
		return
	}
	return
}

type iinitDb interface {
	GetVersion() (err error)
	SetVersion() (err error)
	SetCreateFunc(...CreateTableFunc)
	SetDropFunc(...DropTableFunc)
	OnInitDb(mode string, ch chan<- DbInitStateType) (err error)
	onCreate() (err error)
	onUpgrade() (err error)
	onInitData() (err error)
	onDrop() (err error)
}
