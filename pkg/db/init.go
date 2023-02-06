package db

import (
	"github.com/clz.skywalker/event.shop/kernal/pkg/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// 数据库初始化状态
type DbInitStateType int

const (
	lastVersion                   = 1
	DbInitFailure DbInitStateType = iota - 1
	DbCreating
	DbUpgrading
	DbInitSuccess
)

func InitDatabase(dbPath string, ch chan<- DbInitStateType) (db *gorm.DB, err error) {
	defer func(ch chan<- DbInitStateType) {
		if err != nil {
			ch <- DbInitFailure
		}
	}(ch)
	utils.ZapLog.Info("init database start")
	db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return
	}
	var s iinitDb = &sqliteDbStruct{
		db:          db,
		lastVersion: lastVersion,
		migrateList: make([]autoMigrateFunc, 0),
	}
	err = s.getSqliteVersion()
	if err != nil {
		return
	}
	ch <- DbCreating
	err = s.onCreate()
	if err != nil {
		return
	}
	ch <- DbUpgrading
	err = s.onUpgrade()
	if err != nil {
		return
	}
	utils.ZapLog.Info("init database end")
	return
}

type autoMigrateFunc func() (err error)

type iinitDb interface {
	getSqliteVersion() (err error)
	onCreate() (err error)
	onUpgrade() (err error)
}
